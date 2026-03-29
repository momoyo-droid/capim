package main

import (
	"github.com/gin-gonic/gin"
	"github.com/momoyo-droid/capim/api/internal/config"
	"github.com/momoyo-droid/capim/api/internal/handler"
	"github.com/momoyo-droid/capim/api/internal/repository"
	"github.com/momoyo-droid/capim/api/internal/repository/postgres"
	"github.com/momoyo-droid/capim/api/internal/service"
	"go.uber.org/zap"
)

// logger initializes a zap.Logger instance for logging throughout the application.
// It logs the startup message and returns the logger instance for use in the main function
// and other parts of the application.
func logger() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())

	logger.Info("Starting clinical management API")

	return logger
}

func main() {
	cfg, err := config.LoadConfig()

	zapLogger := logger()

	defer func() {
		_ = zapLogger.Sync()
	}()

	if err != nil {
		zapLogger.Fatal("Failed to load configuration", zap.Error(err))
	}

	db, err := postgres.NewDatabaseConnection(cfg)
	if err != nil {
		zapLogger.Fatal("Failed to connect to database", zap.Error(err))
	}

	repository := repository.NewSellerRepository(db)
	sellerService := service.NewSellerService(repository, zapLogger)
	sellerHandler := handler.SellerHandler{
		Service: sellerService,
		Logger:  zapLogger,
	}

	router := gin.Default()

	router.GET("/health", handler.HealthCheck)
	router.POST("/sellers", sellerHandler.CreateSeller)
	router.GET("/sellers", sellerHandler.GetAllSellers)
	router.GET("/sellers/:id", sellerHandler.GetSellerByID)
	router.DELETE("/sellers/:id", sellerHandler.DeleteSellerByID)
	router.PATCH("/sellers/:id", sellerHandler.UpdateSellerByID)
	router.PATCH("/owners/:id", sellerHandler.UpdateOwnerByID)

	zapLogger.Info("Server is running on port " + cfg.Port)

	if err := router.Run("localhost:" + cfg.Port); err != nil {
		zapLogger.Fatal("Failed to start server", zap.Error(err))
	}

}
