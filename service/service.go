package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pingye007/godoing/db"

	"github.com/Pingye007/godoing/config"
	"github.com/Pingye007/godoing/log"
)

func UserService(rw http.ResponseWriter, req *http.Request) {
	var greq GoRequest
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Log.Infoln("read body failed")
		Failure(rw, "Invalid Body")
		return
	}
	defer req.Body.Close()
	log.Log.Infoln("body:", body)

	err = json.Unmarshal(body, &greq)
	if err != nil {
		log.Log.Infoln("unmarshel failed")
		Failure(rw, "Invalid Data")
		return
	}

	switch greq.Cmd {
	case CmdGetUser:
		id, _ := strconv.Atoi(greq.Params[KeyUserId])
		user, err := db.QueryUserById(id)
		if err != nil {
			Failure(rw, err.Error())
			return
		}
		Success(rw, user)

	case CmdGetAllUser:
	case CmdCountUser:
	}

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
