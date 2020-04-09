// @Author:WY
// @Time:2020/4/7 13:38
package main

import (
	"gin_vue/dao"
	"gin_vue/models"
	"gin_vue/routers"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	defer dao.CloseDB()
	//绑定模型
	dao.DB.AutoMigrate(&models.User{},&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}


