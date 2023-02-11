package utils

import (
	"fmt"
	"go_learn/common"
	"golang.org/x/crypto/bcrypt"
)

func MakePassword(password string) (string, error) {
	salt := common.Salt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s%s", password, salt)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CheckPassword(password string, sqlPassword string) bool {
	salt := common.Salt
	err := bcrypt.CompareHashAndPassword([]byte(sqlPassword), []byte(fmt.Sprintf("%s%s", password, salt)))
	if err != nil {
		return false
	}
	return true
}