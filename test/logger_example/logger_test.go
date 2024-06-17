package collection_example

import (
	"context"
	"testing"

	"github.com/ajaypp123/goutils/logger"
)

func TestLogger(t *testing.T) {
	ctx := context.Background()

	logger.Logger.Info(ctx, "Test")
}
