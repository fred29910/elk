package res

import (
	"encoding/json"
	"github.com/masseelch/elk/internal/ov"
)

func SuccessRes(data any) string {
	rsp := ov.Response{
		Code: 0,
		Data: data,
	}
	return marshal(rsp)
}

func marshal(rsp ov.Response) string {
	bs, _ := json.Marshal(rsp)
	return string(bs)
}

func FailRes(msg string, code int) string {
	rsp := ov.Response{
		Code: code,
		Msg:  msg,
	}
	return marshal(rsp)
}
