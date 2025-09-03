package model

import "time"

type userDomain struct {
	ID                  string
	email               string
	password            string
	name                string
	age                 int8
	password_expiration time.Time
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}
func (ud userDomain) GetID() string {
	return ud.ID
}
func (ud userDomain) GetEmail() string {
	return ud.email
}
func (ud userDomain) GetPassword() string {
	return ud.password
}
func (ud userDomain) GetName() string {
	return ud.name
}
func (ud userDomain) GetAge() int8 {
	return ud.age
}

func (ud userDomain) GetPasswordExpiration() time.Time {
	return ud.password_expiration
}

func (ud *userDomain) SetPasswordExpiration(t time.Time) {
	ud.password_expiration = t
}
