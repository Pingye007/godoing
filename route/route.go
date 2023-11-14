package route

import (
	"net/http"

	"github.com/Pingye007/godoing/service"
)

const (
	ApiVersion = "v1"
	ApiGet     = "get"
	ApiAdd     = "add"
	ApiUpdate  = "update"
	ApiDelete  = "delete"
	User       = "user"
)

func InitRoutes() {
	http.HandleFunc("/api/"+ApiVersion+"/"+ApiGet, service.Get)
	http.HandleFunc("/api/"+ApiVersion+"/"+ApiAdd, service.Add)
	http.HandleFunc("/api/"+ApiVersion+"/"+ApiUpdate, service.Update)
	http.HandleFunc("/api/"+ApiVersion+"/"+ApiDelete, service.Delete)
	http.HandleFunc("/api/"+ApiVersion+"/", service.DatabaseService)
}
