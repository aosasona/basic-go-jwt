package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
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
