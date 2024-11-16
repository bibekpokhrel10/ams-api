package main

import "golang.org/x/crypto/bcrypt"

func main() {
	hasPassword := "$2a$10$KIZLv3PyK4sZIOjbuxsYDevTHWjK/ZoTj63L02dqhfgSdRoJt4Cv6"
	err := CheckPassword("Bibek@1234", hasPassword)
	if err != nil {
		panic(err)
	}
}
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
