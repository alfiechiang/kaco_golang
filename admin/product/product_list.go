package product

import (
	"kaco/admin"
	"kaco/service/product"

	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录接口
func ProductList(c *gin.Context) {
	var service product.ProductListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ProductList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}
}
