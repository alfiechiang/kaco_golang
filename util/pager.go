package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pager struct {
	Page     int `form:"page"  json:"page"`
	PageSize int `form:"page_size"  json:"page_size"`
}

func GetPager(c *gin.Context) Pager {
	p1 := c.Param("page_size")
	p2 := c.Param("page")
	page_size, _ := strconv.Atoi(p1)
	page, _ := strconv.Atoi(p2)

	return Pager{
		PageSize: page_size,
		Page:     page,
	}
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
