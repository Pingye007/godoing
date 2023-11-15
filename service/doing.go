package service

import (
	"errors"
	"github.com/Pingye007/godoing/db"
	"net/http"
	"strconv"
	"time"
)

func DoingService(goReq *GoRequest) (*GoResponse, error) {
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
	case CmdGetDoing:
		id, _ := strconv.Atoi(goReq.Params[KeyUserId])
		data, err = db.QueryDoingById(id)
	case CmdGetAllDoing:
	case CmdGetAllDoingUser:
	case CmdCountDoingUser:
	case CmdCountDoing:
	}

	if err != nil {
		return nil, err
	}
	res.Data = data
	return &res, nil
}
