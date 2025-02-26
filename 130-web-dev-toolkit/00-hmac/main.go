package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

// NOTE: 🔹 HMAC (Hash-based Message Authentication Code)
/*
	✅ Purpose: Ensures data integrity and authenticity using a hash function and a secret key.
	✅ Use Case: Verifying API requests, securing tokens, message authentication.
	✅ Process:
		•	Takes a message and a secret key.
		•	Uses a hash function (e.g., SHA-256).
		•	Produces a fixed-length hash (MAC) for authentication.
	✅ Key Feature: Fast & deterministic (same input = same output).
	❌ Not designed for password hashing—it’s used for verifying message authenticity.
*/

// NOTE: 🔹 Bcrypt
/*
	✅ Purpose: Secure password hashing with a built-in salt and slow hashing for brute-force resistance.
	✅ Use Case: Storing passwords securely in databases.
	✅ Process:
		•	Takes a password and generates a hashed version with a random salt.
		•	Uses the Blowfish cipher and applies multiple rounds of hashing.
		•	Makes brute-force attacks much slower.
	✅ Key Feature: Adaptive security—you can increase the number of hashing rounds as computing power grows.
	❌ Not used for message authentication—bcrypt is for password storage, not integrity verification.
*/

// NOTE: 🔹 HMAC vs Bcrypt: Key Differences
/*
	Feature			HMAC									Bcrypt
	Purpose			Message authentication					Secure password hashing
	Uses Key?		✅ Yes (secret key)						❌ No (uses salt instead)
	Output			Fixed-length hash (e.g., 256-bit)		Variable-length hash with salt
	Speed			⚡ Fast (meant for real-time checks)		🐢 Slow (to resist brute force)
	Use Cases		API authentication, data integrity		Password storage, login securit
*/

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {
	// * Hash-based Message Authentication Code
	h := hmac.New(sha256.New, []byte("our_key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
