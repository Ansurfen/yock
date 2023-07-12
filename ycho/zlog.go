// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
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

func (opt YchoOpt) String() string {
	return ""
}

type zlog struct {
	log *zap.Logger
	Vlog
}

var _ yocki.Ycho = (*zlog)(nil)

func (z *zlog) Write(p []byte) (int, error) {
	return len(p), nil
}

func (z *zlog) Info(msg string) {
	z.log.Info(msg)
}

func (z *zlog) Infof(msg string, v ...any) {
	z.log.Info(fmt.Sprintf(msg, v...))
}

func (z *zlog) Fatal(msg string) {
	z.log.Fatal(msg)
}

func (z *zlog) Fatalf(msg string, v ...any) {
	z.log.Fatal(fmt.Sprintf(msg, v...))
}

func (z *zlog) Debug(msg string) {
	z.log.Debug(msg)
}

func (z *zlog) Debugf(msg string, v ...any) {
	z.log.Debug(fmt.Sprintf(msg, v...))
}

func (z *zlog) Warn(msg string) {
	z.log.Warn(msg)
}

func (z *zlog) Warnf(msg string, v ...any) {
	z.log.Warn(fmt.Sprintf(msg, v...))
}

func (z *zlog) Error(msg string) {
	z.log.Error(msg)
}

func (z *zlog) Errorf(msg string, v ...any) {
	z.log.Error(fmt.Sprintf(msg, v...))
}

func NewZLog(conf YchoOpt) (*zlog, error) {
	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	writeSyncer, err := getLogWriter(conf)
	if err != nil {
		return nil, err
	}
	encoder := getEncoder(conf)
	level, ok := logLevel[conf.Level]
	if !ok {
		level = logLevel["info"]
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	return &zlog{log: logger}, nil
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
			enc.AppendString(t.Format(defaultTimeFormat))
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
		conf.Path = util.Pathf("@/log")
	}
	if err := util.SafeMkdirs(conf.Path); err != nil {
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
