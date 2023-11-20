package models

import (
	"bytes"
	"time"
)

type MetricChannel struct {
	Body          []byte
	Config        QueryConfig
	StartTimeUnix string
	EndTimeUnix   string
	Cluster       string
}

type SeriesChannel struct {
	ByteData    bytes.Buffer
	Config      QueryConfig
	ExtractTime time.Time
	Cluster     string
}

type JSONSchema struct {
	Start      interface{}       `json:"start"`
	End        interface{}       `json:"end"`
	Metric     map[string]string `json:"metric"`
	Interval   string            `json:"interval"`
	Step       string            `json:"step"`
	Values     [][]interface{}   `json:"values"`
	Expression string            `json:"expression"`
	Ds         string            `json:"ds"`
}

type ThanosMetricResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Values [][]interface{}   `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

type ThanosStoreResponse struct {
	Status string `json:"status"`
	Data   struct {
		Sidecar []struct {
			Name      string    `json:"name"`
			LastCheck time.Time `json:"lastCheck"`
			LastError any       `json:"lastError"`
			LabelSets []struct {
				ClusterID string `json:"cluster_id"`
			} `json:"labelSets"`
			MinTime int64 `json:"minTime"`
			MaxTime int64 `json:"maxTime"`
		} `json:"sidecar"`
		Store []struct {
			Name      string    `json:"name"`
			LastCheck time.Time `json:"lastCheck"`
			LastError any       `json:"lastError"`
			LabelSets []struct {
				ClusterID string `json:"cluster_id"`
			} `json:"labelSets"`
			MinTime int64 `json:"minTime"`
			MaxTime int64 `json:"maxTime"`
		} `json:"store"`
	} `json:"data"`
}
