package controller

import (
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthResponse struct {
	Response
	dto.AuthDTO
}

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

// RegisterController handles `/user/register/`
func (c *AuthController) RegisterController(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	// get AuthDTO
	authDTO := c.authService.CreateUser(username, password)
	if authDTO == nil {
		ctx.JSON(http.StatusOK, AuthResponse{
			Response: NewResponse(1, "The user already exists"),
		})
	} else {
		ctx.JSON(http.StatusOK, AuthResponse{
			Response: NewResponse(
				0, "Successfully Register",
			),
			AuthDTO: *authDTO,
		})
	}
}

// LoginController handles `/user/login/`
func (c *AuthController) LoginController(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	if authDTO := c.authService.Login(username, password); authDTO == nil {
		ctx.JSON(http.StatusOK, AuthResponse{
			Response: NewResponse(
				0, "No such user",
			),
		})

	} else {
		ctx.JSON(http.StatusOK, AuthResponse{
			Response: NewResponse(
				0, "Successfully Login",
			),
			AuthDTO: *authDTO,
		})
	}
}
