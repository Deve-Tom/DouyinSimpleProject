package router

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/controller"
	"DouyinSimpleProject/middleware"
	"DouyinSimpleProject/service"

	"github.com/gin-gonic/gin"
)

var (
	userService  = service.NewUserService()
	videoService = service.NewVideoService()

	userController  = controller.NewUserController(userService)
	videoController = controller.NewVideoController(videoService)
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// change to `gin.ReleaseMode` in production.
	gin.SetMode(gin.DebugMode)

	// public directory is used to serve static resources
	r.Static("/static", config.STATIC_ROOT_PATH)

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/register/", userController.Register)
	apiRouter.POST("/user/login/", userController.Login)
	apiRouter.GET("/user/", middleware.JWTMiddleware(), userController.UserInfo)

	apiRouter.POST("/publish/action/", middleware.JWTMiddleware(), videoController.PublishVideo)
	apiRouter.GET("/publish/list/", middleware.JWTMiddleware(), videoController.ListVideo)

	return r
}
