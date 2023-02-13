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
