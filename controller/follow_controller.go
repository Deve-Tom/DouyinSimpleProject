package controller

import (
	"DouyinSimpleProject/service"
	"DouyinSimpleProject/utils"

	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followService service.FollowService
}

func NewFollowController(followService service.FollowService) FollowController {
	return FollowController{followService}
}

func (c *FollowController) Action(ctx *gin.Context) {
	fuid, err := utils.String2uint(ctx.Query("to_user_id"))
	if err != nil {
		ErrorResponse(ctx, "invalid video_id")
		return
	}
	actionType, err := utils.String2uint(ctx.Query("action_type"))
	if err != nil {
		ErrorResponse(ctx, "invalid action_type")
		return
	}
	uid := GetUID(ctx)

	err = c.followService.Action(uid, fuid, actionType)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		SuccessResponseWithoutData(ctx, "Successfully do this action")
	}

}
