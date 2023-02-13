package middleware

import (
	"DouyinSimpleProject/controller"
	"DouyinSimpleProject/utils"
	"net/http"
	"time"

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
		// no such user
		if tokenString == "" {
			ctx.JSON(http.StatusOK, controller.Response{
				StatusCode: 1,
				StatusMsg:  "No such user",
			})
			ctx.Abort()
			return
		}
		// validate token
		claims, ok := utils.ParseToken(tokenString)
		if !ok {
			ctx.JSON(http.StatusOK, controller.Response{
				StatusCode: 1,
				StatusMsg:  "Incorrect token",
			})
			ctx.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			ctx.JSON(http.StatusOK, controller.Response{
				StatusCode: 1,
				StatusMsg:  "Expired token",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
