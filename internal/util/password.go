package util

import "golang.org/x/crypto/bcrypt"

var passwordHashingCost = 12

func HashedPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), passwordHashingCost)
	return string(bytes)
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
