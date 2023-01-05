package product

import (
	"kaco/admin"
	"kaco/service/product"

	"github.com/gin-gonic/gin"
)

func ProducCreate(c *gin.Context) {
	var service product.ProductCreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ProductCreate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}
}
