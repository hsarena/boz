package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsarena/boz/config"
	"github.com/hsarena/boz/controllers"
	dbConn "github.com/hsarena/boz/db/sqlc"
	"github.com/hsarena/boz/routes"

	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbConn.Queries

	SheepController controllers.SheepController
	SheepRoutes     routes.SheepRoutes
)

func init() {
	config := config.LoadConfig(".env",".")

	postgresSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	conn, err := sql.Open("postgres", postgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	SheepController = *controllers.NewSheepController(db)
	SheepRoutes = routes.NewSheepRoutes(SheepController)

	server = gin.Default()
}

func main() {
	config := config.LoadConfig(".env",".")

	rg := server.Group("/")

	rg.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Boz Server is healthy"})
	})

	SheepRoutes.SheepRoute(rg)
	log.Fatal(server.Run(":" + config.Port))
}
