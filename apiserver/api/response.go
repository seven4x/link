package api

import "github.com/seven4x/link/app"

type (
	SimpleResult struct {
		Ok   bool        `json:"ok"`
		Data interface{} `json:"data,omitempty"`
	}
	PageResult struct {
		Ok   bool        `json:"ok"`
		Data interface{} `json:"data,omitempty"`
		Page Page        `json:"page,omitempty"`
	}

	ErrorResult struct {
		Ok        bool              `json:"ok"`
		MsgId     string            `json:"msgId,omitempty"`
		Msg       string            `json:"msg,omitempty"`
		ErrorData map[string]string `json:"errorData,omitempty"`
	}
	Page struct {
		HasMore bool `json:"hasMore"`
		Total   int  `json:"total,omitempty"`
		NextId  int  `json:"nextId"`
	}
)

func Fail(msg string) (res interface{}) {
	return &ErrorResult{Ok: false, Msg: msg}
}

func FailMsgId(msgId string) (res interface{}) {
	return &ErrorResult{Ok: false, MsgId: msgId}
}

func Success(data interface{}) (res interface{}) {
	return &SimpleResult{Ok: true, Data: data}
}

func ResponseHasMore(data interface{}, hasMore bool) (res interface{}) {
	res = &PageResult{
		Ok:   true,
		Data: data,
		Page: Page{HasMore: hasMore},
	}
	return res
}

func ResponsePage(data interface{}, hasMore bool, next int) (res interface{}) {

	res = &PageResult{
		Ok:   true,
		Data: data,
		Page: Page{
			HasMore: hasMore,
			NextId:  next,
		},
	}

	return
}

func Response(date interface{}, svrErr *app.Err) (res interface{}) {
	if svrErr == nil {
		res = &SimpleResult{
			Ok:   true,
			Data: date,
		}
	} else {
		if svrErr.Data != nil {
			return &ErrorResult{Ok: false, MsgId: svrErr.MsgId, ErrorData: svrErr.Data}
		} else {
			return &ErrorResult{Ok: false, MsgId: svrErr.MsgId}
		}
	}

	return
}
