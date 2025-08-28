package service

import (
	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId,userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("UpdateUser service executed sucessfully",
		zap.String("userid", userId),
		zap.String("journey", "updateUser"))

	return nil
}