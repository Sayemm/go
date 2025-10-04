package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse JWT
	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	headerArr := strings.Split(header, " ")
	if len(headerArr) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken := headerArr[1]
	tokenParts := strings.Split(accessToken, ".")
	if len(tokenParts) != 3 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	signature := tokenParts[2]

	// Create Signature/Hash from Header, Payload and Secret Key
	message := jwtHeader + "." + jwtPayload
	byteArrMessage := []byte(message)
	byteArrSecret := []byte(config.GetConfig().JwtSecretKey)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	hash := h.Sum(nil)
	newSignature := base64UrlEncoder(hash)

	// Match New Signature with the signature that we got from Frontend
	if newSignature != signature {
		http.Error(w, "Unauthorized-Hacker Alert", http.StatusUnauthorized)
		return
	}

	// if matches then we will create product

	var newProd database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProd)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Give me valid json", 400)
		return
	}

	createdProd := database.Store(newProd)

	util.SendData(w, createdProd, 201)
}

func base64UrlEncoder(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
