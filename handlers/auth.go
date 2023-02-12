package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

// require API key to add new recipe
func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") !=
			os.Getenv("X_API_KEY") {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}
