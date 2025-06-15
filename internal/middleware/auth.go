package middleware

import (
	"net/http"

	"golang_server/pkg/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware для проверки авторизации
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		token := utils.ExtractTokenFromHeader(authHeader)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Сохраняем информацию о пользователе в контексте
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}

// GetUserID извлекает ID пользователя из контекста
func GetUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	
	id, ok := userID.(uint)
	return id, ok
}

// GetUserEmail извлекает email пользователя из контекста
func GetUserEmail(c *gin.Context) (string, bool) {
	email, exists := c.Get("user_email")
	if !exists {
		return "", false
	}
	
	userEmail, ok := email.(string)
	return userEmail, ok
} 