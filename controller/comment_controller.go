package controller

import (
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/service"
	"DouyinSimpleProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return CommentController{commentService}
}

func (c *CommentController) Action(ctx *gin.Context) {
	uid := GetUID(ctx)
	vid, err := utils.String2uint(ctx.Query("video_id"))
	if err != nil {
		ErrorResponse(ctx, "invalid video id")
	}
	actionType, err := utils.String2uint(ctx.Query("action_type"))
	if err != nil {
		ErrorResponse(ctx, "invalid action type")
	}
	content := ctx.Query("comment_text")
	rawCommentID := ctx.Query("comment_id")
	commentDTO, err := c.commentService.Action(uid, vid, actionType, content, rawCommentID)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		ctx.JSON(http.StatusOK, dto.CommentResponse{
			Response: dto.Response{
				StatusCode: 0,
				StatusMsg:  "Successfully do this action",
			},
			Comment: commentDTO,
		})
	}
}

func (c *CommentController) List(ctx *gin.Context) {
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
	vid, err := utils.String2uint(ctx.Query("video_id"))
	if err != nil {
		ErrorResponse(ctx, err.Error())
	}
	commentDTOList, err := c.commentService.List(uid, vid)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	}
	ctx.JSON(http.StatusOK, dto.CommentResponse{
		Response: dto.Response{
			StatusCode: 0,
			StatusMsg:  "Successfully get the comment list",
		},
		CommentList: commentDTOList,
	})
}
