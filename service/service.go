package service

import (
	"encoding/json"
	"fmt"
	"github.com/Pingye007/godoing/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pingye007/godoing/config"
	"github.com/Pingye007/godoing/log"
)

type GoRequest struct {
}

type GoResponse struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func InternalError(rw http.ResponseWriter, msg string) {
	res := GoResponse{
		Data: nil,
		Code: http.StatusInternalServerError,
		Msg:  msg,
	}

	b, err := json.Marshal(&res)
	if err != nil {
		panic(err.Error())
	}

	_, err = rw.Write(b)
	if err != nil {
		panic(err.Error())
	}
}

func Success(rw http.ResponseWriter, data any) {
	res := GoResponse{
		Data: data,
		Code: http.StatusOK,
		Msg:  "Success",
	}

	b, err := json.Marshal(&res)
	if err != nil {
		panic(err.Error())
	}

	_, err = rw.Write(b)
	if err != nil {
		panic(err.Error())
	}
}

func Failure(rw http.ResponseWriter) {
	InternalError(rw, "Failure")
}

func DatabaseService(rw http.ResponseWriter, req *http.Request) {

}

func UserService(rw http.ResponseWriter, req *http.Request) {
	key := "user_id"
	err := req.ParseForm()
	if err != nil {
		log.Log.Errorf("parse %s form failed \n", req.RequestURI)
		InternalError(rw, "Form data error")
		return
	}

	// Do the service
	idStr := req.Form.Get(key)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		InternalError(rw, "Id conversion error")
		return
	}

	user, err := db.QueryUserById(id)
	if err != nil {
		InternalError(rw, err.Error())
		return
	}

	Success(rw, user)
}

func StartServer() {

	addrStr := fmt.Sprintf("%s:%d", config.Cfg.Server.Ip, config.Cfg.Server.Port)
	addrStr = strings.ToLower(addrStr)
	server := &http.Server{
		Addr: addrStr,
	}

	log.Log.Infoln("start server done")
	server.ListenAndServe()
}
