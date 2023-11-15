package service

import (
	"encoding/json"
	"errors"
	"github.com/Pingye007/godoing/log"
	"io"
	"net/http"
	"time"
)

const (
	ApiVersion          = "v1"
	UrlUser             = "user"
	UrlDoing            = "doing"
	UrlResult           = "result"
	CmdGetUser          = "getUserById"
	CmdGetAllUser       = "getAllUser"
	CmdCountUser        = "countUser"
	CmdGetDoing         = "getDoingById"
	CmdGetAllDoing      = "getAllDoing"
	CmdGetAllDoingUser  = "getAllDoingOfUser"
	CmdCountDoingUser   = "countDoingOfUser"
	CmdCountDoing       = "countDoing"
	CmdGetResult        = "getResultById"
	CmdGetAllResult     = "getAllResult"
	CmdGetAllResultUser = "getAllResultOfUser"
	CmdCountResultUser  = "countResultOfUser"
	CmdCountResult      = "countResult"

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
	if res == nil {
		log.Log.Errorln("nil pointer")
		return
	}
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

func ParseRequest(req *http.Request) (*GoRequest, error) {
	if req == nil {
		return nil, errors.New("Nil Pointer")
	}

	var grep GoRequest
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Log.Infoln("read body failed")
		return nil, errors.New("Invalid Body")
	}
	defer req.Body.Close()
	log.Log.Infoln("body:", body)

	err = json.Unmarshal(body, &grep)
	if err != nil {
		log.Log.Infoln("unmarshel failed")
		return nil, errors.New("Invalid Data")
	}

	return &grep, nil
}

func (req GoRequest) String() string {
	str := "cmd:" + req.Cmd + ",params:["
	for _, p := range req.Params {
		str = str + p + ","
	}
	str = str + "]"
	return str
}
