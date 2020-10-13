package api

type (
	FailResp struct {
		OK  int
		Msg string
	}
	SuccResp struct {
		OK   int
		Data interface{}
	}
)

func Fail(msg string) (res *FailResp) {
	return &FailResp{OK: 1, Msg: msg}
}

func Succ(data interface{}) (res *SuccResp) {
	return &SuccResp{
		OK:   0,
		Data: data,
	}
}
