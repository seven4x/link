package response

type LoginResponse struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}
