package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createBadgeRequest struct {
	Name string `json:"name" binding:"required,alphanum"`
}

func (server *Server) createBadge(ctx *gin.Context) {
	var req createBadgeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	badge, err := server.store.CreateBadge(ctx, req.Name)
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

	ctx.JSON(http.StatusOK, badge)
}
