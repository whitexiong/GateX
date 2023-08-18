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

// Claims 自定义的 payload 数据
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 根据用户名生成一个JWT令牌
func (j *JWT) GenerateToken(username string) (string, error) {
	// 创建一个新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.Expiry).Unix(),
			Issuer:    j.Issuer,
		},
	})

	// 使用 SecretKey 进行签名并获得完整编码的字符串令牌
	return token.SignedString([]byte(j.SecretKey))
}

// ParseToken 解析JWT令牌并返回用户名
func (j *JWT) ParseToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", fmt.Errorf("invalid token")
}
