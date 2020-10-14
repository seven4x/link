package api

type (
	FailResp struct {
		Ok  int    `json:"ok"`
		Msg string `json:"msg"`
	}
	SuccResp struct {
		Ok   int `json:"ok"`
		Data interface{}
	}
)

func Fail(msg string) (res *FailResp) {
	return &FailResp{Ok: 1, Msg: msg}
}

func Succ(data interface{}) (res *SuccResp) {
	return &SuccResp{
		Ok:   0,
		Data: data,
	}
}
