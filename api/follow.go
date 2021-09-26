package api

import (
	"database/sql"
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
