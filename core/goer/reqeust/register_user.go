package reqeust

type RegisterUserRequest struct {
	UserName        string `json:"userName"`
	RawPassword     string `json:"rawPassword"`
	ConfirmPassword string `json:"confirmPassword"` // 确认密码
}
