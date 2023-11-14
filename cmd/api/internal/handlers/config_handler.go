package handlers

import (
	"net/http"

	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/controllers"
	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/logs"
	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/models"
	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/utils"

	"github.com/gin-gonic/gin"
)

var log = logs.GetLog()

// BindConfig godoc
// @Summary Bind Config
// @Description Bind the received JSON to the Config struct
// @Tags Config
// @Accept json
// @Produce json
// @Param config body models.QueryConfig true "QueryConfig"
// @Success 202 {object} models.QueryConfig
// @Failure 400 {object} map[string]any
// @Router /on_demand [post]
func BindConfig(metricChannel chan models.MetricChannel) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		newConfig := models.GetConfig()

		// Call BindJSON to bind the received JSON to newConfig
		if err := c.BindJSON(&newConfig); err != nil {
			log.Error("Bind JSON error: " + err.Error())
		}

		// Response
		_, err := utils.DatesValidation(newConfig.StartDate, newConfig.EndDate, newConfig.Interval)
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusAccepted, gin.H{"config": newConfig})
			go controllers.MetricRequest(newConfig, metricChannel)
			return
		}
	}
	return fn
}
