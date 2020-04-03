package util

import (
	"bytes"
	"encoding/json"
)

type Message struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
func Deal(data interface{}) string {
	msg:=Message{}
	if data==nil {
		msg.Code=0
		msg.Message="失败"
		msg.Data=data
	}else {
		msg.Code=1
		msg.Message="成功"
		msg.Data=data
	}
	st, _ :=JSONMarshal(msg)
	return string(st)
}
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}