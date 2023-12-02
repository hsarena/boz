package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/hsarena/boz/db/sqlc"
	"github.com/hsarena/boz/schemas"
)

type SheepController struct {
	db *db.Queries
}

func NewSheepController(db *db.Queries) *SheepController {
	return &SheepController{db}
}

func (shc *SheepController) CreateSheep(ctx *gin.Context) {
	var payload *schemas.CreateSheep

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.CreateSheepParams{
		Name:  payload.Name,
		Breed: payload.Breed,
		Wool:  payload.Wool,
		Color: payload.Color,
	}

	sheep, err := shc.db.CreateSheep(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "sheep": sheep})
}

func (shc *SheepController) GetSheep(ctx *gin.Context) {

	sheepId, _ := strconv.Atoi(ctx.Param("sheepId"))
	sheep, err := shc.db.GetSheep(ctx, int64(sheepId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No sheep with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "sheep": sheep})
}

func (shc *SheepController) GetAllSheeps(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	args := &db.ListSheepsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	Sheeps, err := shc.db.ListSheeps(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if Sheeps == nil {
		Sheeps = []db.Sheep{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(Sheeps), "data": Sheeps})
}

func (shc *SheepController) CallSheep(ctx *gin.Context) {

	sheepId, _ := strconv.Atoi(ctx.Param("sheepId"))
	sheep, err := shc.db.GetSheep(ctx, int64(sheepId))

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"Sheep": sheep.Name, "Baa": "Baaa"}})
}
