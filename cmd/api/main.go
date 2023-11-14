package main

import (
	_ "github.com/Dev-Mw/thanos-s3-adapter/cmd/api/docs"
	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/controllers"
	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/handlers"
	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/logs"
	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/models"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
)

var log = logs.GetLog()

// @title Thanos to S3 Adapter APIs
// @version 2.0
// @description Extracts metrics collected by Thanos and stores them in an S3 Bucket in JSON GZ format.
// @contact.name Cloud Data
// @contact.email mld-governance-cloud-economics-cloud-data@dars.dev
// @host localhost:9001
// @BasePath /api/v1
// @schemes http
func main() {
	// Get config
	runConfig := models.GetConfig()
	log.Debug(runConfig)

	// Get AWS config
	awsConfig := models.GetAWSConfig()
	log.Debug(awsConfig)

	// Set channels
	metricChannel := make(chan models.MetricChannel)
	seriesChannel := make(chan models.SeriesChannel)

	// Set Series Formatter
	go controllers.SeriesFormat(metricChannel, seriesChannel)

	// Set Series Store
	go controllers.SeriesStore(awsConfig, seriesChannel)

	// Run
	go controllers.Run(runConfig, metricChannel)

	// Gin Setup
	router := gin.New()
	router.Use(ginlogrus.Logger(log), gin.Recovery())

	// Routes
	router.POST("/api/v1/on_demand", handlers.BindConfig(metricChannel))
	router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run server
	router.Run(":9001")
}
