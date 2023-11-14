package logs

import (
	"time"

	"github.com/sirupsen/logrus"
)

func GetLog() *logrus.Logger {
	// Logger Setup
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339, FullTimestamp: true})
	return log
}
