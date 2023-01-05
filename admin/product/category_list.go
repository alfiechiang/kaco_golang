package product

import (
	"kaco/admin"
	"kaco/service/product"

	"github.com/gin-gonic/gin"
)

func CategoryList(c *gin.Context) {
	var service product.CategoryListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CategoryList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}
}
