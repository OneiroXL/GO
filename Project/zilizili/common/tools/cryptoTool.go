package cryptoTool

import (
	"golang.org/x/crypto/bcrypt"
)

/*
加密密码
*/
func CryptoPassword(password string,cost int) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "",err
	}
	return string(bytes),err
}

/*
验证密码
*/
func CheckPassword(password string,cryptoPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cryptoPassword), []byte(password))
	return err == nil
}
