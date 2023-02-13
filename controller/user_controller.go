package controller

import (
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"DouyinSimpleProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userInfoService service.UserInfoService
}

func NewUserController(userInfoService service.UserInfoService) UserController {
	return UserController{
		userInfoService: userInfoService,
	}
}

func (c *UserController) UserInfo(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	userInfoDTO := c.userInfoService.GetUserInfo(utils.String2uint(userID))
	if userInfoDTO == nil {
		ctx.JSON(http.StatusOK, dto.UserInfoResponse{
			Response: dto.Response{
				StatusCode: 1,
				StatusMsg:  "No such user",
			},
		})
	} else {
		ctx.JSON(http.StatusOK, dto.UserInfoResponse{
			Response: dto.Response{
				StatusCode: 0,
				StatusMsg:  "Sucessfully get user info",
			},
			UserInfoDTO: *userInfoDTO,
		})
	}
}
