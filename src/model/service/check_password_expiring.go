package service

import (
	"fmt"
	"time"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) CheckExpiringPasswords() {
	users, err := ud.userRepository.FindAll()
	if err != nil {
		logger.Error("Error trying to call repository", err,
			zap.String("journey", "checkExpiring"))
		return
	}
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now()
	for _, user := range users {
		if user.GetPasswordExpiration().Sub(now) <= 72*time.Hour && user.GetPasswordExpiration().After(now) {
			exp := user.GetPasswordExpiration().In(loc)
			fmt.Printf("⚠️ Usuário %s tem senha expirando em %v\n",
			user.GetEmail(), exp.Format("02/01/2006 15:04:05"))

		}
	}
}