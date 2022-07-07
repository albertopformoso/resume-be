package main

import (
	"log"
	"os"
	"strings"

	"github.com/albertopformoso/resume-be/internal/db"
	"github.com/albertopformoso/resume-be/internal/model"
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

	// Driver selection implementation
	dbDriver := os.Getenv("DRIVER")
	log.Printf("Selected driver: %q\n", dbDriver)

	var driver model.DBDriver
	switch strings.ToLower(dbDriver) {
	case "postgres", "pg":
		driver = db.DriverPG
	default:
		driver = db.DriverUndefined
	}

	database := db.NewDB(driver)
	log.Println(database) // temporary line to dismiss errors

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
