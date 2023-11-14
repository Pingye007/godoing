package service

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	ApiVersion = "v1"
	ApiGet     = "get"
	ApiAdd     = "add"
	ApiUpdate  = "update"
	ApiDelete  = "delete"
	User       = "user"
	Doing      = "doing"
	Result     = "result"

	CmdGetUser    = "getUser"
	CmdGetAllUser = "getAllUser"
	CmdCountUser  = "countUser"

	KeyUserId = "user_id"
)

type GoRequest struct {
	Cmd    string            `json:"cmd"`
	Params map[string]string `json:"params"`
}

type GoResponse struct {
	Data      any    `json:"data"`
	Code      int    `json:"code"`
	Msg       string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func commonWrite(rw http.ResponseWriter, res *GoResponse) {
	b, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	_, err = rw.Write(b)
	if err != nil {
		panic(err.Error())
	}
}

func Success(rw http.ResponseWriter, data any) {
	res := &GoResponse{
		Data:      data,
		Code:      http.StatusOK,
		Msg:       "Success",
		Timestamp: time.Now().Unix(),
	}

	commonWrite(rw, res)
}

func Failure(rw http.ResponseWriter, msg string) {
	res := &GoResponse{
		Data:      nil,
		Code:      http.StatusInternalServerError,
		Msg:       msg,
		Timestamp: time.Now().Unix(),
	}

	commonWrite(rw, res)
}
