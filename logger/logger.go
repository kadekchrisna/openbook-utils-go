package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type loggerInterface interface {
	Printf(string, ...interface{})
	Print(v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{GetLogOutput()},
		Level:       zap.NewAtomicLevelAt(GetLogLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}

}

func GetLogOutput() string {
	output := strings.TrimSpace(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}
	return output
}

func GetLogger() loggerInterface {
	return log
}

func (l logger) Printf(message string, v ...interface{}) {
	if len(v) == 0 {
		Info(message)
	} else {
		Info(fmt.Sprintf(message, v...))
	}

}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

// Info Overiding Info for flushing the buffer
func Info(message string, tags ...zap.Field) {
	log.log.Info(message)
	log.log.Sync()
}

// Error Overiding Error for flushing the buffer
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(message)
	log.log.Sync()
}
