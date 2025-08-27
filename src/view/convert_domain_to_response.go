package view

import (
	"github.com/LucasAntonioC-137/crud-go/src/controller/model/response"
	"github.com/LucasAntonioC-137/crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID: userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}