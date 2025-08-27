package main

import (
	"github.com/LucasAntonioC-137/crud-go/src/controller"
	"github.com/LucasAntonioC-137/crud-go/src/model/repository"
	"github.com/LucasAntonioC-137/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface{

	//Init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)

}