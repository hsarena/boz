package schemas

import db "github.com/hsarena/boz/db/sqlc"



type CreateSheep struct {
	Name  string `json:"name" binding:"required"`
	Breed string `json:"breed" binding:"required"`
	Wool  db.Wool    `json:"wool" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type UpdateSheep struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Wool  db.Wool   `json:"wool"`
	Color string `json:"color"`
}
