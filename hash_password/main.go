package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	userPassword := "123456"
	passwordbyte, err := GeneratePassword(userPassword)
	if err != nil {
		fmt.Println("wrong hash")
	}
	fmt.Println(passwordbyte)

	mysql_password := "$2a$10$I8WaWXgiBw8j/IBejb3t/.s5NoOYLvoQzL6mIM2g3TEu4z0HenzqK"
	isOk, err := ValidatePassword(userPassword, mysql_password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(isOk)
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false, errors.New("password wrong")
	}
	return true, nil
}
