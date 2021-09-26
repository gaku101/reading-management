// Code generated by sqlc. DO NOT EDIT.
// source: follow.sql

package db

import (
	"context"
)

const createFollow = `-- name: CreateFollow :one
INSERT INTO follow (following_id, follower_id)
VALUES ($1, $2)
RETURNING id, following_id, follower_id
`

type CreateFollowParams struct {
	FollowingID int64 `json:"following_id"`
	FollowerID  int64 `json:"follower_id"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error) {
	row := q.db.QueryRowContext(ctx, createFollow, arg.FollowingID, arg.FollowerID)
	var i Follow
	err := row.Scan(&i.ID, &i.FollowingID, &i.FollowerID)
	return i, err
}
