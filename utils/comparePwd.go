package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword, plainPassword string) (error){
	err:=  bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err
}