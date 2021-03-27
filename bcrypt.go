package main

import (
	"crypto/sha256"
)

/*
func cryptpass(userpassword string) []byte {

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userpassword), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	fmt.Println("Hash to store:", hash)
	return hash
}

// Store this "hash" somewhere, e.g. in your database

// After a while, the user wants to log in and you need to check the password he entered
func decryptpass(userpassword, hashstore []byte) bool {

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(hashstore), []byte(userpassword)); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
		return false
	}

	fmt.Println("Password was correct!")
	return true
}
*/

func hashfunc(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
