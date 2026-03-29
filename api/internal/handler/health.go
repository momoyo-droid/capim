package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is a simple handler function that responds with a JSON object
// indicating the health status of the API.
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}
