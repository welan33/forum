package forum

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hpwd(pwd string) string {
	if pwd == "" {
		println("erreur: Password vide")
		return ""
	}
	hash, err := HashPassword(pwd) // ignore error for the sake of simplicity
	if err != nil {
		panic(err)
	}

	/*println("Password:", pwd)
	println("Hash:    ", hash)*/

	return hash
}
