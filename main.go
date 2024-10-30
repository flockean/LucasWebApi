package main

import (
	"LucasApi/api/dbutils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/static"
)

func NewRouter() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	expectedHost := "localhost:8080"

	// Setup Security Headers
	router.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}

		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./public/", true)))

	// Setup route group for the API
	api := router.Group("/api")
	projets := api.Group("/projects")
	services := api.Group("/services")
	{
		apiHandler := func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Uniform API",
			})
		}

		projectHandler := func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"project_id":  "1",
				"name":        "Example",
				"description": "ExampleDescription",
			})
		}

		serviceHandler := func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"service_id":  "1",
				"name":        "Example",
				"description": "ExampleDescription",
			})
		}

		api.GET("", apiHandler)
		api.GET("/", apiHandler)
		projets.GET("", projectHandler)
		services.GET("", serviceHandler)

	}
	return router
}

func main() {
	_, err := dbutils.ConnectDatabase()
	if err != nil {
		return
	}

	httpPort := os.Getenv("API_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	// Initialize router
	router := NewRouter()
	// Create server with timeout
	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: router,
		// set timeout due CWE-400 - Potential Slowloris Attack
		ReadHeaderTimeout: 5 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
