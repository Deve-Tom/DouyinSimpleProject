package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
)

type UserInfoService interface {
	FindUserByID(id uint) *entity.User
	GetUserInfo(id uint) *dto.UserInfoDTO
}

type userInfoService struct {
}

func NewUserInfoService() UserInfoService {
	return &userInfoService{}
}

func (s *userInfoService) GetUserInfo(id uint) *dto.UserInfoDTO {
	user := s.FindUserByID(id)
	if user == nil {
		return nil
	}
	// TODO
	isFollow := true
	return &dto.UserInfoDTO{
		ID:            id,
		Name:          user.Nickname,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
}

func (s *userInfoService) FindUserByID(id uint) *entity.User {
	u := dao.Q.User
	user, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return user
}
