package user

import (
	"kaco/serializer"

	"github.com/gin-gonic/gin"
)

func UserLogout(c *gin.Context) {

	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
