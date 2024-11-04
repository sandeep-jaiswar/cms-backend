package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"github.com/sandeep-jaiswar/cms-backend/responses"
)

func GetProducts(c *gin.Context) {
	products, err := repositories.ProductRepo.FindAll()
	if err != nil {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse("Failed to retrieve products"))
		return
	}
	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(products))
}

func CreateProducts(c *gin.Context) {
	var products []models.Product
	if err := c.ShouldBindJSON(&products); err != nil {
		responses.WriteResponse(c, http.StatusBadRequest, responses.NewErrorResponse("Invalid input"))
		return
	}

	var errors []string
	for _, product := range products {
		if err := repositories.ProductRepo.Create(&product); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse("Failed to create some products"))
		return
	}
	responses.WriteResponse(c, http.StatusCreated, responses.NewSuccessResponse(products))
}

func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.WriteResponse(c, http.StatusBadRequest, responses.NewErrorResponse("Invalid input"))
		return
	}

	product, err := repositories.ProductRepo.FindByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			responses.WriteResponse(c, http.StatusNotFound, responses.NewErrorResponse("Product not found"))
			return
		}
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse("Failed to retrieve product"))
		return
	}

	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(product))
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.WriteResponse(c, http.StatusBadRequest, responses.NewErrorResponse("Invalid input"))
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		responses.WriteResponse(c, http.StatusBadRequest, responses.NewErrorResponse("Invalid input"))
		return
	}

	product.ID = uint(id)

	if err := repositories.ProductRepo.Update(&product); err != nil {
		if err.Error() == "record not found" {
			responses.WriteResponse(c, http.StatusNotFound, responses.NewErrorResponse("Product not found"))
			return
		}
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse("Failed to update product"))
		return
	}

	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(product))
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.WriteResponse(c, http.StatusBadRequest, responses.NewErrorResponse("Invalid input"))
		return
	}

	if err := repositories.ProductRepo.Delete(uint(id)); err != nil {
		// Handle not found case
		if err.Error() == "record not found" {
			responses.WriteResponse(c, http.StatusNotFound, responses.NewErrorResponse("Product not found"))
			return
		}
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse("Failed to delete product"))
		return
	}

	c.Status(http.StatusNoContent) // No content returned on successful deletion
}
