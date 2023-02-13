package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
)

type AuthService interface {
	Login(username, password string) *dto.AuthDTO
	CreateUser(username, password string) *dto.AuthDTO
	FindByUsername(username string) *entity.User
	ComparePassword(oldPassword, newPassword string) bool
}

type authService struct {
}

// NewAuthService is a instance of AuthService
func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Login(username, password string) *dto.AuthDTO {
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

func (s *authService) CreateUser(username, password string) *dto.AuthDTO {
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

func (s *authService) FindByUsername(username string) *entity.User {
	u := dao.Q.User
	user, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		return nil
	}
	return user
}

func (s *authService) ComparePassword(oldPassword, newPassword string) bool {
	return oldPassword == newPassword
}
