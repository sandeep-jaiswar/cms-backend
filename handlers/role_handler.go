package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"github.com/sandeep-jaiswar/cms-backend/responses"
)

func GetRoles(c *gin.Context) {
	roles, err := repositories.RoleRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve roles"})
		return
	}
	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(roles))

}

func CreateRoles(c *gin.Context) {
	var roles []models.Role
	if err := c.ShouldBindJSON(&roles); err != nil {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))

		return
	}
	var errors []string
	for _, role := range roles {
		if err := repositories.RoleRepo.Create(&role); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse(errors[0]))

		return
	}
	responses.WriteResponse(c, http.StatusCreated, responses.NewSuccessResponse(roles))

}
