// Code generated by sqlc. DO NOT EDIT.
// source: notes.sql

package db

import (
	"context"
)

const createNote = `-- name: CreateNote :one
INSERT INTO notes (author, post_id, body, page, line)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, author, post_id, body, page, line, created_at
`

type CreateNoteParams struct {
	Author string `json:"author"`
	PostID int64  `json:"post_id"`
	Body   string `json:"body"`
	Page   int16  `json:"page"`
	Line   int16  `json:"line"`
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, createNote,
		arg.Author,
		arg.PostID,
		arg.Body,
		arg.Page,
		arg.Line,
	)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.PostID,
		&i.Body,
		&i.Page,
		&i.Line,
		&i.CreatedAt,
	)
	return i, err
}

const deleteNote = `-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1
`

func (q *Queries) DeleteNote(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteNote, id)
	return err
}

const deleteNotes = `-- name: DeleteNotes :exec
DELETE FROM notes
WHERE post_id = $1
`

func (q *Queries) DeleteNotes(ctx context.Context, postID int64) error {
	_, err := q.db.ExecContext(ctx, deleteNotes, postID)
	return err
}

const getNote = `-- name: GetNote :one
SELECT id, author, post_id, body, page, line, created_at
FROM notes
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetNote(ctx context.Context, id int64) (Note, error) {
	row := q.db.QueryRowContext(ctx, getNote, id)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.PostID,
		&i.Body,
		&i.Page,
		&i.Line,
		&i.CreatedAt,
	)
	return i, err
}

const listNotes = `-- name: ListNotes :many
SELECT id, author, post_id, body, page, line, created_at
FROM notes
WHERE post_id = $1
ORDER BY id
LIMIT $2 OFFSET $3
`

type ListNotesParams struct {
	PostID int64 `json:"post_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListNotes(ctx context.Context, arg ListNotesParams) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, listNotes, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Note{}
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.PostID,
			&i.Body,
			&i.Page,
			&i.Line,
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

const updateNote = `-- name: UpdateNote :one
UPDATE notes
SET body = $2,
  page = $3,
  line = $4
WHERE id = $1
RETURNING id, author, post_id, body, page, line, created_at
`

type UpdateNoteParams struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
	Page int16  `json:"page"`
	Line int16  `json:"line"`
}

func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, updateNote,
		arg.ID,
		arg.Body,
		arg.Page,
		arg.Line,
	)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.PostID,
		&i.Body,
		&i.Page,
		&i.Line,
		&i.CreatedAt,
	)
	return i, err
}
