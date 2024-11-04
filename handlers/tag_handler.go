package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"github.com/sandeep-jaiswar/cms-backend/responses"
)

func GetTags(c *gin.Context) {
	tags, err := repositories.TagRepo.FindAll()
	if err != nil {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))
		return
	}
	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(tags))

}

func CreateTags(c *gin.Context) {
	var tags []models.Tag
	if err := c.ShouldBindJSON(&tags); err != nil {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))

		return
	}
	var errors []string
	for _, tag := range tags {
		if err := repositories.TagRepo.Create(&tag); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(errors[0]))

		return
	}
	responses.WriteResponse(c, http.StatusCreated, responses.NewSuccessResponse(tags))

}
