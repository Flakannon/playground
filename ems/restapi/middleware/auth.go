package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegisteredInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input UserRegisteredInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Example validation logic here (this should be replaced with your actual user authentication logic)
		if input.Username == "" || input.Password == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			ctx.Abort()
			return
		}

		// Add the username to the context after successful authentication
		ctx.Set("username", input.Username)

		// Continue down the middleware chain
		ctx.Next()
	}
}
