package product

import (
	"kaco/model"
	"kaco/serializer"

	"github.com/gin-gonic/gin"
)

type ProductListService struct {
}

func (service *ProductListService) ProductList(c *gin.Context) serializer.Response {

	var res []model.Product
	model.DB.Find(&res)

	return serializer.Response{Code: 200, Msg: "", Data: res}
}
