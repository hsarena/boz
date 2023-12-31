// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: sheep.sql

package db

import (
	"context"
)

const createSheep = `-- name: CreateSheep :one
INSERT INTO Sheeps (
  name,
  breed,
  wool,
  color
) VALUES (
  $1, $2, $3, $4
) RETURNING id, name, breed, wool, color, created_at
`

type CreateSheepParams struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Wool  Wool   `json:"wool"`
	Color string `json:"color"`
}

func (q *Queries) CreateSheep(ctx context.Context, arg CreateSheepParams) (Sheep, error) {
	row := q.db.QueryRowContext(ctx, createSheep,
		arg.Name,
		arg.Breed,
		arg.Wool,
		arg.Color,
	)
	var i Sheep
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Breed,
		&i.Wool,
		&i.Color,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSheep = `-- name: DeleteSheep :exec
DELETE FROM Sheeps
WHERE id = $1
`

func (q *Queries) DeleteSheep(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSheep, id)
	return err
}

const getSheep = `-- name: GetSheep :one
SELECT id, name, breed, wool, color, created_at FROM Sheeps
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSheep(ctx context.Context, id int64) (Sheep, error) {
	row := q.db.QueryRowContext(ctx, getSheep, id)
	var i Sheep
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Breed,
		&i.Wool,
		&i.Color,
		&i.CreatedAt,
	)
	return i, err
}

const listSheeps = `-- name: ListSheeps :many
SELECT id, name, breed, wool, color, created_at FROM Sheeps
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListSheepsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSheeps(ctx context.Context, arg ListSheepsParams) ([]Sheep, error) {
	rows, err := q.db.QueryContext(ctx, listSheeps, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sheep
	for rows.Next() {
		var i Sheep
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Breed,
			&i.Wool,
			&i.Color,
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

const updateSheep = `-- name: UpdateSheep :one
UPDATE Sheeps 
SET name = $2
WHERE id = $1
RETURNING id, name, breed, wool, color, created_at
`

type UpdateSheepParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateSheep(ctx context.Context, arg UpdateSheepParams) (Sheep, error) {
	row := q.db.QueryRowContext(ctx, updateSheep, arg.ID, arg.Name)
	var i Sheep
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Breed,
		&i.Wool,
		&i.Color,
		&i.CreatedAt,
	)
	return i, err
}
