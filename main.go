package main

import (
	"Todolist/config"
	"Todolist/pkg/utils"
	"Todolist/repository/db/dao"
	"Todolist/repository/db/model"
	"Todolist/router"
	"context"
	"fmt"
)

func init() {
	config.InitConfig()
	dao.MySQLInit()
	utils.InitLogger()
}
func main() {

	fmt.Printf("%#v \n", *config.Cfg.System)
	fmt.Printf("%#v \n", *config.Cfg.MySql["default"])
	fmt.Printf("%#v \n", *config.Cfg.Redis)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("dao test")
	userDao := dao.NewUserDao(context.TODO())
	newUser := &model.User{
		UserName:       "zhangsan",
		PasswordDigest: "12893721iu3hiuoasiof3hro",
	}
	userDao.CreateUser(newUser)

	r := router.NewRouter()
	_ = r.Run(config.Cfg.System.HttpPort)
}
