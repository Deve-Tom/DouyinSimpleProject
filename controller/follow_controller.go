package controller

import (
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"DouyinSimpleProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followService service.FollowService
}

func NewFollowController(followService service.FollowService) FollowController {
	return FollowController{followService}
}

// Action handles `/relation/action/`
func (c *FollowController) Action(ctx *gin.Context) {
	fuid, err := utils.String2uint(ctx.Query("to_user_id"))
	if err != nil {
		ErrorResponse(ctx, "invalid user_id")
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

// FollowList handles `/relation/follow/list`
func (c *FollowController) FollowList(ctx *gin.Context) {
	uid := GetUID(ctx)
	userDTOs, err := c.followService.GetFollowList(uid)
	if err != nil {
		ErrorResponse(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, dto.FollowInfoResponse{
		Response: dto.Response{
			StatusCode: 0,
			StatusMsg:  "Successfully get follow list",
		},
		UserList: userDTOs,
	})
}
