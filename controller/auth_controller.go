package controller

import (
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

// Register handles `/user/register/`
func (c *AuthController) Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	// get AuthDTO
	authDTO := c.authService.CreateUser(username, password)
	if authDTO == nil {
		ctx.JSON(http.StatusOK, dto.AuthResponse{
			Response: dto.Response{StatusCode: 1, StatusMsg: "The user already exists"},
		})
	} else {
		ctx.JSON(http.StatusOK, dto.AuthResponse{
			Response: dto.Response{StatusCode: 0, StatusMsg: "Successfully Register"},
			AuthDTO:  *authDTO,
		})
	}
}

// Login handles `/user/login/`
func (c *AuthController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	if authDTO := c.authService.Login(username, password); authDTO == nil {
		ctx.JSON(http.StatusOK, dto.AuthResponse{
			Response: dto.Response{StatusCode: 0, StatusMsg: "No such user"},
		})

	} else {
		ctx.JSON(http.StatusOK, dto.AuthResponse{
			Response: dto.Response{StatusCode: 0, StatusMsg: "Successfully Login"},
			AuthDTO:  *authDTO,
		})
	}
}
