package logging

import (
	"github.com/go-kit/log"
	"os"
	"sync"
)

var (
	logger log.Logger
	once   sync.Once
)

func GetLogger() log.Logger {
	once.Do(func() {
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	})
	return logger
}
