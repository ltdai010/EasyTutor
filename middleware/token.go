package middleware

import (
	"EasyTutor/consts"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaims struct {
	Username string
	TypeToken     string
	jwt.StandardClaims
}

var KeyFunc = func(token *jwt.Token) (interface{}, error) {
	return []byte(consts.SecretKey), nil
}

func GenerateToken(username string, typeToken string) (string, error) {
	claims := TokenClaims{
		Username:       username,
		TypeToken:      typeToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 360000,
			IssuedAt: time.Now().Unix(),
			Issuer:    "EasyTutor",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(consts.SecretKey))
}

//return the username and the usertype
func ValidateToken(tokenString string) (string, string) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, KeyFunc)
	if err != nil {
		return "", ""
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		if claim.Username != "" {
			return claim.Username, claim.TypeToken
		}
	}
	return "", ""
}