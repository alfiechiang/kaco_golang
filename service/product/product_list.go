package product

import (
	"kaco/model"
	"kaco/serializer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductListService struct {
	Category string `form:"category"`
	Name     string `form:"name"`
}

func (service *ProductListService) ProductList(c *gin.Context) serializer.Response {

	var res []model.Product
	model.DB.Scopes(service.productListSearch()).Find(&res)

	return serializer.Response{Code: 200, Msg: "", Data: res}
}

func (service *ProductListService) productListSearch() func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		if service.Category != "" {
			db.Where("category", service.Category)
		}

		if service.Name != "" {
			db.Where("name", service.Name)
		}

		return db
	}

}
