package controller

import (
	"net/http"
	"net/mail"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// FindUserByID retrieves user information based on the provided user ID. 
// @Summary Find User by ID
// @Description Retrieves user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Success 400 {object} rest_err.RestErr "Error: Invalid user ID"
// @Success 404 {object} rest_err.RestErr "User not found"
// @Router /getUserById/{userId} [get]
func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller", zap.String("journey", "FindUserByID"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err,
			zap.String("journey", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID service", err,
			zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed sucessfully", zap.String("journey", "FindUserByID"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}


// FindUserByEmail retrieves user information based on the provided user Email. 
// @Summary Find User by Email
// @Description Retrieves user details based on the Email provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userEmail path string true "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid user Email"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserByEmail/{userEmail} [get]
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate email", err,
			zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail service", err,
			zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed sucessfully", zap.String("journey", "FindUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
