package api

type (
	Result struct {
		Ok   int         `json:"ok"`
		Msg  string      `json:"msg,omitempty"`
		Data interface{} `json:"data,omitempty"`
	}
)

func Fail(msg string) (res *Result) {
	return &Result{Ok: 1, Msg: msg}
}

func Succ(data interface{}) (res *Result) {
	return &Result{
		Ok:   0,
		Data: data,
	}
}

func Response(b bool, date interface{}) (res *Result) {
	if b {
		res = Succ(date)
	} else {
		res = Fail(date.(string))
	}
	return
}
