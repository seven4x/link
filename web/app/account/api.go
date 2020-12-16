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
