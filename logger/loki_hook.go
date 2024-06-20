package logger

import (
	"fmt"
	"net/url"
	"time"

	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki-client-go/pkg/urlutil"
	"github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

type LokiHook struct {
	client *loki.Client
}

func NewLokiHook(lokiConfig *loki.Config) (*LokiHook, error) {
	client, err := loki.New(*lokiConfig)
	if err != nil {
		return nil, err
	}

	return &LokiHook{client: client}, nil
}

func GetSimpleLokiConfig(lokiUrl string) (*loki.Config, error) {
	u, err := url.Parse(lokiUrl)
	if err != nil {
		return nil, err
	}
	return &loki.Config{
		URL: urlutil.URLValue{
			URL: u,
		},
	}, nil
}

func (hook *LokiHook) Fire(entry *log.Entry) error {
	labels := model.LabelSet{
		"level": model.LabelValue(entry.Level.String()),
		"app":   model.LabelValue("test-app"),
	}

	logLine := fmt.Sprintf("%s - %s", entry.Time.Format("2006-01-02T15:04:05.000Z"), entry.Message)

	err := hook.client.Handle(labels, time.Now(), logLine)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (hook *LokiHook) Levels() []log.Level {
	return log.AllLevels
}
