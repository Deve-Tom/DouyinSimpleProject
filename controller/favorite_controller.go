package controller

import (
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"DouyinSimpleProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct {
	favoriteService service.FavoriteService
}

func NewFavoriteController(favoriteService service.FavoriteService) FavoriteController {
	return FavoriteController{favoriteService}
}

func (c *FavoriteController) Action(ctx *gin.Context) {
	vid, err := utils.String2uint(ctx.Query("video_id"))
	if err != nil {
		ErrorResponse(ctx, "invalid video_id")
		return
	}
	actionType, err := utils.String2uint(ctx.Query("action_type"))
	if err != nil {
		ErrorResponse(ctx, "invalid action_type")
		return
	}
	rawID, _ := ctx.Get("user_id")
	uid, _ := rawID.(uint)

	err = c.favoriteService.Action(uid, vid, actionType)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		SuccessResponseWithoutData(ctx, "Successfully do this action")
	}

}

func (c *FavoriteController) FavoriteList(ctx *gin.Context) {
	uid, err := utils.String2uint(ctx.Query("user_id"))
	if err != nil {
		ErrorResponse(ctx, "invalid user_id")
	}
	videoDTOList, err := c.favoriteService.GetFavoriteList(uid)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		ctx.JSON(http.StatusOK, dto.VideoResponse{
			Response: dto.Response{
				StatusCode: 0,
				StatusMsg:  "Successfully get favorite list",
			},
			VideoList: videoDTOList,
		})
	}
}
