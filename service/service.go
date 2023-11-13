package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Pingye007/godoing/config"
	"github.com/Pingye007/godoing/log"
)

func UserService(rw http.ResponseWriter, req *http.Request) {

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
