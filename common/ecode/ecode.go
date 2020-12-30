package ecode

import "fmt"

type ECode int64

var codeList = map[int64]string{}

// 创建一个错误码
func New(code int64, msg string) ECode {
	if code <= 0 {
		panic(fmt.Sprintf("illegal error code: %d", code))
	}
	return add(code, msg)
}

// 添加一个错误码
func add(code int64, msg string) ECode {
	if _, ok := codeList[code]; ok {
		panic(fmt.Sprintf("error code: %d already exist", code))
	}
	codeList[code] = msg

	return ECode(code)
}

// 通过code获取一个存在的验证码
func Build(code int64) ECode {
	if _, ok := codeList[code]; !ok {
		panic(fmt.Sprintf("error code: %d don't exist", code))
	}
	return ECode(code)
}

func (e ECode) Code() int64 {
	return int64(e)
}

func (e ECode) Msg() string {
	if _, ok := codeList[e.Code()]; !ok {
		panic(fmt.Sprintf("error code: %d don't exist", e.Code()))
	}
	return codeList[e.Code()]
}
