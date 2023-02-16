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

func GetUIDFromToken(ctx *gin.Context) uint {
	rawID, _ := ctx.Get("user_id")
	uid, _ := rawID.(uint)
	return uid
}
