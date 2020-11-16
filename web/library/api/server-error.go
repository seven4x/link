package api

import "encoding/json"

type Err struct {
	MsgId string
	Data  map[string]string
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
func NewError(msg string) *Err {
	return &Err{
		MsgId: msg,
	}
}
func NewErrorWithData(msg string, key1, value1 string) *Err {
	return &Err{
		MsgId: msg,
		Data:  map[string]string{key1: value1},
	}
}
