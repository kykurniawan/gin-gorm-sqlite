package main

import (
	"fmt"
	"gin-gorm-sqlite/database"
	"gin-gorm-sqlite/handlers"
	"gin-gorm-sqlite/repositories"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("failed to load env variable")
	}

	database.Connect()
}

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{
		"127.0.0.1",
	})

	noteRepository := repositories.NewNoteRepository(database.DB)

	noteHandler := handlers.NewNoteHandler(noteRepository)

	noteRoutes := router.Group("/notes")

	noteRoutes.GET("/", noteHandler.GetAll)
	noteRoutes.GET("/:id", noteHandler.GetById)
	noteRoutes.POST("/", noteHandler.Create)
	noteRoutes.PUT("/:id", noteHandler.UpdateById)
	noteRoutes.DELETE("/:id", noteHandler.DeleteById)

	router.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))
}
