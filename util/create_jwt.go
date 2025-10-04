package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"` //user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// Header B64
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncoder(byteArrHeader)

	// Payload B64
	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncoder(byteArrData)

	// Signature b64
	byteArrSecret := []byte(secret)
	message := headerB64 + "." + payloadB64
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncoder(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64
	return jwt, nil
}

func base64UrlEncoder(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

/*
Header -  Hbase64 							(Secret)
				=> header.payload(b64) => HMAC-SHA256 => Hash(Signarute) => Hbase64=> Header.Payload.Hash(b64) = JWT
Payload - Pbase64

- Who has that secret (backend) only that  will know whether that is valid or not
- We don't keep sensitive information to jwt, the information that are okay to view we only keep those info


- So what's the point if everyone can see that?
=> if they make any changes it will be possible to figure out

Frontend (email, pass)  -> Backend
Backend (JWT) -> Frontend

- Frontend will decode payload (that's okay)
- Frontend will send other requests (/create) using that JWT
- Backend will use the header, payload and use it's secret key to generate hash
- if backend can match the hash that backend provied before than if will confirm that signature is correct
- if does not match then we will know that there are some changes on payload/header
	=> won't response to the request - unauthorized.
-
*/
