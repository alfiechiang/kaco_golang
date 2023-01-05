package server

import (
	"kaco/admin/product"
	"kaco/admin/user"
	"kaco/service/product"

	"kaco/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	Admin := r.Group("/admin")
	{
		// 用户登录
		Admin.POST("user/login", user.UserLogin)
		Admin.GET("product", product.ProductList)

		// 需要登录保护的
		auth := Admin.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
		}
	}
	return r
}
