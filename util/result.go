package util

import "encoding/json"

type Result struct {
	Code int         `json:"code" form:"code"`
	Msg  string      `json:"msg" form:"msg"`
	Data interface{} `json:"data" form:"data"`
}

func NewResult(code int, msg string, data ...interface{}) []byte {
	if len(data) > 0 {
		res := Result{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
		buf, _ := json.Marshal(res)
		return buf
	} else {
		res := Result{
			Code: code,
			Msg:  msg,
		}
		buf, _ := json.Marshal(res)
		return buf
	}
}
