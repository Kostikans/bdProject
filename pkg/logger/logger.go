package logger

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/kostikans/bdProject/configs"
	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	*logrus.Logger
}

func NewLogger(writer io.Writer) *CustomLogger {
	Logger := &CustomLogger{logrus.New()}
	Formatter := new(logrus.JSONFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Logger.SetFormatter(Formatter)
	Logger.SetOutput(writer)
	return Logger
}

func (l *CustomLogger) relative(path string) string {
	return strings.TrimPrefix(filepath.ToSlash(path), configs.PrefixPath)
}

func (l *CustomLogger) LogError(err error) {
	l.WithFields(logrus.Fields{}).Error(err)
}

func (l *CustomLogger) StartReq(r http.Request, rid string) {
	l.WithFields(logrus.Fields{
		"id":         rid,
		"usr_addr":   r.RemoteAddr,
		"req_URI":    r.RequestURI,
		"method":     r.Method,
		"user_agent": r.UserAgent(),
	}).Info("request started")
}

func (l *CustomLogger) EndReq(start time.Time) {
	l.WithFields(logrus.Fields{

		"elapsed_time,μs": time.Since(start).Microseconds(),
	}).Info("request ended")
}

func (l *CustomLogger) LogWarning(msg string) {
	l.WithFields(logrus.Fields{}).Warn(msg)
}
