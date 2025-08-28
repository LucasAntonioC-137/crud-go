package service

import (
	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDService( id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services", zap.String("journey", "FindUserByID"))

	userByIdRepository, err := ud.userRepository.FindUserByID(id)
	if err != nil {
		logger.Error("Error trying to call repository",err,
		zap.String("journey", "FindUserByID"))
		return nil, err
	}	 	
	
	return userByIdRepository, nil
}

func (ud *userDomainService) FindUserByEmailService( email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("journey", "FindUserByEmail"))

	userByEmailRepository, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call repository",err,
		zap.String("journey", "FindUserBy"))
		return nil, err
	}	 	
	
	return userByEmailRepository, nil
}