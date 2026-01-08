package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("==> Authentication Middleware")
		// Parse JWT
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 { // Bearer, Token
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
		byteArrSecret := []byte(m.cnf.JwtSecretKey)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEncoder(hash)

		// Match New Signature with the signature that we got from Frontend
		if newSignature != signature {
			http.Error(w, "Unauthorized-Hacker Alert", http.StatusUnauthorized)
			return
		}

		// if matches then we will proceed to next
		next.ServeHTTP(w, r)
	})
}

func base64UrlEncoder(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
