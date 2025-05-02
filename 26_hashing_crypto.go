package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "password123"
	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))
	fmt.Println(password)
	fmt.Println(hash)
	fmt.Println(hash512)
	fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}
	signUpHash1 := hashPassword(password, salt)
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("Hash:", signUpHash1)
	hashOrigPass:=sha256.Sum256([]byte(password))
	fmt.Println("Hash pass str without salt", base64.StdEncoding.EncodeToString(hashOrigPass[:]))

	//verify decode
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode", err)
		return
	}
	loginHash := hashPassword(password, decodedSalt)

	if signUpHash1 == loginHash {
		fmt.Println("Password is correct!")
	} else {
		fmt.Println("Password is incorrect!")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
