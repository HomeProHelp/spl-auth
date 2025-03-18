package token

import (
	"fmt"
	"github/LissaiDev/spl-auth/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte(utils.GetEnv("JWT_SECRET", "spl-auth"))
)

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id string) (*string, string) {
	claims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "spl-auth",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secretKey)

	if err != nil {
		return nil, utils.AuthenticationCodes["internal_server_error"]
	}

	return &signed, utils.AuthenticationCodes["success"]
}

func ValidateToken(tokenString string) (*Claims, string) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, utils.AuthenticationCodes["invalid_token"]
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, utils.AuthenticationCodes["success"]
	}

	return nil, utils.AuthenticationCodes["invalid_token"]
}

func RefreshToken(tokenString string) (string, string) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", utils.AuthenticationCodes["invalid_token"]
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signed, err := newToken.SignedString(secretKey)
		if err != nil {
			return "", utils.AuthenticationCodes["internal_server_error"]
		}
		return signed, utils.AuthenticationCodes["success"]
	}

	return "", utils.AuthenticationCodes["invalid_token"]
}
