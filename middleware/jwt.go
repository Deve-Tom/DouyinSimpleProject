package middleware

import (
	"DouyinSimpleProject/controller"
	"DouyinSimpleProject/utils"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware intercepts *Token* and parses `user_id`
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token string
		tokenString := ctx.Query("token")
		if tokenString == "" {
			tokenString = ctx.PostForm("token")
		}

		claims, msg := utils.ValidToken(tokenString)
		if claims == nil {
			controller.ErrorResponse(ctx, msg)
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
