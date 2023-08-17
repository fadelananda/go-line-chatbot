package utils

import (
	"github.com/sirupsen/logrus"
)

// TODO:
// 1. create logger middleware to log request and response code
// 2. create logger to function
// Log is the global logger object
var logger *logrus.Logger

// InitLogger initializes the global logger
func InitLogger() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func LogInfo(msg string, obj map[string]interface{}) {
	loggerObj := logrus.Fields{}
	for key, value := range obj {
		loggerObj[key] = value
	}

	entries := logger.WithFields(loggerObj)
	entries.Info(msg)
}
