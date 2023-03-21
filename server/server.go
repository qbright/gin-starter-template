package server

import (
	"gin-starter-template/config"
)

func Init() {
	config := config.GetConfig()
	r := Router()
	r.Run("0.0.0.0:" + config.GetString("server.port"))
}
