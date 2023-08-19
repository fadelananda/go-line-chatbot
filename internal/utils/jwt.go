package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: make this generic :>
func GenerateOauthJWT(lineId string) (string, error) {
	claims := jwt.MapClaims{
		"line_id": lineId,
	}
	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["line_id"] = lineId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenString, err
}

func GenerateJWT[T jwt.Claims](claims T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenString, err
}

func ValidateJWT[T jwt.Claims](tokenString string) (T, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	var zeroClaim T

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretKey), nil
	})

	if err != nil {
		return zeroClaim, err
	}

	if claims, ok := token.Claims.(T); ok && token.Valid {
		return claims, nil
	} else {
		return zeroClaim, err
	}
}

// TODO: also make this generic :>
func ValidateOauthJWT(tokenString string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
