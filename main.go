package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go_learn/route"
)

//func init() {
//	common.DB.AutoMigrate(&model.User{}, &model.Admin{})
//}

func main() {
	router := route.InitRoute()
	port := viper.GetString("server.port")
	fmt.Println("当前端口", port)
	if port != "" {
		router.Run(":" + port)
	} else {
		router.Run()
	}
}