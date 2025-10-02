package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// import (
// 	"ecommerce/cmd"
// )

func main() {
	// cmd.Serve()

	secret := []byte("secret")
	message := []byte("hello")

	h := hmac.New(sha256.New, secret)
	h.Write(message)

	text := h.Sum(nil)
	fmt.Println(text)
}

/*
JWT (JSON Web Token) - Authentication
======================================
BASE64 - Method for encoding binary data into an ASCII string format (A-Z, a-z, 0-9, +, /)
       - coverts binary data (images, files, bytes) into a set of 64 printable ASCII characters.

Purpose of BASE64
-----------------
- When transferring data from one system to another system base64 is the faster way


SHA 1 -> SHA 256 -> SHA 512 - Secure Hash Algorithm
		- same input same output always
		- cannot generate input from output


HMAC - Hash-based Message Authentication Code
		- same like SHA but input is TEXT & SECRET KEY then output hash
		- SECRET KEY change - same text but different output


HMAC-SHA-256 - input text and secret key but the hashing algo is SHA 256
*/
