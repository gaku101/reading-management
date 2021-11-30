package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createFollowRequest struct {
	FollowingID int64 `json:"followingId" binding:"required,min=1"`
}

func (server *Server) createFollow(ctx *gin.Context) {
	var req createFollowRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateFollowParams{
		FollowingID: req.FollowingID,
		FollowerID:  user.ID,
	}
	if arg.FollowingID == arg.FollowerID {
		err := errors.New("can not follow yourself")
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	follow, err := server.store.CreateFollow(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, follow)
}

type getFollowRequest struct {
	FollowingID int64 `uri:"followingId" binding:"required,min=1"`
}

func (server *Server) getFollow(ctx *gin.Context) {
	var req getFollowRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.GetFollowParams{
		FollowingID: req.FollowingID,
		FollowerID:  user.ID,
	}
	follow, err := server.store.GetFollow(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, nil)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, follow)
}

type listFollowParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listFollow(ctx *gin.Context) {
	var param listFavoritePostsParams
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.ListFollowParams{
		FollowerID: user.ID,
		Limit:      param.PageSize,
		Offset:     (param.PageID - 1) * param.PageSize,
	}
	followings, err := server.store.ListFollow(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []db.ListFollowRow
	for i := range followings {
		follow := followings[i]
		preUrl := getPresignedUrl(follow.Image)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		follow.Image = preUrl
		response = append(response, follow)
	}
	ctx.JSON(http.StatusOK, response)
}

type deleteFollowRequest struct {
	FollowongID int64 `uri:"followingId" binding:"required,min=1"`
}

func (server *Server) deleteFollow(ctx *gin.Context) {
	var req deleteFollowRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.DeleteFollowParams{
		FollowingID: req.FollowongID,
		FollowerID:  user.ID,
	}
	err = server.store.DeleteFollow(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, nil)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
