package controller

import (
	"DouyinSimpleProject/service"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
	videoService service.VideoService
}

func NewVideoController(videoService service.VideoService) VideoController {
	return VideoController{
		videoService: videoService,
	}
}

func (c *VideoController) PublishVideo(ctx *gin.Context) {
	//////// 1. Get parameters
	// we already set user_id in *JWT middleware*
	rawID, _ := ctx.Get("user_id")
	uid, _ := rawID.(uint)

	// get title and video from form-data
	title := ctx.PostForm("title")
	form, err := ctx.MultipartForm()
	if err != nil {
		ErrorResponse(ctx, err.Error())
		return
	}
	// only support single-file upload
	videoFile := form.File["data"][0]

	msg, ok := c.videoService.Publish(ctx, uid, title, videoFile)
	if !ok {
		ErrorResponse(ctx, msg)
	} else {
		SuccessResponseWithoutData(ctx, msg)
	}
}
