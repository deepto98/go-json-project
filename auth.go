package main

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v4"
)

//Middleware to protect authenticated routes - Decorator Pattern
func withJWTAuth(handlerFunc http.HandlerFunc, store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling JWT Auth Middleware")

		tokenString := r.Header.Get("x-jwt-token")
		token, err := validateJWT(tokenString)
		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{err.Error()})
			return
		}

		if !token.Valid {
			WriteJSON(w, http.StatusForbidden, ApiError{err.Error()})
			return
		}
		userID, err := getID(r)
		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{err.Error()})
			return
		}

		account, err := store.GetAccountById(userID)
		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{err.Error()})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(int64(claims["accountNumber"].(float64)))
		fmt.Println(account.Number)

		if account.Number != int64(claims["accountNumber"].(float64)) {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "Invalid JWT"})
			return
		}

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	// From Go docs:s
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

}

func createJWT(account *Account) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"expiresAt":     15000,
		"accountNumber": account.Number,
	}
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
