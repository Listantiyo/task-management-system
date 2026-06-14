package jwt

import (
	apperr "task-management-system/internal/delivery/error"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secret = "just-harcode-secret"

type CustomClaim struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type User struct {
	ID       uuid.UUID
	Email    string
	Username string
}

type UserPayload interface {
	GetID() string
	GetEmail() string
	GetUsername() string
}

func GenerateJWT(userPayload UserPayload) (tokenString string, err error) {
	claims := CustomClaim{
		UserID:   userPayload.GetID(),
		Email:    userPayload.GetEmail(),
		Username: userPayload.GetUsername(),
		RegisteredClaims: jwt.RegisteredClaims{

			Issuer:    "task-management-system",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(tokenStr string) (*CustomClaim, error) {
	var claims CustomClaim
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperr.New(apperr.ErrTokenInvalid, "invalid token signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, apperr.New(apperr.ErrTokenInvalid, "invalid token")
	}

	return &claims, nil
}
