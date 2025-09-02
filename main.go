package main

import (
	"context"
	"log"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/database/mongodb"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/LucasAntonioC-137/crud-go/docs"
)

// @title Meu Primeiro CRUD em Go | Lucas
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @schemes http
// @license MIT
func main() {
	logger.Info("About to start appliccation")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)
	router := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
