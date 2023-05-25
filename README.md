#  <img src="docs/static/yock.ico" width = "60" height = "60" alt="logo" align=center />Yock

[![Go Report Card](https://goreportcard.com/badge/github.com/ansurfen/cushion)](https://goreportcard.com/report/github.com/ansurfen/yock)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/ansurfen/yock?status.svg)](https://pkg.go.dev/github.com/ansurfen/yock)

English | [简体中文](./docs/zh_cn/README.md)

Yock is a solution of cross platform to compose distributed build stream. It's able to act as software package tool, like Homebrew, rpm, winget and so on. It also is used for dependency manager (pip, npm, maven, etc.) of programming languages. On top of this, yock also implements distributed build tasks based on grpc and goroutines (and can even build cluster for this). You can think of it as the lua version of the nodejs framework, except that it focuses on composition and is more lightweight.

## Architecture
![arch](docs/static/arch.png)

* Yctl: it's used for scheduling yock's commands.
* YockPack: it's mainly used for preprocessing lua file, such as schema decomposition, decomposing a lua source file into multiple lua files according to a given modes for distributed system.
* YockScheduler: the scheduler is responsible for running the lua code, and launchs goroutines to execute in tasks.
* YPM: yock package manager, used for completion and loading dependencies.

## Installation

You can download the binary version here, or try the following two methods.
`NOTE`: After downloading, you must mount yock to the local environment, and need to manually run the depoly script in the compressed package to use it.

#### Build by yock

Yock implements something like "bootstrap", meaning it's able to build itself. Of course, the prerequisite for all this also requires go compiler.

The first, to fetch source code of yock
```cmd
git clone https://github.com/Ansurfen/yock.git
```

Execute the go command and schedule the yock build script to build yock
```cmd
cd ctl
go run . run ../auto/build.lua all
```

#### Embed in Go

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

## Document

You can find information about module development and yock development here.

## Center Repository

If you want to register a module to yock so that it can use an identifier index instead of URL, see here.

## TODO

- [ ] implement mock service and compose middlewares (YockCloud)
- [ ] implement preprocess and DSL to enhance lua syntax

## License

This software is licensed under the MIT license, see [LICENSE](./LICENSE) for more information.