package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func HashPassword(password string) (string ,error){

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}



func CheckPasswordHash(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func main() {

	password := "secret"

	hash, _:= HashPassword(password)

	fmt.Println("password", password)
	fmt.Println("hash:    ", hash )

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)

}