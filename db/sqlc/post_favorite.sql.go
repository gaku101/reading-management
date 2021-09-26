// Code generated by sqlc. DO NOT EDIT.
// source: post_favorite.sql

package db

import (
	"context"
)

const createPostFavorite = `-- name: CreatePostFavorite :one
INSERT INTO post_favorites (post_id, user_id)
VALUES ($1, $2)
RETURNING id, post_id, user_id
`

type CreatePostFavoriteParams struct {
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
}

func (q *Queries) CreatePostFavorite(ctx context.Context, arg CreatePostFavoriteParams) (PostFavorite, error) {
	row := q.db.QueryRowContext(ctx, createPostFavorite, arg.PostID, arg.UserID)
	var i PostFavorite
	err := row.Scan(&i.ID, &i.PostID, &i.UserID)
	return i, err
}

const getMyFavoritePost = `-- name: GetMyFavoritePost :one
SELECT id, post_id, user_id
FROM post_favorites
WHERE post_id = $1
  AND user_id = $2
LIMIT 1
`

type GetMyFavoritePostParams struct {
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
}

func (q *Queries) GetMyFavoritePost(ctx context.Context, arg GetMyFavoritePostParams) (PostFavorite, error) {
	row := q.db.QueryRowContext(ctx, getMyFavoritePost, arg.PostID, arg.UserID)
	var i PostFavorite
	err := row.Scan(&i.ID, &i.PostID, &i.UserID)
	return i, err
}

const getPostFavorite = `-- name: GetPostFavorite :many
SELECT id
FROM post_favorites
WHERE post_id = $1
`

func (q *Queries) GetPostFavorite(ctx context.Context, postID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getPostFavorite, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFavoritePosts = `-- name: ListFavoritePosts :many
SELECT posts.id,
  author,
  title,
  body,
  created_at,
  updated_at
FROM posts
  JOIN post_favorites ON posts.id = post_id
  AND user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3
`

type ListFavoritePostsParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFavoritePosts(ctx context.Context, arg ListFavoritePostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listFavoritePosts, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Title,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
