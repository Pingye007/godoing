package main

import (
	"github.com/Pingye007/godoing/log"
	"github.com/Pingye007/godoing/service"
)

func main() {
	log.Log.Debugln("godoing start to run...")
	service.StartServer()
}
