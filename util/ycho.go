// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Ycho is yock's logging component, which is packaged from zap.

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

// YchoOpt indicates the configuration of ycho.
// It is stored in the yockConf, loaded via init.
// You can set it using yock conf command.
type YchoOpt struct {
	Level          string `yaml:"level"`
	Format         string `yaml:"format"`
	Path           string `yaml:"path"`
	FileName       string `yaml:"filename"`
	FileMaxSize    int    `yaml:"fileMaxSize"`
	FileMaxBackups int    `yaml:"fileMaxBackups"`
	MaxAge         int    `yaml:"maxAge"`
	Compress       bool   `yaml:"compress"`
	Stdout         bool   `yaml:"stdout"`
}

var Ycho *zap.Logger

func initYcho(conf YchoOpt) error {
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

func getEncoder(conf YchoOpt) zapcore.Encoder {
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

func getLogWriter(conf YchoOpt) (zapcore.WriteSyncer, error) {
	if len(conf.Path) == 0 {
		conf.Path = Pathf("@/log")
	}
	if err := utils.SafeMkdirs(conf.Path); err != nil {
		return nil, err
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
