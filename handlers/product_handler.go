package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
)

func GetProducts(c *gin.Context) {
	products, err := repositories.ProductRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProducts(c *gin.Context) {
	var products []models.Product
    if err := c.ShouldBindJSON(&products); err != nil {
        c.Error(err)
        return
    }
    var errors []string
    for _, product := range products {
        if err := repositories.ProductRepo.Create(&product); err != nil {
            errors = append(errors, err.Error())
        }
    }
    if len(errors) > 0 {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":  "Failed to create some products",
            "details": errors,
        })
        return
    }
	c.JSON(http.StatusCreated, products)
}
