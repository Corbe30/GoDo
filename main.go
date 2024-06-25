package main

import (
	"log"
	"os"

	"github.com/Corbe30/GoDo/database"
	"github.com/Corbe30/GoDo/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	dbUrl := os.Getenv("DBURL")
	if dbUrl == "" {
		dbUrl = "mongodb://localhost:27017"
	}
	database.MongoDBInit(dbUrl)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	routers.TasksRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
