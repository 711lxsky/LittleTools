package reqeust

type UserLoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	UserName        string `json:"userName"`
	Email           string `json:"email"`
	RawPassword     string `json:"rawPassword"`
	ConfirmPassword string `json:"confirmPassword"` // 确认密码
}

type UserUpdateSelfInfoRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
