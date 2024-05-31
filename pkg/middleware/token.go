package middleware

import (
	"fmt"
	"time"

	"github.com/dewciu/timetrove_api/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateToken(user_id uuid.UUID) (string, error) {
	config, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	token_hour_lifetime := config.Server.TokenHourLifetime

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_hour_lifetime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Server.ApiSecret))
}

func TokenValid(c *gin.Context) bool {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		config, err := config.GetConfig()
		if err != nil {
			return 0, err
		}
		return []byte(config.Server.ApiSecret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.GetHeader("Authorization")
	fmt.Println(bearerToken)
	if bearerToken != "" {
		return bearerToken
	}
	return ""
}

func ExtractUserIDFromToken(c *gin.Context) (string, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		config, err := config.GetConfig()
		if err != nil {
			return 0, err
		}
		return []byte(config.Server.ApiSecret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user_id := claims["user_id"].(string)
		return user_id, nil
	}
	return "", nil
}
