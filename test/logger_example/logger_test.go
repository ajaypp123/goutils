package collection_example

import (
	"testing"

	"github.com/ajaypp123/goutils/logger"
)

func TestLogger(t *testing.T) {
	logData := logger.LogData{}
	logData.InitLogger()
	logger.Logger.Info("Test")
}
