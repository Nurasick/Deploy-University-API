package jwt

import (
	"errors"
	"time"
	"university/config"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(config.Load().JWTSecret)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
	jwt.StandardClaims
}

func GenerateToken(userID, roleID int, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RoleID: roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateToken(tknStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tknStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
