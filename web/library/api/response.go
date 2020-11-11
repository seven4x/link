package api

type (
	Result struct {
		Ok        bool              `json:"ok"`
		Data      interface{}       `json:"data,omitempty"`
		MsgId     string            `json:"msgId,omitempty"`
		Msg       string            `json:"msg,omitempty"`
		ErrorData map[string]string `json:"errorData,omitempty"`
	}
)

func Fail(msg string) (res *Result) {
	return &Result{Ok: false, Msg: msg}
}

func Success(data interface{}) (res *Result) {
	return &Result{Ok: true, Data: data}
}

func Response(date interface{}, svrErr *Err) (res *Result) {
	if svrErr == nil {
		res = &Result{
			Ok:   true,
			Data: date,
		}
	} else {
		if svrErr.Data != nil {
			return &Result{Ok: false, MsgId: svrErr.MsgId, ErrorData: svrErr.Data}
		} else {
			return &Result{Ok: false, MsgId: svrErr.MsgId}
		}
	}

	return
}
