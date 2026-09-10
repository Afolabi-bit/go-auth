package middleware

import (
	"go-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserIDKey = "auth.userId"
	ctxRoleKey   = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format. expected 'Bearer <token>"})
			return
		}

		scheme := strings.ToLower(strings.TrimSpace(parts[0]))
		tokenStr := parts[1]

		if scheme != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format. expected 'Bearer <token>"})
			return
		}

		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			return
		}

		claims, err := auth.ParseToken(jwtSecret, tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token:" + err.Error()})
			return
		}

		c.Set(ctxUserIDKey, claims.Subject)
		c.Set(ctxRoleKey, claims.Role)

		c.Next()
	}
}

func GetUserID(c *gin.Context) (string, bool) {
	val, ok := c.Get(ctxUserIDKey)
	if !ok {
		return "", false
	}

	userId, ok := val.(string)
	return userId, ok
}

func GetRole(c *gin.Context) (string, bool) {
	val, ok := c.Get(ctxRoleKey)
	if !ok {
		return "", false
	}

	role, ok := val.(string)
	return role, ok
}
