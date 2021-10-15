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

type createNoteRequest struct {
	PostID int64  `json:"postId" binding:"required,min=1"`
	Body   string `json:"body" binding:"required"`
	Page   int16  `json:"page"`
	Line   int16  `json:"line"`
}

func (server *Server) createNote(ctx *gin.Context) {
	var req createNoteRequest
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
	if post.Author != authPayload.Username {
		err := errors.New("Not allowed to create note on other user's post")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	arg := db.CreateNoteParams{
		Author: authPayload.Username,
		PostID: req.PostID,
		Body:   req.Body,
		Page:   req.Page,
		Line:   req.Line,
	}

	comment, err := server.store.CreateNote(ctx, arg)
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
