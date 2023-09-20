package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWT struct {
	SecretKey string
	Issuer    string
	Expiry    time.Duration
}

type Claims struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (j *JWT) GenerateToken(username string, userID int64, role string) (string, error) {
	// 创建一个新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: username,
		UserID:   userID,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.Expiry).Unix(),
			Issuer:    j.Issuer,
		},
	})

	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWT) ParseToken(tokenStr string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
