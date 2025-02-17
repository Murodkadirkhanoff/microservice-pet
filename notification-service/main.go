package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 🔑 Секретный ключ (должен совпадать с Laravel Passport)
var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAm2QDKVyEAR38zzBtI8mB
KET6gc/SfLkXrS1NcFTX0LhiNZVSbP++DE/RPOVm9gUtMOXGN7qiCSxOR6nmNstY
cCvq1xIodCtD4C8ToOyntf6e8hpW9TGz4JNBzaU7g3F43rfF+Xe8sRnl+xaxKqjI
F076Co73tKZY3AeYNaIFZ5Lv2xMZpbqlwcU3j2Pj/GKM9xeZbFuFki/H5XI2INkx
Xr0Lsm32lwwb5vBl+m1rXqgqCQu2qYcqxiorlhIopi6BiXOBsV/1PTNqGW555/x9
VhVispFP7hV5HTtgHbiAkCwI9RC+JTul3oIfPfawGwxM6UpRIOMb7kqW6MKIw3Fj
C+VvSpJ/P/FYhbNVuko8JwJj2mvC96RrMJ9OAlz2w2ffaHRswRr4QESbHOm2I9HH
aoLTHYTZKzbp4UwUt4KT6zREi+A2MDN+XZ3UxgCbOiiLT+oEaECO2W6X2Lf5wKIc
bV1HBuLUybqSZ7UW8WZADtRKx0M0Qt0+cW+3ZntE6FiNTLqXUnE8x0GjgU50MSgi
sSrm5WPBWpEfAsqrjcW2axx3Q/KryzvwCsVVwhL3mtVsN24hDFfUp5Wf98SEPgbV
qjS2ZJdv5pnSwWfrE+EPfbuSVFhobtt7wPsLPpwEOoOwm2rnxFkhjZiO8Z716FBx
nkMotuGxvKK8HoK5AvlVBN0CAwEAAQ==
-----END PUBLIC KEY-----`)

// 📌 Функция валидации токена
func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Используем публичный ключ RSA вместо HMAC
		parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}
		return parsedKey, nil
	})
}

// 🚀 Middleware для проверки JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Нет токена"})
			c.Abort()
			return
		}

		// Проверяем формат "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный формат токена"})
			c.Abort()
			return
		}

		// Валидация токена
		token, err := validateToken(tokenParts[1])
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		// Если всё ок, продолжаем выполнение
		c.Next()
	}
}

func main() {
	// Создаём сервер Gin
	r := gin.Default()

	// Открытый маршрут (без авторизации)
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Открытый маршрут, авторизация не нужна"})
	})

	// Защищённый маршрут
	r.GET("/secure", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Доступ разрешён! Токен действителен."})
	})

	// Запускаем сервер
	log.Println("🚀 Сервер запущен на :8081")
	log.Fatal(r.Run(":8081"))
}
