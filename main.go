package main

import (
	"encoding/base64"
	"fmt"
)

// import (
// 	"ecommerce/cmd"
// )

func main() {
	// cmd.Serve()

	var s string
	s = "ab"
	byteArr := []byte(s)
	fmt.Println(s, byteArr) // ab [97 98]

	enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	b64str := enc.EncodeToString(byteArr)
	fmt.Println(b64str) //YWI

	decodedStr, _ := enc.DecodeString(b64str)
	fmt.Println(decodedStr) // [97 98]
}

/*
JWT (JSON Web Token) - Authentication
======================================
BASE64 - Method for encoding binary data into an ASCII string format (A-Z, a-z, 0-9, +, /)
       - coverts binary data (images, files, bytes) into a set of 64 printable ASCII characters.

Purpose of BASE64
-----------------
- When transferring data from one system to another system base64 is the faster way

*/
