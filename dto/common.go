package dto

import "errors"

type Response struct {
	StatusCode uint8  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

const (
	MaxUsernameLength = 32 //Maximum length of user name
	MaxPasswordLength = 32 //Maximum password length
	MinPasswordLength = 6  //Minimum password length
)

var (
	ErrorUserNameNull    = errors.New("empty username")
	ErrorUserNameExtend  = errors.New("incorrect username length")
	ErrorPasswordNull    = errors.New("empty password")
	ErrorPasswordLength  = errors.New("incorrect password length")
	ErrorUserExit        = errors.New("username already exists")
	ErrorUserNotExit     = errors.New("the user not exists")
	ErrorFullPossibility = errors.New("the user not exists, incorrect username or password")
	ErrorPasswordFalse   = errors.New("incorrect username or password")
)
