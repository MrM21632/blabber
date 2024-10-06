package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func GenerateAccessToken(user_id uuid.UUID) (string, error) {
	token_lifespan := GetJWTLifespan()
	token_secret, err := base64.StdEncoding.DecodeString(GetJWTSecret())
	if err != nil {
		log.Error("Failed to retrieve and decode JWT secret, cannot generate new access token")
		return "", err
	}

	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user_id.String(),
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
	}

	// Using HMAC for simplicity, demo purposes. In a real application, we should
	// utilize an asymmetrical algorithm like RS512 or ES512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(token_secret)
}

func GenerateRefreshToken() (string, error) {
	token_secret, err := base64.StdEncoding.DecodeString(GetJWTSecret())
	if err != nil {
		log.Error("Failed to retrieve and decode JWT secret, cannot generate new refresh token")
		return "", err
	}

	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	}

	// Using HMAC for simplicity, demo purposes. In a real application, we should
	// utilize an asymmetrical algorithm like RS512 or ES512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(token_secret)
}

func ValidateToken(context *gin.Context) error {
	token_string := ExtractToken(context)
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret, err := base64.StdEncoding.DecodeString(GetJWTSecret())
		if err != nil {
			return nil, err
		}
		return secret, nil
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("token is invalid")
	}
	return nil
}

func ExtractToken(context *gin.Context) string {
	token := context.Query("token")
	if token != "" {
		return token
	}

	bearer_token := context.Request.Header.Get("Authorization")
	chunks := strings.Split(bearer_token, " ")
	if len(chunks) == 2 {
		return chunks[1]
	}

	return ""
}
