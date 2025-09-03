package controller

import (
	"net/http"
	// "time"

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

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequestCreate true "User information for registration"
// @Param Authorization header string true "Insert your acess token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /createUser [post]
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