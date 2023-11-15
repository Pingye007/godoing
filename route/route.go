package route

import (
	"net/http"

	"github.com/Pingye007/godoing/service"
)

func InitRoutes() {
	http.HandleFunc("/api/"+service.ApiVersion+"/", service.ApiService)
}
