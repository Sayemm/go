package main

import (
	"ecommerce/util"
	"fmt"
)

// import (
// 	"ecommerce/cmd"
// )

func main() {
	// cmd.Serve()

	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         45,
		FirstName:   "Seer",
		LastName:    "Sayem",
		Email:       "sayemseer@gmail.com",
		IsShopOwner: false,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
}

/*
JWT (JSON Web Token) - Authentication
======================================
BASE64 - Method for encoding binary data into an ASCII string format (A-Z, a-z, 0-9, +, /)
       - coverts binary data (images, files, bytes) into a set of 64 printable ASCII characters.
	   - data -> byte -> base64

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


------------------
JWT - 3 Parts (jwt.io) - all 3 parts are encoded to base64 then joined using dot(.)
	= Header
	= Payload/Claim
	= Signature


Header
------
- algorithm
- type - JWT

Payload/Claim
-------------
- data that we send i.e. frontend
- data (email, pass, ...)

Signature
---------

*/
