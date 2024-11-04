package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/cms-backend/config"
	"github.com/sandeep-jaiswar/cms-backend/handlers"
)

func main() {
	config.LoadConfig()
	config.InitDB()
	router := gin.Default()
	handlers.SetupRoutes(router)
	port := config.AppConfig.Server.Port
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
