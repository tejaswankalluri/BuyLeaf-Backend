package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func JwtSignIn[subT string | uint | interface{}](sub subT, Time time.Duration) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour * Time).Unix(),
	})
	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString, err
}

func JwtValid(tokenString string) (ok bool, err error, claims jwt.MapClaims) {
	//	Validate the token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return false, errors.New("not a Valid token"), nil
		}
		return true, nil, claims
	} else {
		return false, errors.New("not a Valid token"), nil
	}
}
