package dto

import (
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
)

type CommentDTO struct {
	ID         uint        `json:"id"`
	User       UserInfoDTO `json:"user"`
	Content    string      `json:"content"`
	CreateDate string      `json:"create_date"`
}

type CommentResponse struct {
	Response
	Comment     *CommentDTO   `json:"comment,omitempty"`
	CommentList []*CommentDTO `json:"comment_list,omitempty"`
}

func NewCommentDTO(user *UserInfoDTO, comment *entity.Comment) *CommentDTO {
	createDate := utils.GetCommentReponseDate(comment.CreatedAt)
	return &CommentDTO{
		ID:         comment.ID,
		User:       *user,
		Content:    comment.Content,
		CreateDate: createDate,
	}
}
