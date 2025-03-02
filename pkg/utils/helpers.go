package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your-secret-key") // Change this to a strong secret key

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(data any, datakey string, exipresIn int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		datakey: data,
		"exp":   exipresIn,
		// "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	return token.SignedString(secretKey)
}

func VerifyJWT(tokenString string, datakey string) (interface{}, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check that the signing method is HMAC and that it matches what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("invalid signing")
			return nil, errors.New("invalid request")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		log.Println("invalid token")
		return nil, errors.New("invalid token")
	}

	// Get the claims and check expiration
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if data, exists := claims[datakey]; exists {
			// Check expiration
			if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
				return nil, errors.New("token expired")
			}
			return data, nil
		}
	}

	return nil, errors.New("data key not found")
}

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func GetBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host
}
