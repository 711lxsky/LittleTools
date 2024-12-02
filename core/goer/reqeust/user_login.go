package reqeust

type UserLoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
