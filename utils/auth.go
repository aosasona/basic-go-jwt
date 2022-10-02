package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func GenerateJWT(uuid string) (string, error) {

	claims := jwt.MapClaims{
		"uuid": uuid,
		"exp":  time.Now().Add(30 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		fmt.Printf("something went wrong: %s", err.Error())
		return "", errors.New("unable to generate token")
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ReadUserUUID(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("authorization header is required")
	}

	token := header[7:]
	claims, err := ParseJWT(token)

	if err != nil {
		return "", errors.New("invalid token")
	}

	return claims["uuid"].(string), nil
}
