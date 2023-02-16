package controller

import (
	"DouyinSimpleProject/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, dto.Response{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func SuccessResponseWithoutData(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, dto.Response{
		StatusCode: 0,
		StatusMsg:  msg,
	})
}

// GetUID get user id from token if the token is valid.
func GetUID(ctx *gin.Context) uint {
	rawID, _ := ctx.Get("user_id")
	uid, _ := rawID.(uint)
	return uid
}
