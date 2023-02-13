package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
)

// UserService serves as register, login and user info
type UserService interface {
	// auth
	Login(username, password string) *dto.AuthDTO
	CreateUser(username, password string) *dto.AuthDTO
	FindByUsername(username string) *entity.User
	ComparePassword(oldPassword, newPassword string) bool
	// user info
	FindUserByID(id uint) *entity.User
	GetUserInfo(id uint) *dto.UserInfoDTO
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Login(username, password string) *dto.AuthDTO {
	// 1. check whether the user already exists with username
	user := s.FindByUsername(username)
	if user == nil {
		return nil
	}
	// 2. compare passwords
	if !s.ComparePassword(user.Password, password) {
		return nil
	}
	// 3. generate token according username
	token, _ := utils.GenToken(user.ID)
	return &dto.AuthDTO{
		UserID: user.ID,
		Token:  token,
	}
}

func (s *userService) CreateUser(username, password string) *dto.AuthDTO {
	if user := s.FindByUsername(username); user != nil {
		return nil
	}
	u := dao.Q.User
	newUser := entity.User{
		Username: username, Password: password, Nickname: username,
	}
	_ = u.Create(&newUser)
	token, _ := utils.GenToken(newUser.ID)
	return &dto.AuthDTO{
		UserID: newUser.ID,
		Token:  token,
	}
}

func (s *userService) FindByUsername(username string) *entity.User {
	u := dao.Q.User
	user, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		return nil
	}
	return user
}

func (s *userService) ComparePassword(oldPassword, newPassword string) bool {
	return oldPassword == newPassword
}

func (s *userService) GetUserInfo(id uint) *dto.UserInfoDTO {
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

func (s *userService) FindUserByID(id uint) *entity.User {
	u := dao.Q.User
	user, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return user
}
