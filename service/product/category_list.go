package product

import (
	"kaco/model"
	"kaco/serializer"

	"github.com/gin-gonic/gin"
)

type CategoryListService struct {
}

func (service *CategoryListService) CategoryList(c *gin.Context) serializer.Response {

	var res []string
	model.DB.Model(&model.Product{}).Distinct().Pluck("category", &res)
	
	return serializer.Response{Code: 200, Msg: "", Data: res}
}
