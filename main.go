package main

import (
	"gin-starter-template/config"
	"gin-starter-template/db"
	"gin-starter-template/db/dbutil"
	"gin-starter-template/server"
	"gin-starter-template/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Init()

	gin.SetMode(config.GetConfig().GetString("mode"))
	utils.InitLogger()

	db.Init()

	dbutil.Migrate(db.GetDb())

	server.Init()

}
