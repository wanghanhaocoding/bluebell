package models

// 定义请求的参数结构体
// 注册
type ParamsSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
