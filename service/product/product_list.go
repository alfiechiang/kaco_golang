package product

import (
	"kaco/model"
	"kaco/serializer"
	"kaco/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductListService struct {
	util.Pager

	Category string `form:"category"`
	Name     string `form:"name"`
}

func (service *ProductListService) ProductList(c *gin.Context) serializer.Response {

	var res []model.Product
	var total int64
	model.DB.Scopes(util.Paginate(service.Page, service.PageSize)).
		Scopes(service.productListSearch()).Find(&res).
		Limit(-1).Offset(-1).Count(&total)

	pageRes := serializer.PageListFormat(service.Page, service.PageSize, total, res)

	return serializer.Response{Code: 200, Msg: "", Data: pageRes}
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
