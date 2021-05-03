package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) string {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	} 

	return string(hash)
}

func Compare(hashedPwd string, plainPwd string) bool { 
	byteHash := []byte(hashedPwd)
	bytePlainPwd := []byte(plainPwd)
	log.Println("Password:", byteHash, bytePlainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
