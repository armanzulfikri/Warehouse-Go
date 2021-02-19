package response

//LoginResponse ...
type LoginResponse struct {
	Code   uint   `json:"code"`
	Status string `json:"status"`
	Token  string `json:"token"`
}
