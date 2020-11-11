package api

type (
	Page struct {
		HasMore bool
		Total   int
	}
	Result struct {
		Ok        bool              `json:"ok"`
		Data      interface{}       `json:"data,omitempty"`
		MsgId     string            `json:"msgId,omitempty"`
		Msg       string            `json:"msg,omitempty"`
		ErrorData map[string]string `json:"errorData,omitempty"`
		Page      Page              `json:"page,omitempty"`
	}
)

func Fail(msg string) (res *Result) {
	return &Result{Ok: false, Msg: msg}
}

func Success(data interface{}) (res *Result) {
	return &Result{Ok: true, Data: data}
}

func ResponseHasMore(data interface{}, hasMore bool) (res *Result) {
	res = &Result{
		Ok:   true,
		Data: data,
		Page: Page{HasMore: hasMore},
	}

	return res
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
