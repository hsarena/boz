package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hsarena/boz/controllers"
)

type SheepRoutes struct {
	sheepController controllers.SheepController
}

func NewSheepRoutes(sheepController controllers.SheepController) SheepRoutes {
	return SheepRoutes{sheepController}
}

func (shr *SheepRoutes) SheepRoute(rg *gin.RouterGroup) {

	router := rg.Group("/boz")
	router.POST("/baa", shr.sheepController.CallSheep)
	router.GET("/sheeps", shr.sheepController.GetAllSheeps)
	router.GET("/sheeps/:id", shr.sheepController.GetSheep)
	router.POST("/sheeps", shr.sheepController.CreateSheep)
}
