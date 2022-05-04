package api

import (
	"gin-web/user/services"
	"github.com/gin-gonic/gin"
)

//login
func Login(context *gin.Context) {
	var login services.Login
	context.ShouldBind(&login)
	context.JSON(200, login.LoginLogic())

}
