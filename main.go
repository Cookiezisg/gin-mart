package main

import (
	"fmt"
	"golang-mvc-ecommerce/utils/hash"
)

func main() {
	// Example usage of the hash package
	password := "mySecurePassword"

	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hashed Password:", hashedPassword)

	isValid := hash.CheckPasswordHash(password, hashedPassword)
	if isValid {
		println("Password is valid!")
	} else {
		println("Invalid password.")
	}
}
