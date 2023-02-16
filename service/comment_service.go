package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CommentService interface {
	Action(uid, vid, actionType uint, content, rawCommentID string) (*dto.CommentDTO, error)
	DO(uid, vid uint, content string) (*dto.CommentDTO, error)
	Delete(uid, vid, commentID uint) (*dto.CommentDTO, error)
	List(uid, vid uint) ([]*dto.CommentDTO, error)
}

type commentService struct{}

func NewCommentService() CommentService {
	return &commentService{}
}

func (s *commentService) Action(uid, vid, actionType uint, content, rawCommentID string) (*dto.CommentDTO, error) {
	if actionType == 1 {
		return s.DO(uid, vid, content)
	} else if actionType == 2 {
		commentID, err := utils.String2uint(rawCommentID)
		if err != nil {
			return nil, err
		}
		return s.Delete(uid, vid, commentID)
	} else {
		return nil, errors.New("invalid action type")
	}
}

// DO submits a comment
func (s *commentService) DO(uid, vid uint, content string) (*dto.CommentDTO, error) {
	now := time.Now()
	var commentEntity = &entity.Comment{
		Model: gorm.Model{
			CreatedAt: now,
			UpdatedAt: now,
		},
		Content: content,
		UserID:  uid,
		VideoID: vid,
	}
	err := dao.Q.Transaction(func(tx *dao.Query) error {
		cq := tx.Comment
		vq := tx.Video
		// insert comment
		err := cq.Create(commentEntity)
		if err != nil {
			return err
		}
		// video.comment_count + 1
		_, err = vq.Where(vq.ID.Eq(vid)).UpdateSimple(vq.CommentCount.Add(1))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	uq := dao.Q.User
	user, err := uq.Where(uq.ID.Eq(uid)).First()
	if err != nil {
		return nil, err
	}
	userInfoDTO := dto.NewUserInfoDTO(user, uid)
	commentDTO := dto.NewCommentDTO(userInfoDTO, commentEntity)
	return commentDTO, nil
}

// Delete deletes a comment
func (s *commentService) Delete(uid, vid, commentID uint) (*dto.CommentDTO, error) {
	var commentEntity *entity.Comment
	err := dao.Q.Transaction(func(tx *dao.Query) error {
		cq := tx.Comment
		// find comment according to uid, vid, and comment id
		comment, err := cq.Preload(cq.User).Where(cq.ID.Eq(commentID)).Where(cq.UserID.Eq(uid)).Where(cq.VideoID.Eq(vid)).First()
		if err != nil {
			return err
		}
		commentEntity = comment
		// delete comment
		_, err = cq.Unscoped().Delete(comment)
		if err != nil {
			return err
		}
		// video.comment_cout - 1
		vq := tx.Video
		_, err = vq.Where(vq.ID.Eq(vid)).UpdateSimple(vq.CommentCount.Sub(1))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	uq := dao.Q.User
	user, err := uq.Where(uq.ID.Eq(uid)).First()
	if err != nil {
		return nil, err
	}
	userInfoDTO := dto.NewUserInfoDTO(user, uid)
	commentDTO := dto.NewCommentDTO(userInfoDTO, commentEntity)
	return commentDTO, nil
}

func (s *commentService) List(uid, vid uint) ([]*dto.CommentDTO, error) {
	cq := dao.Q.Comment
	comments, err := cq.Preload(cq.User).Where(cq.UserID.Eq(uid)).Where(cq.VideoID.Eq(vid)).Find()
	if err != nil {
		return nil, err
	}
	commentDTOList := make([]*dto.CommentDTO, len(comments))
	for i, comment := range comments {
		userInfoDTO := dto.NewUserInfoDTO(&comment.User, uid)
		commentDTOList[i] = dto.NewCommentDTO(userInfoDTO, comment)
	}
	return commentDTOList, nil
}
