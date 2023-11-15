package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Pingye007/godoing/config"
	"github.com/Pingye007/godoing/log"
)

func DatabaseService(rw http.ResponseWriter, req *http.Request) {

}

func ApiService(rw http.ResponseWriter, req *http.Request) {
	log.Log.Infoln("route url:", req.URL.Path)
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

	target := req.URL.Path[idx+len("/"):]
	goReq, err := ParseRequest(req)
	if err != nil {
		log.Log.Infoln("parse body data to request failed")
		Failure(rw, "Wrong Format Body Data")
		return
	}

	log.Log.Infoln("request:", goReq.String())

	var goRes *GoResponse
	switch target {
	case UrlUser:
		goRes, err = UserService(goReq)
	case UrlDoing:
		goRes, err = DoingService(goReq)
	case UrlResult:
		goRes, err = ResultService(goReq)
	default:
		Failure(rw, "Wrong URL Target")
		return
	}

	if err != nil {
		Failure(rw, "Service Error")
		return
	}

	Success(rw, goRes)
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
