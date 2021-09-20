package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createPostRequest struct {
	Author     string `json:"author" binding:"required,alphanum"`
	Title      string `json:"title" binding:"required"`
	Body       string `json:"body" binding:"required"`
	CategoryID int64  `json:"categoryId"`
}
type postResponse struct {
	Id        int64       `json:"id"`
	Author    string      `json:"author"`
	Title     string      `json:"title"`
	Body      string      `json:"body"`
	Category  db.Category `json:"category"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func newPostResponse(post db.Post, category db.Category) postResponse {
	return postResponse{
		Id:        post.ID,
		Author:    post.Author,
		Title:     post.Title,
		Body:      post.Body,
		Category:  category,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, valid := server.validUser(ctx, req.Author)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Username != authPayload.Username {
		err := errors.New("Not allowed to create other user's post")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		Author: req.Author,
		Title:  req.Title,
		Body:   req.Body,
	}

	post, err := server.store.CreatePost(ctx, arg)
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
	category, err := server.store.GetCategory(ctx, req.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("post_id = %v's category not set", post.ID)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	arg2 := db.CreatePostCategoryParams{
		PostID:     post.ID,
		CategoryID: category.ID,
	}
	if category.ID != 0 {
		server.store.CreatePostCategory(ctx, arg2)
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
	}

	rsp := newPostResponse(post, category)

	ctx.JSON(http.StatusOK, rsp)
}

type getPostRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.store.GetPost(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	category, err := server.store.GetPostCategory(ctx, post.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("post_id = %v's category not set", post.ID)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	rsp := newPostResponse(post, category)
	ctx.JSON(http.StatusOK, rsp)
}

type listPostRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listPosts(ctx *gin.Context) {
	var req listPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListPostsParams{
		Author: authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	posts, err := server.store.ListPosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []postResponse
	for i := range posts {
		post := posts[i]
		category, err := server.store.GetPostCategory(ctx, post.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v's category not set", post.ID)
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

		}
		rsp := newPostResponse(post, category)
		response = append(response, rsp)
	}

	ctx.JSON(http.StatusOK, response)
}

type updatePostRequest struct {
	ID         int64  `json:"id" binding:"required,min=1"`
	Author     string `json:"author" binding:"required,alphanum"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryID int64  `json:"categoryId"`
}

func (server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, valid := server.validUser(ctx, req.Author)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Username != authPayload.Username {
		err := errors.New("Not allowed to update other user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdatePostParams{
		ID:    req.ID,
		Title: req.Title,
		Body:  req.Body,
	}

	post, err := server.store.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
	arg2 := db.UpdatePostCategoryParams{
		PostID:     post.ID,
		CategoryID: category.ID,
	}
	arg3 := db.CreatePostCategoryParams{
		PostID:     post.ID,
		CategoryID: category.ID,
	}
	if req.CategoryID != 0 {
		_, err = server.store.UpdatePostCategory(ctx, arg2)
		if err != nil {
			if err == sql.ErrNoRows {
				_, err = server.store.CreatePostCategory(ctx, arg3)
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
	}
	rsp := newPostResponse(post, category)
	ctx.JSON(http.StatusOK, rsp)
}
