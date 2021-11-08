package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type commentResponse struct {
	Id          int64     `json:"id"`
	Author      string    `json:"author"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	AuthorImage string    `json:"authorImage"`
	PostID      int64     `json:"post_id"`
}

func newCommentResponse(comment db.Comment, authorImage string) commentResponse {
	return commentResponse{
		Id:          comment.ID,
		Author:      comment.Author,
		Body:        comment.Body,
		CreatedAt:   comment.CreatedAt,
		PostID:      comment.PostID,
		AuthorImage: authorImage,
	}
}

type createCommentRequest struct {
	PostID int64  `json:"postId" binding:"required,min=1"`
	Body   string `json:"body" binding:"required"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	post, err := server.store.GetPost(ctx, req.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if post.Author == authPayload.Username {
		err := errors.New("Not allowed to comment on your own post")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	arg := db.CreateCommentParams{
		Author: authPayload.Username,
		PostID: req.PostID,
		Body:   req.Body,
	}

	comment, err := server.store.CreateComment(ctx, arg)
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

	ctx.JSON(http.StatusOK, comment)
}

type listCommentsRequest struct {
	PostID int64 `uri:"postId" binding:"required,min=1"`
}
type listCommentsParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listComments(ctx *gin.Context) {
	var req listCommentsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var param listCommentsParams
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListCommentsParams{
		PostID: req.PostID,
		Limit:  param.PageSize,
		Offset: (param.PageID - 1) * param.PageSize,
	}

	comments, err := server.store.ListComments(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []commentResponse
	for i := range comments {
		comment := comments[i]
		authorImage, err := server.store.GetUserImage(ctx, comment.Author)
		preUrl := getPresignedUrl(authorImage)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		rsp := newCommentResponse(comment, preUrl)
		response = append(response, rsp)
	}
	ctx.JSON(http.StatusOK, response)
}

type deleteCommentRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteComment(ctx *gin.Context) {
	var req deleteCommentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	comment, err := server.store.GetComment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if comment.Author != authPayload.Username {
		err := errors.New("Not allowed to other user's comment")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	err = server.store.DeleteComment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, nil)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
