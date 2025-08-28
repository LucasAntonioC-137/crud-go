package controller

import (
	"net/http"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/validation"
	"github.com/LucasAntonioC-137/crud-go/src/controller/model/request"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"github.com/LucasAntonioC-137/crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var(
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email, 
		userRequest.Password, 
		userRequest.Name, 
		userRequest.Age)

	domainResult, err := uc.service.CreateUserService(domain);	
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return 
	}

	logger.Info("User created successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))
		
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}