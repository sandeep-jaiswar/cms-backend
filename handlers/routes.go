package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/tags", GetTags)
		api.POST("/tags", CreateTags)

		api.GET("/media", GetMedias)
		api.POST("/media", CreateMedias)

		api.GET("/categories", GetCategories)
		api.POST("/categories", CreateCategories)

		api.GET("/products", GetProducts)
		api.POST("/products", CreateProducts)
		api.PUT("/product/:id", UpdateProduct)
		api.DELETE("/product/:id", DeleteProduct)
		api.GET("/product/:id", GetProductByID)

		api.GET("/orders", GetOrders)
		api.POST("/orders", CreateOrders)
		api.GET("/users", GetUsers)
		api.POST("/users", CreateUsers)
		api.GET("/roles", GetRoles)
		api.POST("/roles", CreateRoles)
	}
}
