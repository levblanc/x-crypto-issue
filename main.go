package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte("123123123"), 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)
}
