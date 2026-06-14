package middleware

import (
	"net/http"
	"strings"
	apperr "task-management-system/internal/delivery/error"
	"task-management-system/internal/pkg/jwt"
	"task-management-system/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.GetHeader("Authorization")
		if authHeader == "" {
			response.AbortFailed(g, http.StatusUnauthorized, apperr.ErrUnauthorized, "Header Authorization not found!")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if authHeader == tokenString {
			response.AbortFailed(g, http.StatusUnauthorized, apperr.ErrUnauthorized, "Invalid Header Token")
			return
		}

		claims, err := jwt.ValidateJWT(tokenString)
		if err != nil {
			if custErr, ok := apperr.AsType[apperr.AppError](err); ok {
				response.AbortFailed(g, http.StatusUnauthorized, custErr.Code, custErr.Message)
				return
			}
			response.AbortFailed(g, http.StatusInternalServerError, apperr.ErrInternalServer, "Failed process token")
			return
		}

		g.Set("user", claims)

		g.Next()
	}
}
