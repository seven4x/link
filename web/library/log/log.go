package log

import (
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	//todo 定制
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()

}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Debug(args ...interface{}) {
	sugar.Debug(args)
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}
