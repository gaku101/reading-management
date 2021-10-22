package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromUserID int64 `json:"fromUserId" binding:"required,min=1"`
	ToUserID   int64 `json:"toUserId" binding:"required,min=1"`
	Amount     int64 `json:"amount" binding:"required,gt=0"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromUser, err := server.store.GetUserById(ctx, req.FromUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromUser.Username != authPayload.Username {
		err := errors.New("from user id doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	_, err = server.store.GetUserById(ctx, req.ToUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.TransferTxParams{
		FromUserID: req.FromUserID,
		ToUserID:   req.ToUserID,
		Amount:     req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
