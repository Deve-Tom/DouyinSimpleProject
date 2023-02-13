package controller

type Response struct {
	StatusCode uint8  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempy"`
}

func NewResponse(code uint8, msg string) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  msg,
	}
}
