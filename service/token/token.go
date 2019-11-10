package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = "secret"

// New builds new token
func New(payload map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"nbf":     time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}

// Parse checks for valid JWT token
func Parse(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("can't parse token, error: %s", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payload, _ := claims["payload"].(map[string]interface{})
		return payload, nil
	}

	return nil, fmt.Errorf("can't parse error")
}
