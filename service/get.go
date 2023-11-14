package service

import (
	"net/http"
	"strings"

	"github.com/Pingye007/godoing/log"
)

func Get(rw http.ResponseWriter, req *http.Request) {
	log.Log.Infoln("url:", req.URL.Path)
	idx := strings.LastIndexAny(req.URL.Path, "/")
	if idx == -1 {
		log.Log.Infoln("parse request url failed")
		Failure(rw, "Illegal URL")
		return
	} else if idx == len(req.URL.Path)-1 {
		log.Log.Infoln("invalid url")
		Failure(rw, "Invalid URL")
		return
	}

	realCmd := req.URL.Path[idx+1:]
	switch realCmd {
	case User:
		UserService(rw, req)
	case Doing:
	case Result:
	}

}
