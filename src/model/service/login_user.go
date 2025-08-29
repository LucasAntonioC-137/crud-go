package service

import (

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(userDomain model.UserDomainInterface)( 
	model.UserDomainInterface, string, *rest_err.RestErr){
	logger.Info("Init userLogin model.",
			zap.String("journey", "loginUser"))
	
	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordService(
		userDomain.GetEmail(),
		userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
	)
	return user, token, nil
} 