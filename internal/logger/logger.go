package logger

import (
	"github.com/charmbracelet/log"
	"github.com/nqui/go-project-template/internal/config"
	"os"
	"time"
)

var cfg = config.GetLogging()

func NewLogger() *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Level:           cfg.Level,
	})
	return logger
}
