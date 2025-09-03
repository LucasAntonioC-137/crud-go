package model

import (
	"time"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
)

type UserDomainInterface interface{
	GetID() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	SetID(string)
	GetPasswordExpiration() time.Time
	SetPasswordExpiration(time.Time)
	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string, age int8, password_expiration time.Time,
) UserDomainInterface {
	return &userDomain{
		email: email,
		password: password,
		name: name,
		age: age,
		password_expiration: password_expiration, // 90 dias de validade
	}
}

func NewUserUpdateDomain(
	name string, age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age: age,
	}
}

func NewUserLoginDomain(
	email, password string, password_expiration time.Time,
) UserDomainInterface {
	return &userDomain{
		email: email,
		password: password,
		password_expiration: password_expiration,
	}
}