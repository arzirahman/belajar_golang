package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var accessTokenPrivate []byte
var accessTokenPublic []byte

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	accessTokenPrivatePath := os.Getenv("ACCESS_TOKEN_PRIVATE_PATH")
	accessTokenPublicPath := os.Getenv("ACCESS_TOKEN_PUBLIC_PATH")

	var err error
	accessTokenPrivate, err = ioutil.ReadFile(accessTokenPrivatePath)
	if err != nil {
		panic("Error reading access token private key: " + err.Error())
	}

	accessTokenPublic, err = ioutil.ReadFile(accessTokenPublicPath)
	if err != nil {
		panic("Error reading access token public key: " + err.Error())
	}
}

func GenerateAccessToken(claims jwt.MapClaims) string {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(accessTokenPrivate)
	if err != nil {
		fmt.Println("create: parse key: %w", err)
	}
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Error generate access token: " + err.Error())
	}
	return tokenString
}

func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(accessTokenPublic)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}
	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}
	return claims, nil
}
