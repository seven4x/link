package account

type LoginResponse struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}
type (
	Login struct {
		Username string
		Password string
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
