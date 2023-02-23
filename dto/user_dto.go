package dto

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/entity"
)

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
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	FollowCount     uint   `json:"follow_count"`
	FollowerCount   uint   `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	WorkCount       uint   `json:"work_count"`
	FavoriteCount   uint   `json:"favorite_count"`
	TotalFavorCount uint   `json:"total_favorited"`
}

// UserInfoResponse responses to `/user/`
type UserInfoResponse struct {
	Response
	UserInfoDTO `json:"user"`
}

type FollowInfoResponse struct {
	Response
	UserList []*UserInfoDTO `json:"user_list"`
}


func NewUserInfoDTO(user *entity.User, loginUID uint) *UserInfoDTO {
	isFollow := false
	//login user + feed
	if user.ID == loginUID {
		isFollow = true
	} else {
		//login user + feed
		if user.ID == loginUID {
			isFollow = true
		}
		if loginUID != 0 {
			fq := dao.Q.Follow
			cnt, err := fq.Where(fq.UserID.Eq(loginUID)).Where(fq.FollowUserID.Eq(user.ID)).Count()
			if err != nil {
				return nil
			}
			if cnt != 0 {
				isFollow = true
			}
		}
	}

	if loginUID != 0 {
		fq := dao.Q.Follow
		cnt, err := fq.Where(fq.UserID.Eq(loginUID)).Where(fq.FollowUserID.Eq(user.ID)).Count()
		if err != nil {
			return nil
		}
		if cnt != 0 {
			isFollow = true
		}
	}


	return &UserInfoDTO{
		ID:              user.ID,
		Name:            user.Nickname,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        isFollow,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
		TotalFavorCount: user.TotalFavorCount,
	}
}
