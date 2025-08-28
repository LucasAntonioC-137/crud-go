package service

import (
	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser service", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository", err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("DeleteUser service executed sucessfully",
		zap.String("userid", userId),
		zap.String("journey", "deleteUser"))

	return nil
}