package logger

import (
	"fmt"
	"os"

	"github.com/cnh1en/lets_go/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type AppLogger struct {
	config setting.LoggerSetting
}

func NewLogger(config setting.LoggerSetting) *zap.Logger {
	appLogger := &AppLogger{
		config: config,
	}
	return appLogger.getLogger()
}

func (l *AppLogger) getLogger() *zap.Logger {
	encoder := l.getEncoderLog()
	ws := l.getWriterSync()

	logLevel := switchLevel(l.config.Level)

	core := zapcore.NewCore(encoder, ws, logLevel)

	return zap.New(core, zap.AddCaller())
}

// format logger
func (l *AppLogger) getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.TimeKey = "time"

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func (l *AppLogger) getWriterSync() zapcore.WriteSyncer {
	file, err := os.OpenFile(l.config.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	hook := lumberjack.Logger{
		Filename:   l.config.Filename,
		MaxSize:    l.config.MaxSize,
		MaxBackups: l.config.MaxBackups,
		MaxAge:     l.config.MaxAge,
		Compress:   l.config.Compress,
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile, zapcore.AddSync(&hook))
}

func switchLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	}

	return zapcore.DebugLevel
}
