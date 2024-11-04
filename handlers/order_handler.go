package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"github.com/sandeep-jaiswar/cms-backend/responses"
)

func GetOrders(c *gin.Context) {
	orders, err := repositories.OrderRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}
	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(orders))
}

func CreateOrders(c *gin.Context) {
	var orders []models.Order
	if err := c.ShouldBindJSON(&orders); err != nil {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))

		return
	}
	var errors []string
	for _, order := range orders {
		if err := repositories.OrderRepo.Create(&order); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(errors[0]))

		return
	}
	responses.WriteResponse(c, http.StatusCreated, responses.NewSuccessResponse(orders))

}
