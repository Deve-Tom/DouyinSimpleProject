package dto

type Response struct {
	StatusCode uint8  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
