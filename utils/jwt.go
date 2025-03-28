package utils

import (
	"bank-api/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.JWTSecret)

func GenerateToken(username string) (map[string]interface{}, error) {
	var jwtExpiration = time.Hour * 24 // 24 hours expiration time
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(jwtExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	returnedData := map[string]interface{}{
		"token":      signedToken,
		"expired_at": time.Now().Add(jwtExpiration).Format(time.RFC3339),
	}

	return returnedData, nil
}

type TokenValidationResult struct {
	Claims jwt.MapClaims
	Token  string
}

func ValidateToken(tokenStr string) (TokenValidationResult, error) {
	removedBearer := RemoveBearerPrefix(tokenStr)

	token, err := jwt.Parse(removedBearer, func(t *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return TokenValidationResult{}, errors.New("Invalid token (parse)" + err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return TokenValidationResult{
			Claims: claims,
			Token:  removedBearer,
		}, nil
	}

	return TokenValidationResult{}, errors.New("Invalid token (validateToken2)")
}

func RemoveBearerPrefix(token string) string {
	if len(token) > 7 && token[:7] == "Bearer " {
		return token[7:]
	}

	return token
}
