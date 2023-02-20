package dto

import "DouyinSimpleProject/entity"

type AuthDTO struct {
	UserID uint   `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

// AuthResponse responses to `/user/register/` or `/user/login/`
type AuthResponse struct {
	Response
	AuthDTO
}

type UserInfoDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	WorkCount     uint   `json:"work_count"`
	FavoriteCount uint   `json:"favorite_count"`
}

// UserInfoResponse responses to `/user/`
type UserInfoResponse struct {
	Response
	UserInfoDTO `json:"user"`
}

func NewUserInfoDTO(user *entity.User, loginUID uint) *UserInfoDTO {
	// TODO: get isFollow with loginUID
	isFollow := false
	if user.ID == loginUID {
		isFollow = true
	}
	return &UserInfoDTO{
		ID:            user.ID,
		Name:          user.Nickname,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
		WorkCount:     user.WorkCount,
		FavoriteCount: user.FavoriteCount,
	}
}
