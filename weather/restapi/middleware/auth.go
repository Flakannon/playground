package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	API_KEY = "SECURE"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		suppliedKey := ctx.GetHeader("X-API-KEY")

		if suppliedKey != API_KEY {
			ctx.String(http.StatusUnauthorized, "api key is invalid")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
