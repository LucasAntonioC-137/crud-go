package service

import (
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"github.com/LucasAntonioC-137/crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIDService(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(string) *rest_err.RestErr
	LoginUserService(model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
	CheckExpiringPasswords()
}
