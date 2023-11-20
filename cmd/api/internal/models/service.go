package models

import (
	"os"
	"strconv"

	"github.com/Dev-Mw/thanos-s3-adapter/cmd/api/internal/logs"
)

type QueryConfig struct {
	EndpointAddress string `json:"EndpointAddress" swaggerignore:"true"`
	Query           string `json:"Query" swaggerignore:"true"`
	QueryStep       string `json:"QueryStep" example:"60s"`
	MetricBucket    string `json:"MetricBucket" swaggerignore:"true"`
	MetricName      string `json:"MetricName" swaggerignore:"true"`
	Interval        int    `json:"Interval" example:"3600" minimum:"3600"`
	Schedule        string `json:"Schedule" swaggerignore:"true"`
	StartDate       string `json:"StartDate" example:"2023-01-01T00:00:00Z" format:"2023-01-01T00:00:00Z"`
	EndDate         string `json:"EndDate" example:"2023-01-01T06:00:00Z" format:"2023-01-01T00:00:00Z"`
}

var log = logs.GetLog()

func GetConfig() QueryConfig {
	// Check Interval is a valid integer
	intInterval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		log.Errorf("Configuration error: %s", err.Error())
		return QueryConfig{}
	}

	// Set the configuration
	conf := QueryConfig{
		EndpointAddress: os.Getenv("ENDPOINT_ADDRESS"),
		Query:           os.Getenv("QUERY_STRING"),
		QueryStep:       os.Getenv("QUERY_STEP"),
		MetricBucket:    os.Getenv("METRIC_BUCKET"),
		MetricName:      os.Getenv("METRIC_NAME"),
		Interval:        intInterval,
		Schedule:        os.Getenv("SCHEDULE"),
	}
	return conf
}
