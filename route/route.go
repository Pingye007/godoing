package route

import (
	"net/http"

	"github.com/Pingye007/godoing/service"
)

const (
	GetUser = "db/getUser"
)

func InitRoutes() {
	http.HandleFunc("/api/v1/"+GetUser, service.UserService)
}
