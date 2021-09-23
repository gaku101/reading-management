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
	user, err := server.store.GetUserById(ctx, arg.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Username != authPayload.Username {
		err := errors.New("Not allowed to make other user's favorite")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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
