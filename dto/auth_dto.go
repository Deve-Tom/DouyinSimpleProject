package dto

// AuthDTO is a response for `/user/register/` and `/user/login/`
type AuthDTO struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

// NewAuthDTO creates a new instance of AuthDTO
func NewAuthDTO(uid uint, token string) AuthDTO {
	return AuthDTO{
		UserID: uid,
		Token:  token,
	}
}
