package main

import (
	"Todolist/config"
	"Todolist/pkg/utils"
	"Todolist/repository/cache"
	"Todolist/repository/db/dao"
	"Todolist/router"
)

func init() {
	config.InitConfig()
	dao.MySQLInit()
	utils.InitLogger()
	cache.InitRedis()
}
func main() {
	r := router.NewRouter()
	_ = r.Run(config.Cfg.System.HttpPort)
}
