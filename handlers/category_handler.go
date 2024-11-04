package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"github.com/sandeep-jaiswar/cms-backend/responses"
)

func GetCategories(c *gin.Context) {
	categories, err := repositories.CategoryRepo.FindAll()
	if err != nil {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))
		return
	}
	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(categories))

}

func CreateCategories(c *gin.Context) {
	var categories []models.Category
	if err := c.ShouldBindJSON(&categories); err != nil {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))

		return
	}
	var errors []string
	for _, category := range categories {
		if err := repositories.CategoryRepo.Create(&category); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(errors[0]))

		return
	}
	responses.WriteResponse(c, http.StatusCreated, responses.NewSuccessResponse(categories))

}
