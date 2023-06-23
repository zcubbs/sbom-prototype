package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zcubbs/zlogger/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LoggerType string

const (
	LogrusLogger LoggerType = "logrus"
	ZapLogger    LoggerType = "zap"
)

func SetupLogger(loggerType LoggerType) {
	if loggerType == LogrusLogger {
		setupLogrus()
		return
	}
	if loggerType == ZapLogger {
		setupZap()
		return
	}

	panic("invalid logger type")
}

func setupLogrus() {
	logrusLog := logrus.New()
	logrusLog.SetReportCaller(false)
	logrusLog.SetFormatter(&logrus.JSONFormatter{})
	logrusLog.SetOutput(os.Stdout)
	logrusLog.SetLevel(logrus.InfoLevel)
	log, _ := logger.NewLogrusLogger(logrusLog)
	logger.ReplaceGlobals(log)
}

func setupZap() {
	consoleEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewCore(consoleEncoder,
		zapcore.Lock(zapcore.AddSync(os.Stderr)),
		zapcore.DebugLevel)
	zapLogger := zap.New(core)
	log, _ := logger.NewZapLogger(zapLogger)
	logger.ReplaceGlobals(log)
}
