package user

import (
	"kaco/service"
	"kaco/admin"
	"github.com/gin-gonic/gin"
)


// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}
}



