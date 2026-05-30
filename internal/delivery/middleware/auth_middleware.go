package middleware

import (
	"net/http"
	"strings"
	"task-management-system/internal/utils/dto"
	"task-management-system/internal/utils/jwt"
	"task-management-system/internal/utils/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.GetHeader("Authorization")
		if authHeader == "" {
			response.AbortFailed(g, http.StatusUnauthorized, dto.ErrUnauthorized, "Header Authorization not found!")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if authHeader == tokenString {
			response.AbortFailed(g, http.StatusUnauthorized, dto.ErrUnauthorized, "Invalid Header Token")
			return
		}

		claims, err := jwt.ValidateJWT(tokenString)
		if err != nil {
			if custErr, ok := dto.AsType[dto.ErrorCode](err); ok {
				response.AbortFailed(g, http.StatusUnauthorized, custErr, "Invalid Token")
				return
			}
			response.AbortFailed(g, http.StatusInternalServerError, dto.ErrInternalServer, "Failed process token")
			return
		}

		g.Set("user", claims)

		g.Next()
	}
}
