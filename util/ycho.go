package util

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ansurfen/cushion/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOpt struct {
	Level          string
	Format         string
	Path           string
	FileName       string
	FileMaxSize    int
	FileMaxBackups int
	MaxAge         int
	Compress       bool
	Stdout         bool
}

var Ycho *zap.Logger

func initLogger(conf LoggerOpt) error {
	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	writeSyncer, err := getLogWriter(conf)
	if err != nil {
		return err
	}
	encoder := getEncoder(conf)
	level, ok := logLevel[conf.Level]
	if !ok {
		level = logLevel["info"]
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return nil
}

func getEncoder(conf LoggerOpt) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000 -0700"))
		},
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	if conf.Format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(conf LoggerOpt) (zapcore.WriteSyncer, error) {
	DefaultLogPath := ""
	if len(DefaultLogPath) == 0 {
		DefaultLogPath = Pathf("@/log")
	}
	if exist := utils.IsExist(conf.Path); !exist {
		if conf.Path == "" {
			conf.Path = DefaultLogPath
		}
		if err := os.MkdirAll(conf.Path, os.ModePerm); err != nil {
			conf.Path = DefaultLogPath
			if err := os.MkdirAll(conf.Path, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.Path, conf.FileName),
		MaxSize:    conf.FileMaxSize,
		MaxBackups: conf.FileMaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}
	if conf.Stdout {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
	} else {
		return zapcore.AddSync(lumberJackLogger), nil
	}
}
