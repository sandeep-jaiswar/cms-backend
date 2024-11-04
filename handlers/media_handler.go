package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"github.com/sandeep-jaiswar/cms-backend/responses"
)

func GetMedias(c *gin.Context) {
	medias, err := repositories.MediaRepo.FindAll()
	if err != nil {
		log.Printf("Error retrieving media: %v", err)
		responses.WriteResponse(c, http.StatusInternalServerError, responses.NewErrorResponse("Failed to retrieve media"))
		return
	}
	responses.WriteResponse(c, http.StatusOK, responses.NewSuccessResponse(medias))
}

func CreateMedias(c *gin.Context) {
	form, err := c.MultipartForm()
	log.Print(form)
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		responses.WriteResponse(c, http.StatusBadRequest, responses.NewErrorResponse("Invalid file data"))
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		responses.WriteResponse(c, http.StatusBadRequest, responses.NewErrorResponse("At least one file is required"))
		return
	}

	var mediaList []models.Media
	var errors []string

	for _, file := range files {
		uniqueFileName := uuid.New().String() + filepath.Ext(file.Filename)
		savePath := filepath.Join("uploads", uniqueFileName)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			log.Printf("Error saving file %s: %v", file.Filename, err)
			errors = append(errors, "Failed to save file: "+file.Filename)
			continue
		}

		media := models.Media{
			URL:        savePath,
			FileType:   file.Header.Get("Content-Type"),
			FileSize:   int(file.Size),
			UploadedBy: 0,
			CreatedAt:  time.Now(),
		}

		if err := repositories.MediaRepo.Create(&media); err != nil {
			log.Printf("Error saving media details for file %s: %v", file.Filename, err)
			errors = append(errors, "Failed to save media details for file: "+file.Filename)
			continue
		}

		mediaList = append(mediaList, media)
	}

	if len(errors) > 0 {
		responses.WriteResponse(c, http.StatusPartialContent, responses.NewErrorResponse("Some files could not be processed: " + formatErrors(errors)))
		return
	}

	responses.WriteResponse(c, http.StatusCreated, responses.NewSuccessResponse(mediaList))
}

func formatErrors(errors []string) string {
	return strings.Join(errors, "; ")
}
