package main

import (
	"log"
	"todo-api/internal/config"
	"todo-api/internal/database"
	"todo-api/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	/* Config load */
	var cfg *config.Config
	var err error
	cfg, err = config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	/* estd database connection */
	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer pool.Close()

	/* Router setup */
	var router *gin.Engine
	router = gin.Default() // or router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {

		// H =>
		// map[string]interface{} => older go way
		// map[string]any{} => new go way
		c.JSON(200, gin.H{
			"message":  "Todo API is running",
			"status":   "Success",
			"database": "Connected",
		})
	})

	router.POST("/todos", handlers.CreateTodoHandler(pool))

	router.Run(":" + cfg.Port)

}
