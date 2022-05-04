package main

import (
	"fmt"
	"gin-web/data"
	"gin-web/data/model"
	"gin-web/user/api"
	"gin-web/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	var user []model.UserBase
	data.UserDb.First(&user)
	fmt.Println(user)
	for v, k := range user {
		fmt.Println(v, k)
	}
	uid := utils.GetConfig().Get("mysql.uid")
	fmt.Println(uid)
	uid = utils.GetConfig().Get("mysql.uid")
	fmt.Println(uid)
	app := gin.Default()
	admin := app.Group(fmt.Sprintf("/api/%v", utils.GetConfig().Get("app")))
	admin.Use()
	{
		admin.POST("login", api.Login)
	}
	app.Run(fmt.Sprintf(":%d", utils.GetConfig().Get("port")))
}
