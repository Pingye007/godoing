package route

import (
	"net/http"

	"github.com/Pingye007/godoing/service"
)

func InitRoutes() {
	http.HandleFunc("/api/"+service.ApiVersion+"/"+service.ApiGet, service.Get)
	http.HandleFunc("/api/"+service.ApiVersion+"/"+service.ApiAdd, service.Add)
	http.HandleFunc("/api/"+service.ApiVersion+"/"+service.ApiUpdate, service.Update)
	http.HandleFunc("/api/"+service.ApiVersion+"/"+service.ApiDelete, service.Delete)
}
