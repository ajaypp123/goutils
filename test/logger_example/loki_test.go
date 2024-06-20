package collection_example

import (
	"testing"

	"github.com/ajaypp123/goutils/logger"
)

func TestLokiLogger(t *testing.T) {
	conf, err := logger.GetSimpleLokiConfig("http://192.168.0.115:3100/loki/api/v1/push")
	if err != nil {
		panic("loki config failed")
	}
	lokiHook, err := logger.NewLokiHook(conf)
	if err != nil {
		panic("loki hook failed")
	}
	logData := logger.LogData{
		AppName: "golang_testapp",
		Hook:    lokiHook,
	}
	err = logData.InitLogger()
	if err != nil {
		panic("logger initialization failed")
	}
	logger.Logger.Info("Test")
}
