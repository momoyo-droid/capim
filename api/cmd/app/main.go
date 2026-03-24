package main

import (
	"github.com/gin-gonic/gin"
	"github.com/momoyo-droid/capim/api/internal/config"
	"go.uber.org/zap"
)

func logger() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()

	logger.Info("Starting clinical management API")
	return logger
}

func main() {
	cfg, err := config.LoadConfig()

	logger := logger()

	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	router := gin.Default()

	logger.Info("Server is running on port " + cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

}
