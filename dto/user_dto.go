package dto

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
}

type UserInfoResponse struct {
	Response
	UserInfoDTO
}
