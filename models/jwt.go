package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWT_KEY = []byte("3gm2bV237FVbMi3rmk32407FH9y7CAaI8yi2rI408yuIR")


type Claims struct {
	Id          int64
	User        *User `json:"payload"`
	jwt.StandardClaims
}

func getExpirationJWT() time.Time {
	// Expiration, tomorrow at same hour
	now := time.Now()
	yyyy, mm, dd := now.Date()
	expiration := time.Date(yyyy, mm, dd+1, now.Hour(), 0, 0, 0, now.Location())
	return expiration
}

func (user *User) GenerateTokenJWT() (tokenStr string, exp time.Time, err error) {
	exp = getExpirationJWT()
	claims := Claims{
		Id:          exp.Unix(),
		User:        user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			Issuer:    "sapip.mx",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(JWT_KEY)
	if err != nil {
		return
	}
	return
}

func ValidateTokenJWT(tokenStr string, obj string, act string) (claims Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("no autorizado - %v", err.Error())
		}
		return JWT_KEY, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = fmt.Errorf("sesion no valida - %v", err.Error())
		return
	}
	return
}

