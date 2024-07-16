package middlewares

import (
	"github.com/cnh1en/lets_go/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "valid-token" {
			ctx.Next()
		}

		response.ErrorResponse(ctx, 403, "token invalid")
		ctx.Abort()
	}
}
