package service

import (
	"github.com/Pingye007/godoing/log"
	"github.com/Pingye007/godoing/route"
	"net/http"
	"strings"
)

func Get(rw http.ResponseWriter, req *http.Request) {
	log.Log.Infoln("url:", req.URL.Path)
	idx := strings.LastIndexAny(req.URL.Path, "/")
	if idx == -1 {
		log.Log.Infoln("parse request url failed")
		InternalError(rw, "Url illegal")
		return
	} else if idx == len(req.URL.Path)-1 {
		log.Log.Infoln("url invalid")
		InternalError(rw, "Url invalid")
		return
	}

	realCmd := req.URL.Path[idx+1:]
	if realCmd == route.User {

	}

}
