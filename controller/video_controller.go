package controller

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"DouyinSimpleProject/utils"
	"net/http"
	"time"

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

// Feed handles `/feed/`
func (c *VideoController) Feed(ctx *gin.Context) {
	// get latest_time from request
	rawLatestTime := ctx.Query("latest_time")
	var latestTime time.Time
	if rawLatestTime == "" {
		latestTime = time.Now()
	} else {
		latestTime = utils.UnmarshalJSTimeStamp(rawLatestTime)
	}
	// get user_id from request
	var uid uint = 0
	tokenString := ctx.Query("token")
	if tokenString != "" {
		claims, err := utils.ValidToken(tokenString)
		if err != nil {
			ErrorResponse(ctx, err.Error())
			return
		}
		uid = claims.UserID
	}

	videoDTOs, err := c.videoService.GetVideoDTOList(config.VIDEO_LIMIT, latestTime, uid, true)
	if err != nil {
		ErrorResponse(ctx, err.Error())
		return
	}

	if len(videoDTOs) == 0 {
		SuccessResponseWithoutData(ctx, "Ah, no any videos")
		return
	}
	nextTime := utils.MarshalJSTimeStamp(videoDTOs[len(videoDTOs)-1].CreatedAt)
	ctx.JSON(http.StatusOK, dto.VideoResponse{
		Response: dto.Response{
			StatusCode: 0,
			StatusMsg:  "Successfully fetch videos",
		},
		NextTime:  nextTime,
		VideoList: videoDTOs,
	})

}

// PublishVideo handles `/publish/action/`
func (c *VideoController) PublishVideo(ctx *gin.Context) {
	//////// 1. Get parameters
	// we already set user_id in *JWT middleware*
	uid := GetUID(ctx)

	// get title and video from form-data
	title := ctx.PostForm("title")
	form, err := ctx.MultipartForm()
	if err != nil {
		ErrorResponse(ctx, err.Error())
		return
	}
	// only support single-file upload
	videoFile := form.File["data"][0]

	err = c.videoService.Publish(ctx, uid, title, videoFile)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		SuccessResponseWithoutData(ctx, "sucessfully publish video")
	}
}

// ListVideo handles `/publish/list/`
func (c *VideoController) ListVideo(ctx *gin.Context) {
	uid, err := utils.String2uint(ctx.Query("user_id"))
	if err != nil {
		ErrorResponse(ctx, "invalid user_id")
		return
	}
	videoDTOs, err := c.videoService.GetVideoDTOList(-1, time.Now(), uid, false)
	if err != nil {
		ErrorResponse(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, dto.VideoResponse{
		Response: dto.Response{
			StatusCode: 0,
			StatusMsg:  "Successfully get video list",
		},
		VideoList: videoDTOs,
	})
}
