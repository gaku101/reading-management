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

type updatePostCategoryRequest struct {
	PostID     int64 `json:"postId" binding:"required,min=1"`
	CategoryID int64 `json:"categoryId" binding:"required,min=1"`
}

func (server *Server) updatePostCategory(ctx *gin.Context) {
	var req updatePostCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	post, err := server.store.GetPost(ctx, req.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if post.Author != authPayload.Username {
		err := errors.New("Not allowed to update other user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, req.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("post_id = %v's category not set", post.ID)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	arg := db.UpdatePostCategoryParams{
		PostID:     post.ID,
		CategoryID: category.ID,
	}
	arg2 := db.CreatePostCategoryParams{
		PostID:     post.ID,
		CategoryID: category.ID,
	}
	_, err = server.store.UpdatePostCategory(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = server.store.CreatePostCategory(ctx, arg2)
			if err != nil {
				if pqErr, ok := err.(*pq.Error); ok {
					switch pqErr.Code.Name() {
					case "foreign_key_violation", "unique_violation":
						ctx.JSON(http.StatusForbidden, errorResponse(err))
						return
					}
				} else {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, category)
}
