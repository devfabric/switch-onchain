package log

import (
	log "github.com/devfabric/HP-Log/logger"
	"go.uber.org/zap"
)

type Logger struct {
	ModName string
	Log     *zap.SugaredLogger
}

func GetLogger(ModName string, hpLog *log.HPLog) *Logger {
	return &Logger{
		ModName: ModName,
		Log:     hpLog.GetLogger(ModName),
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.Log.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Log.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Log.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Log.Error(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.Log.Debugf(template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.Log.Infof(template, args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.Log.Warnf(template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.Log.Errorf(template, args)
}
