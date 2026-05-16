package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTempRoutes(r *gin.RouterGroup) {
	tmp := r.Group("/tmp")
	tmp.GET("/health", healthcheck)
	tmp.GET("/redirect", redirect)
	tmp.GET("/status-found", statusFound)
	tmp.GET("/result", result)
	tmp.GET("/internal-redirect", internalRedirect)
}

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
}

func statusFound(c *gin.Context) {
	c.Redirect(http.StatusFound, "./result")
}

func result(c *gin.Context) {
	c.String(http.StatusOK, "Redirected here!")
}

func internalRedirect(c *gin.Context) {
	result(c)
}
