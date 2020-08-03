package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

var (
	adminUsername = "admin"
	adminPassword = "admin"
	jwtSecret     = "iu4fcn0qnua"
)

func AuthenticateUser(r *http.Request) (string, error) {
	tokenString, err := ParseToken(r)
	if err != nil {
		return "", err
	}
	token, err := VerifyToken(tokenString)
	if err != nil {
		return "", err
	}
	return ExtractUsername(token)
}

func ParseToken(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	tokenArr := strings.Split(tokenString, " ")
	if len(tokenArr) != 2 {
		err := errors.New("invalid token")
		return "", err
	}
	return tokenArr[1], nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractUsername(token *jwt.Token) (string, error) {
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "", errors.New("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("Invalid token")
		}
		return username, nil
	}
	return "", nil
}

func CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["admin"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := at.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}