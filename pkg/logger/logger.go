package logger

import (
	"github.com/AbdulwahabNour/movies/config"
	log "github.com/sirupsen/logrus"
)

type Logger interface {
	ErrorLog(args ...any)
	ErrorLogWithFields(f log.Fields, args ...any)
	InfoLog(args ...any)
	InfoLogWithFields(f log.Fields, args ...any)
	WarnLog(args ...any)
	WarnLogWithFields(f log.Fields, args ...any)
	DebugLog(args ...any)
	DebugLogWithFields(f log.Fields, args ...any)
}

type apiLogger struct {
	logger *log.Logger
}

func NewApiLogger(conf *config.Config) Logger {
	logger := log.New()
	setLogLevel(conf.Logger.Level, logger)
	setLogFormate(conf.Logger.Formate, logger)
	return &apiLogger{
		logger: logger,
	}
}

func setLogLevel(level string, logger *log.Logger) {

	switch level {
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "trace":
		logger.SetLevel(log.TraceLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	case "debug":
		logger.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
func setLogFormate(formate string, logger *log.Logger) {

	switch formate {
	case "json":
		logger.SetFormatter(&log.JSONFormatter{})
	default:
		logger.SetFormatter(&log.TextFormatter{})
	}
}
func (l *apiLogger) ErrorLog(args ...any) {
	l.logger.Error(args...)
}
func (l *apiLogger) ErrorLogWithFields(f log.Fields, args ...any) {
	l.logger.WithFields(f).Error(args...)
}

func (l *apiLogger) InfoLog(args ...any) {
	l.logger.Info(args...)
}

func (l *apiLogger) InfoLogWithFields(f log.Fields, args ...any) {
	l.logger.WithFields(f).Info(args...)
}

func (l *apiLogger) WarnLog(args ...any) {
	l.logger.Warn(args...)
}
func (l *apiLogger) WarnLogWithFields(f log.Fields, args ...any) {
	l.logger.WithFields(f).Warn(args...)
}

func (l *apiLogger) DebugLog(args ...any) {
	l.logger.Debug(args...)
}
func (l *apiLogger) DebugLogWithFields(f log.Fields, args ...any) {
	l.logger.WithFields(f).Debug(args...)
}
