package service

import (
	"errors"
	"github.com/Pingye007/godoing/db"
	"net/http"
	"strconv"
	"time"
)

func UserService(goReq *GoRequest) (*GoResponse, error) {
	if goReq == nil {
		return nil, errors.New("Nil Pointer")
	}
	var data any
	var err error
	res := GoResponse{
		Code:      http.StatusOK,
		Msg:       "Success",
		Timestamp: time.Now().Unix(),
	}
	switch goReq.Cmd {
	case CmdGetUser:
		id, _ := strconv.Atoi(goReq.Params[KeyUserId])
		data, err = db.QueryUserById(id)
	case CmdGetAllUser:
	case CmdCountUser:
	}

	if err != nil {
		return nil, err
	}
	res.Data = data
	return &res, nil
}
