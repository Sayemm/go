package main

import (
	"crypto/sha256"
	"fmt"
)

// import (
// 	"ecommerce/cmd"
// )

func main() {
	// cmd.Serve()

	data := []byte("Hello")
	hash := sha256.Sum256(data)
	fmt.Println(data, hash)

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
*/
