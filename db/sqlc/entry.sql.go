// Code generated by sqlc. DO NOT EDIT.
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (user_id, amount)
VALUES ($1, $2)
RETURNING id, user_id, amount, created_at
`

type CreateEntryParams struct {
	UserID int64 `json:"user_id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.UserID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntry = `-- name: GetEntry :one
SELECT id, user_id, amount, created_at
FROM entries
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, user_id, amount, created_at
FROM entries
WHERE user_id = $1
  AND amount > 0
`

func (q *Queries) ListEntries(ctx context.Context, userID int64) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.CreatedAt,
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
