package main

import (
	"chijji-moni-backend-go/bootstrap"
	"chijji-moni-backend-go/routes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	r := app.Router

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(app)

	r.GET("/health", func(c *gin.Context) {
		// Return JSON respons
		c.String(http.StatusOK, "OK")
	})

	const port = 8080
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run(fmt.Sprintf(":%d", port))
}
