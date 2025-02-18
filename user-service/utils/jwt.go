package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const secretKey = "secret"

func GenerateToken(email string, userId uuid.UUID) (string, error) {
	// Преобразуем UUID в int64
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId, // Сохраняем userId как строку
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (string, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if !parsedToken.Valid {
		return "", errors.New("token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	// Проверка наличия userId в claims
	userIdStr, exists := claims["userId"].(string)
	if !exists {
		return "", errors.New("userId claim is missing or not a string")
	}

	return userIdStr, nil
}
