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

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" { // Bearer, Token
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerParts[1]
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
		expectedSignature := createSignature(message, m.cnf.JwtSecretKey)

		// Match New Signature with the signature that we got from Frontend
		if expectedSignature != signature {
			http.Error(w, "Unauthorized-Hacker Alert", http.StatusUnauthorized)
			return
		}

		// if matches then we will proceed to next
		next.ServeHTTP(w, r)
	})
}

func createSignature(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	hash := h.Sum(nil)
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash)
}

func base64UrlEncoder(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
