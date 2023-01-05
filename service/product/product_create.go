package product

import (
	"kaco/model"
	"kaco/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductCreateService struct {
	Name        string `form:"name"`
	Price       int    `form:"price"`
	Category    string `form:"category"`
	HotSpot     string `form:"hot_spot" json:"hot_spot"`
	Description string `form:"description"`
}

func (service *ProductCreateService) ProductCreate(c *gin.Context) serializer.Response {

	product := model.Product{
		Name:        service.Name,
		Price:       service.Price,
		Category:    service.Category,
		ChCategory:  "蔬菜",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		HotSpot:     service.HotSpot,
		Description: service.Description,
	}

	model.DB.Create(&product)

	return serializer.Response{Code: 200, Msg: "", Data: make([]string, 0)}
}
