package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
	"errors"
)

// UserService serves as register, login and user info.
type UserService interface {
	// auth
	Login(username, password string) (*dto.AuthDTO, error)
	CreateUser(username, password string) (*dto.AuthDTO, error)
	// user info
	GetUserInfo(id uint) (*dto.UserInfoDTO, error)

	findUserByUsername(username string) (*entity.User, error)
	comparePassword(oldPassword, newPassword string) bool
	findUserByID(id uint) (*entity.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Login(username, password string) (*dto.AuthDTO, error) {
	// 1. check whether the user already exists with username
	user, err := s.findUserByUsername(username)
	if err != nil {
		return nil, errors.New("incorrect username or password")
	}
	// 2. compare passwords
	if !s.comparePassword(user.Password, password) {
		return nil, errors.New("incorrect username or password")
	}
	// 3. generate token according username
	token, _ := utils.GenToken(user.ID)

	return &dto.AuthDTO{UserID: user.ID, Token: token}, nil
}

func (s *userService) CreateUser(username, password string) (*dto.AuthDTO, error) {
	if _, err := s.findUserByUsername(username); err == nil {
		return nil, errors.New("username already exists")
	}
	u := dao.Q.User
	newUser := entity.User{
		Username: username, Password: password, Nickname: username,
	}
	_ = u.Create(&newUser)
	token, _ := utils.GenToken(newUser.ID)
	return &dto.AuthDTO{UserID: newUser.ID, Token: token}, nil
}

func (s *userService) findUserByUsername(username string) (*entity.User, error) {
	u := dao.Q.User
	user, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// comparePassword just simply compare the uploaded password with the old password.
// Maybe we can encrypt the password.
func (s *userService) comparePassword(oldPassword, newPassword string) bool {
	return oldPassword == newPassword
}

func (s *userService) GetUserInfo(id uint) (*dto.UserInfoDTO, error) {
	user, err := s.findUserByID(id)
	if err != nil {
		return nil, err
	}
	// TODO
	isFollow := true
	userInfoDTO := &dto.UserInfoDTO{
		ID:            id,
		Name:          user.Nickname,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
	return userInfoDTO, nil
}

func (s *userService) findUserByID(id uint) (*entity.User, error) {
	u := dao.Q.User
	user, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, errors.New("the user not exists")
	}
	return user, nil
}
