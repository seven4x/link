package user

type LoginResponse struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}
type (
	Login struct {
		Username string
		Password string
	}
	AccountInfo struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		NickName string `json:"nickName"`
		Avatar   string `json:"avatar"`
	}
)

type WechatUserInfo struct {
	Openid     string
	Nickname   string
	Sex        int
	Province   string
	City       string
	Country    string
	Geadimgurl string
	Privilege  []string
	Unionid    string
}

type RegisterRequest struct {
	Code     string `validate:"required"`
	LoginId  string `validate:"required,min=4,max=32"`
	NickName string `validate:"required,min=2,max=32"`
	Password string `validate:"required,min=6"`
	CreateBy int
}
