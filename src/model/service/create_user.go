package service

import (
	"time"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}
	userDomain.EncryptPassword()
	userDomain.SetPasswordExpiration(time.Now().AddDate(0, 0, 1)) // AddDate(0, 0, 90)) 90 dias

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("createUser service executed sucessfully",
		zap.String("userid", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}