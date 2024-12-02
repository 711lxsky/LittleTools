package reqeust

type UserRegisterRequest struct {
	UserName        string `json:"userName"`
	RawPassword     string `json:"rawPassword"`
	ConfirmPassword string `json:"confirmPassword"` // 确认密码
}
