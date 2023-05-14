# Yock

[![Go Report Card](https://goreportcard.com/badge/github.com/ansurfen/cushion)](https://goreportcard.com/report/github.com/ansurfen/yock)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/ansurfen/yock?status.svg)](https://pkg.go.dev/github.com/ansurfen/yock)

English | [简体中文](./docs/zh_cn/README.md)

Yock is a solution of cross platform to compose distributed build stream.

## Install

### Embed in Go

To start, fetchs library by go mod.
```cmd
go get "github.com/ansurfen/yock"
```

Then, import library to use it on your project.
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