package main

import (
	"fmt"
	"gin-web/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CustomRouterMiddle1(c *gin.Context) {
	t := time.Now()
	fmt.Println("我是自定义中间件第1种定义方式---请求之前")
	//在gin上下文中定义一个变量
	c.Set("example", "CustomRouterMiddle1")
	//请求之前
	c.Next()
	fmt.Println("我是自定义中间件第1种定义方式---请求之后")
	//请求之后
	//计算整个请求过程耗时
	t2 := time.Since(t)
	log.Println(t2)

}

func main() {

	uid := utils.GetConfig().Get("mysql.uid")
	fmt.Println(uid)
	uid = utils.GetConfig().Get("mysql.uid")
	fmt.Println(uid)
	app := gin.Default()

	user := app.Group("/user")
	user.Use(CustomRouterMiddle1)
	{
		user.GET("test", func(context *gin.Context) {
			context.String(http.StatusOK, "232")
		})
	}
	admin := app.Group("/admin")
	admin.Use()
	{
		admin.GET("test", func(context *gin.Context) {
			context.String(http.StatusOK, "232")
		})
	}

	app.Run(":9090")
	log.Println("tet")
}
