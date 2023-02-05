package app

import "encoding/json"

type Err struct {
	MsgId string
	Msg   string
	Data  map[string]string
}

func (e Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
func NewError(msgId string) Err {
	return Err{
		MsgId: msgId,
	}
}
func NewMsgError(msg string) *Err {
	return &Err{
		Msg: msg,
	}
}

func NewErrorWithData(msg string, key1, value1 string) *Err {
	return &Err{
		MsgId: msg,
		Data:  map[string]string{key1: value1},
	}
}
