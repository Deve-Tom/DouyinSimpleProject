package dto

// AuthDTO is a response for `/user/register/` and `/user/login/`
type AuthDTO struct {
	UserID uint   `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

// NewAuthDTO creates a new instance of AuthDTO
func NewAuthDTO(uid uint, token string) AuthDTO {
	return AuthDTO{
		UserID: uid,
		Token:  token,
	}
}
