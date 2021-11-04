package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gin-gonic/gin"
)

type getUserBadgeRequest struct {
	Username string `uri:"username" binding:"required,alphanum"`
}

type getUserBadgeResponse struct {
	User userResponse `json:"user"`
}

func (server *Server) getUserBadge(ctx *gin.Context) {
	var req getUserBadgeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		err := errors.New("Not allowed to delete other user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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
	entries, err := server.store.ListEntries(ctx, user.ID)
	var total int
	for i := range entries {
		entry := entries[i]
		total += int(entry.Amount)
	}
	userBadge, err := server.store.GetUserBadge(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if userBadge.BadgeID == 1 && total >= 50 {
		_, err = server.store.UpdateUserBadge(ctx, db.UpdateUserBadgeParams{
			UserID:  user.ID,
			BadgeID: 2,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		userBadge, err = server.store.GetUserBadge(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	} else if userBadge.BadgeID == 2 && total == 200 {
		_, err = server.store.UpdateUserBadge(ctx, db.UpdateUserBadgeParams{
			UserID:  user.ID,
			BadgeID: 2,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		userBadge, err = server.store.GetUserBadge(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	} else if userBadge.BadgeID == 3 && total >= 500 {
		_, err = server.store.UpdateUserBadge(ctx, db.UpdateUserBadgeParams{
			UserID:  user.ID,
			BadgeID: 2,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		userBadge, err = server.store.GetUserBadge(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	}

	ctx.JSON(http.StatusOK, userBadge)
}
