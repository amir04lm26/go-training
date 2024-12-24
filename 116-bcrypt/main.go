package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	password = "MY_PASSWORD"
	cost     = bcrypt.DefaultCost // 10
)

func main() {
	b, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("hashed:", b)
	fmt.Println("hashed(string):", string(b))

	err = bcrypt.CompareHashAndPassword(b, []byte(password))
	if err != nil {
		fmt.Println("wrong password!")
		return
	}

	fmt.Println("password is match")
}
