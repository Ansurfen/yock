#  <img src="../yock.ico" width = "60" height = "60" alt="logo" align=center />Yock

[![Go Report Card](https://goreportcard.com/badge/github.com/ansurfen/cushion)](https://goreportcard.com/report/github.com/ansurfen/yock)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/ansurfen/yock?status.svg)](https://pkg.go.dev/github.com/ansurfen/yock)

[English](../../README.md) | 简体中文

Yock 是一个跨平台的分布式构建流编排解决方案。

## 安装

### 嵌入Go语言

首先，通过go mod去获取库
```cmd
go get "github.com/ansurfen/yock"
```

接着，在你的项目中引入库使用它
```go
package main

import . "github.com/ansurfen/yock/cmd"

func main() {
	HTTP(HttpOpt{
		Method: "GET",
		Save:   true,
		Debug:  true,
		Dir:    ".",
		Filename: func(s string) string {
			return s
		},
	}, []string{"https://www.github.com"})    
}
```