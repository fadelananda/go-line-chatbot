package utils

import (
	"github.com/sirupsen/logrus"
)

// TODO:
// 1. create logger middleware to log request and response code
// 2. create logger to function
// Log is the global logger object
var Logger *logrus.Logger

// InitLogger initializes the global logger
func InitLogger() {
	Logger = logrus.New()
}
