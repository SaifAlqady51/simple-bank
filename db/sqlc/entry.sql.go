// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: entry.sql

package authors

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO enteries (
    account_id,
    amount 
) VALUES (
    $1, $2
)
 RETURNING id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID int64
	Amount    int64
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entery, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entery
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :one
DELETE FROM enteries
WHERE id = $1
RETURNING id, account_id, amount, created_at
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) (Entery, error) {
	row := q.db.QueryRowContext(ctx, deleteEntry, id)
	var i Entery
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM enteries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entery, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entery
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntry = `-- name: ListEntry :many
SELECT id, account_id, amount, created_at FROM enteries
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEntryParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListEntry(ctx context.Context, arg ListEntryParams) ([]Entery, error) {
	rows, err := q.db.QueryContext(ctx, listEntry, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entery
	for rows.Next() {
		var i Entery
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
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

const updateEntry = `-- name: UpdateEntry :one
UPDATE enteries
SET amount = $2
WHERE $1
RETURNING id, account_id, amount, created_at
`

type UpdateEntryParams struct {
	Column1 interface{}
	Amount  int64
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entery, error) {
	row := q.db.QueryRowContext(ctx, updateEntry, arg.Column1, arg.Amount)
	var i Entery
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
