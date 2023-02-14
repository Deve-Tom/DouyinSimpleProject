package controller

import (
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"DouyinSimpleProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

// Register handles `/user/register/`
func (c *UserController) Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	// get AuthDTO
	authDTO, err := c.userService.CreateUser(username, password)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		ctx.JSON(http.StatusOK, dto.AuthResponse{
			Response: dto.Response{StatusCode: 0, StatusMsg: "Successfully Register"},
			AuthDTO:  *authDTO,
		})
	}
}

// Login handles `/user/login/`
func (c *UserController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	if authDTO, err := c.userService.Login(username, password); err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		ctx.JSON(http.StatusOK, dto.AuthResponse{
			Response: dto.Response{StatusCode: 0, StatusMsg: "Successfully Login"},
			AuthDTO:  *authDTO,
		})
	}
}

// UserInfo handles `/user/`
func (c *UserController) UserInfo(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	userInfoDTO, err := c.userService.GetUserInfo(utils.String2uint(userID))
	if err != nil {
		ErrorResponse(ctx, err.Error())
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
