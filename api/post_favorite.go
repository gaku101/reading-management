package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createPostFavoriteRequest struct {
	PostID int64 `json:"postId" binding:"required,min=1"`
	UserID int64 `json:"userId" binding:"required,min=1"`
}

func (server *Server) createPostFavorite(ctx *gin.Context) {
	var req createPostFavoriteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostFavoriteParams{
		PostID: req.PostID,
		UserID: req.UserID,
	}
	isAuthorized := server.authorizedUser(ctx, arg.UserID)
	if !isAuthorized {
		return
	}
	postFavorite, err := server.store.CreatePostFavorite(ctx, arg)
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

	ctx.JSON(http.StatusOK, postFavorite)
}

type listFavoritePostsParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}
type listFavoritePostsRequest struct {
	UserID int64 `uri:"userId" binding:"required,min=1"`
}

func (server *Server) listFavoritePosts(ctx *gin.Context) {
	var param listFavoritePostsParams
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req listFavoritePostsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListFavoritePostsParams{
		UserID: req.UserID,
		Limit:  param.PageSize,
		Offset: (param.PageID - 1) * param.PageSize,
	}
	isAuthorized := server.authorizedUser(ctx, arg.UserID)
	if !isAuthorized {
		return
	}
	favoritePosts, err := server.store.ListFavoritePosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []postResponse
	for i := range favoritePosts {
		post := favoritePosts[i]
		category, err := server.store.GetPostCategory(ctx, post.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v's category not set", post.ID)
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

		}
		authorImage, err := server.store.GetUserImage(ctx, post.Author)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		favorites := len(server.getFavoriteCount(ctx, post.ID))
		commentsNum := server.getCommentsCount(ctx, post.ID)
		rsp := newPostResponse(post, category, authorImage, favorites, commentsNum)
		response = append(response, rsp)
	}

	ctx.JSON(http.StatusOK, response)
}

type getPostFavoriteRequest struct {
	PostID int64 `uri:"postId" binding:"required,min=1"`
}

func (server *Server) getPostFavorite(ctx *gin.Context) {
	var req getPostFavoriteRequest
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
	arg := db.GetMyFavoritePostParams{
		PostID: req.PostID,
		UserID: user.ID,
	}
	postFavorite, err := server.store.GetMyFavoritePost(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, nil)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, postFavorite)
}

func (server *Server) authorizedUser(ctx *gin.Context, userId int64) bool {
	user, err := server.store.GetUserById(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Username != authPayload.Username {
		err := errors.New("Not allowed to make other user's favorite")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return false
	}

	return true
}
