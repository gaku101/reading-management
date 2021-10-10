package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/infrastructure"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gaku101/my-portfolio/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Profile  string `json:"profile"`
	Image    string `json:"image"`
}

type userResponse struct {
	Id                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Profile           string    `json:"profile"`
	Image             string    `json:"image"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Id:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		Profile:           user.Profile,
		Image:             getPresignedUrl(user.Image),
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Profile:        req.Profile,
		Image:          req.Image,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,alphanum"`
}

type getUserResponse struct {
	User userResponse `json:"user"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}

type updateUserRequest struct {
	ID       int64  `json:"id" binding:"required,min=1"`
	Username string `json:"username" binding:"required,alphanum"`
	Profile  string `json:"profile"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, valid := server.validUser(ctx, req.Username)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Username != authPayload.Username {
		err := errors.New("Not allowed to update other user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID:      req.ID,
		Profile: req.Profile,
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := getUserResponse{
		User: newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) validUser(ctx *gin.Context, username string) (db.User, bool) {
	user, err := server.store.GetUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return user, false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return user, false
	}

	return user, true
}

func getPresignedUrl(image string) string {
	if !strings.Contains(image, "amazonaws") {
		return image
	}
	awsS3 := infrastructure.NewAwsS3()
	svc := awsS3.Svc
	slice := strings.Split(image, awsS3.Keys.Folder+"/")
	fileName := slice[1]
	s3req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(awsS3.Config.Aws.S3.Bucket),
		Key:    aws.String(awsS3.Keys.Folder + "/" + fileName),
	})
	urlStr, err := s3req.Presign(120 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("Pre signed url : ", urlStr)
	return urlStr
}
