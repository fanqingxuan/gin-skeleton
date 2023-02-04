package logx

import (
	"context"
)

var l *fileLogger

func NewLog(loglevel string) {
	l = &fileLogger{
		zapLogger: initLogger(loglevel),
	}
}

func WithContext(ctx context.Context) Logger {
	return &fileLogger{
		zapLogger: l.zapLogger,
		ctx:       ctx,
	}
}

func Debug(message interface{}) {
	l.Debug(message)
}

func Info(message interface{}) {
	l.Info(message)
}

func Warn(message interface{}) {
	l.Warn(message)
}

func Error(message interface{}) {
	l.Error(message)
}

func Debugf(format string, message ...interface{}) {
	l.Debugf(format, message...)
}

func Infof(format string, message ...interface{}) {
	l.Infof(format, message...)
}

func Warnf(format string, message ...interface{}) {
	l.Warnf(format, message...)
}

func Errorf(format string, message ...interface{}) {

	l.Errorf(format, message...)
}
