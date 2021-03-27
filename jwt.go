package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createtowken(email string) (string, error) {

	// the link expire in five days
	tim := time.Now()
	tim = tim.AddDate(0, 0, 5)
	t := tim.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"expirein": t,
	})
	tokenString, err := token.SignedString([]byte("notsecret"))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil

}

func verifytoken(tokenstr string) (string, int64, error) {

	cl := jwt.MapClaims{}
	var mail string
	var expire int64

	token, err := jwt.ParseWithClaims(tokenstr, cl, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("notsecret"), nil
	})

	if err != nil {
		return mail, expire, err
	}

	if !token.Valid {

		err := errors.New("token not valid")
		return mail, expire, err
	}

	email, ok := cl["email"].(string)
	if !ok {
		return mail, expire, err

	}

	experin, ok := cl["expirein"].(float64)
	if !ok {
		return mail, expire, err

	}
	expire = int64(experin)

	return email, expire, nil
}
