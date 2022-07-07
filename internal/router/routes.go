package router

import (
	"github.com/labstack/echo"
)

// ProfileRoutes routes for profile info
func ProfileRoutes(route *echo.Group) {
	profileGroup := route.Group("/profile")
	profileGroup.GET("/", nil)
	profileGroup.PUT("/", nil)
}

// ExperienceRoutes routes for experience data
func ExperienceRoutes(route *echo.Group) {
	experienceGroup := route.Group("/experience")
	experienceGroup.GET("/", nil)
	experienceGroup.POST("/", nil)
	experienceGroup.PUT("/:id", nil)
	experienceGroup.DELETE("/:id", nil)
}

// ProjectsRoutes routes for projects data
func ProjectsRoutes(route *echo.Group) {
	projectsGroup := route.Group("/projects")
	projectsGroup.GET("/", nil)
}

// EducationRoutes routes for education data
func EducationRoutes(route *echo.Group) {
	educationGroup := route.Group("/education")
	educationGroup.GET("/", nil)
	educationGroup.POST("/", nil)
	educationGroup.PUT("/:id", nil)
	educationGroup.DELETE("/:id", nil)
}

// CertificationsRoutes routes for certifications data
func CertificationsRoutes(route *echo.Group) {
	certificationsGroup := route.Group("/certifications")
	certificationsGroup.GET("/", nil)
	certificationsGroup.POST("/", nil)
	certificationsGroup.PUT("/:id", nil)
	certificationsGroup.DELETE("/:id", nil)
}
