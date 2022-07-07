package main

import (
	"os"

	"github.com/albertopformoso/resume-be/internal/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Initialize Echo instance
var e = echo.New()

func main() {
	// Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// TODO: Implement Driver selection

	// TODO: Initialize repository

	// TODO: Initialize service

	// TODO: Initialize controller

	// Echo configuration
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	apiVersionURI := e.Group("/api/v1")
	router.ProfileRoutes(apiVersionURI)
	router.ExperienceRoutes(apiVersionURI)
	router.ProjectsRoutes(apiVersionURI)
	router.EducationRoutes(apiVersionURI)
	router.CertificationsRoutes(apiVersionURI)

	e.Logger.Fatal(e.Start(":" + port))
}
