package router

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/controller"
	"DouyinSimpleProject/middleware"
	"DouyinSimpleProject/service"

	"github.com/gin-gonic/gin"
)

var (
	userService     = service.NewUserService()
	videoService    = service.NewVideoService()
	favoriteService = service.NewFavoriteService()
	commentService  = service.NewCommentService()
	followService   = service.NewFollowService()

	userController     = controller.NewUserController(userService)
	videoController    = controller.NewVideoController(videoService)
	favoriteController = controller.NewFavoriteController(favoriteService)
	commentController  = controller.NewCommentController(commentService)
	followController   = controller.NewFollowController(followService)
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
	apiRouter.GET("/feed/", videoController.Feed)

	apiRouter.POST("/favorite/action/", middleware.JWTMiddleware(), favoriteController.Action)
	apiRouter.GET("/favorite/list/", middleware.JWTMiddleware(), favoriteController.FavoriteList)

	apiRouter.POST("/comment/action/", middleware.JWTMiddleware(), commentController.Action)
	apiRouter.GET("/comment/list/", commentController.List)

	apiRouter.POST("/relation/action/", middleware.JWTMiddleware(), followController.Action)
	apiRouter.GET("/relation/follow/list/", middleware.JWTMiddleware(), followController.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.JWTMiddleware(), followController.FollowerList)
	return r
}
