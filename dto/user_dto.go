package dto

// AuthDTO is a response for `/user/register/` and `/user/login/`
type AuthDTO struct {
	UserID uint   `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

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
