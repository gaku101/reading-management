package api

import (
	"database/sql"
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

type getBadgeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBadge(ctx *gin.Context) {
	var req getBadgeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	badge, err := server.store.GetBadge(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, badge)
}

type listBadgesRequest struct {
}

func (server *Server) listBadges(ctx *gin.Context) {
	badges, err := server.store.ListBadges(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, badges)
}
