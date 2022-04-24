package log

import (
	"go.uber.org/zap"
)

// Logger Log component
var Logger *zap.SugaredLogger

func init() {
	l, _ := zap.NewDevelopment()
	Logger = l.Sugar()
}
