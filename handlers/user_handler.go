package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"github.com/sandeep-jaiswar/cms-backend/internal/utils"
	"github.com/sandeep-jaiswar/cms-backend/responses"
)

func GetUsers(c *gin.Context) {
	users, err := repositories.UserRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(users))

}

func CreateUsers(c *gin.Context) {
	var users []models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// Start transaction
	tx := repositories.UserRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
	}()

	var errors []string
	for i := range users {
		if err := assignRole(repositories.RoleRepo, &users[i]); err != nil {
			errors = append(errors, "User "+users[i].Username+": "+err.Error())
			continue
		}

		hashedPassword, err := utils.HashPassword(users[i].Password)
		if err != nil {
			errors = append(errors, "Failed to hash password for user: "+users[i].Username)
			continue
		}
		users[i].Password = hashedPassword

		if err := tx.Create(&users[i]); err != nil {
			errors = append(errors, "Failed to create user: "+users[i].Username)
			continue
		}
	}

	if len(errors) > 0 {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create some users", "details": errors})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction commit failed", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"users": users})
}

func assignRole(roleRepo repositories.Repository[models.Role], user *models.User) error {
	role, err := roleRepo.FindByID(user.RoleID)
	if err != nil {
		return fmt.Errorf("invalid RoleID %d for user %s", user.RoleID, user.Username)
	}

	user.Role = *role
	return nil
}
