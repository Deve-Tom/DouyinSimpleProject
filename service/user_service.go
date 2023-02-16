package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
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

	IsUserLegal(userName string, passWord string) error
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Login(username, password string) (*dto.AuthDTO, error) {
	//0. Validity verification
	err := s.IsUserLegal(username, password)
	if err != nil {
		return nil, err
	}

	// 1. check whether the user already exists with username
	user, err := s.findUserByUsername(username)
	if err != nil {
		return nil, dto.ErrorFullPossibility
	}

	// 2. compare passwords
	if !s.comparePassword(user.Password, password) {
		return nil, dto.ErrorPasswordFalse
	}
	// 3. generate token according username
	token, _ := utils.GenToken(user.ID)

	return &dto.AuthDTO{UserID: user.ID, Token: token}, nil
}

func (s *userService) CreateUser(username, password string) (*dto.AuthDTO, error) {
	//0. Validity verification
	err := s.IsUserLegal(username, password)
	if err != nil {
		return nil, err
	}
	//1. check whether the user already exists with username
	if _, err := s.findUserByUsername(username); err == nil {
		return nil, dto.ErrorUserExit
	}

	//2. Password encryption
	newPassword, _ := utils.HashAndSalt(password)

	//3. Database storage
	u := dao.Q.User
	newUser := entity.User{
		Username: username, Password: newPassword, Nickname: username,
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
func (s *userService) comparePassword(oldPassword string, newPassword string) bool {
	return utils.ComparePassword(oldPassword, newPassword)
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
		return nil, dto.ErrorUserNotExit
	}
	return user, nil
}

// IsUserLegal Validity verification of username and password
func (s *userService) IsUserLegal(userName string, passWord string) error {
	//1.username
	if userName == "" {
		return dto.ErrorUserNameNull
	}
	if len(userName) > dto.MaxUsernameLength {
		return dto.ErrorUserNameExtend
	}
	//2.password
	if passWord == "" {
		return dto.ErrorPasswordNull
	}
	if len(passWord) > dto.MaxPasswordLength || len(passWord) < dto.MinPasswordLength {
		return dto.ErrorPasswordLength
	}
	return nil
}
