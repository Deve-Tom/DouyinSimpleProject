package router

import (
	"DouyinSimpleProject/controller"
	"DouyinSimpleProject/middleware"
	"DouyinSimpleProject/service"

	"github.com/gin-gonic/gin"
)

var (
	userService = service.NewUserService()

	userController = controller.NewUserController(userService)
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// change to `gin.ReleaseMode` in production.
	gin.SetMode(gin.DebugMode)

	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/register/", userController.Register)
	apiRouter.POST("/user/login/", userController.Login)
	apiRouter.GET("/user/", middleware.JWTMiddleware(), userController.UserInfo)

	return r
}
