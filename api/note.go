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

type listNotesRequest struct {
	PostID int64 `uri:"postId" binding:"required,min=1"`
}
type listNotesParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listNotes(ctx *gin.Context) {
	var req listNotesRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var param listNotesParams
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListNotesParams{
		PostID: req.PostID,
		Limit:  param.PageSize,
		Offset: (param.PageID - 1) * param.PageSize,
	}
	notes, err := server.store.ListNotes(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, notes)
}

type updateNoteRequest struct {
	ID   int64  `json:"id" binding:"required,min=1"`
	Body string `json:"body" binding:"required"`
	Page int16  `json:"page"`
	Line int16  `json:"line"`
}

func (server *Server) updateNote(ctx *gin.Context) {
	var req updateNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	note, err := server.store.GetNote(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if note.Author != authPayload.Username {
		err := errors.New("Not allowed to update other user's note")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdateNoteParams{
		ID:   req.ID,
		Body: req.Body,
		Page: req.Page,
		Line: req.Line,
	}

	note, err = server.store.UpdateNote(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, note)
}

type deleteNotetRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteNote(ctx *gin.Context) {
	var req deleteNotetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	note, err := server.store.GetNote(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if note.Author != authPayload.Username {
		err := errors.New("Not allowed to update other user's note")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	note, err = server.store.DeleteNote(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, nil)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
