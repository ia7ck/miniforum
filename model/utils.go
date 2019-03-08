package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hash(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(b)
}
