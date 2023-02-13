package router

import (
	"DouyinSimpleProject/controller"
	"DouyinSimpleProject/service"

	"github.com/gin-gonic/gin"
)

var (
	authService    = service.NewAuthService()
	authController = controller.NewAuthController(authService)
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// change to `gin.ReleaseMode` in production.
	gin.SetMode(gin.DebugMode)

	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/register/", authController.RegisterController)
	apiRouter.POST("/user/login/", authController.LoginController)

	return r
}
