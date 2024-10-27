package middleware

import (
	"api-gateway/internal/http/token"
	rdb "api-gateway/internal/pkg/redis"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(rdb *rdb.RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		valid, err := token.ValidateToken(authHeader)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		accessToken := authHeader

		isBlacklisted, err := rdb.DB.Exists(context.Background(), accessToken).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify token", "details": err.Error()})
			c.Abort()
			return
		}
		if isBlacklisted == 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in, please login again"})
			c.Abort()
			return
		}

		valid, err = token.ValidateToken(accessToken)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		claims, err := token.ExtractClaim(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims", "details": err.Error()})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
