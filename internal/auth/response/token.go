package response

type Token struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}
