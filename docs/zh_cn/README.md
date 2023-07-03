#  <img src="../static/yock.ico" width = "60" height = "60" alt="logo" align=center />Yock

[![Go Report Card](https://goreportcard.com/badge/github.com/ansurfen/cushion)](https://goreportcard.com/report/github.com/ansurfen/yock)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/ansurfen/yock?status.svg)](https://pkg.go.dev/github.com/ansurfen/yock)

[English](../../README.md) | 简体中文

Yock 是一个跨平台的分布式构建流编排解决方案。它能够作为软件包使用，就像Homebrew, rpm, winget等等。同时它还能充当编程语言的依赖管理的角色（pip，npm，maven等等）。在此基础上，yock还基于grpc和协程实现分布式构建任务（甚至可以为此搭建集群）。你可以将他视作nodejs框架的lua语言版本，不同的是他专注于编排，更加轻量。

## 架构
![arch](../static/arch.png)

* Yctl: 负责调度yock的命令。
* YockPack: 主要用于对lua文件的预处理，例如模式分解，将一份lua代码根据给定的模式分解成多份lua文件供分布式运行。
* YockScheduler: 调度器负责运行lua代码，以task为单位起协程执行。
* YPM: yock包管理，负责补全和装载依赖。

## 安装

你能够在这里下载二进制版本，或者尝试以下两种方式。
`注意`: 下载完后还需要将yock挂载到本地环境中，你需要手动运行压缩包内的depoly脚本去完成这个过程。

#### 使用yock构建

Yock实现了类似"自举"的操作，这意味着它能够自己构建自己。当然，这一切的前提还需要go语言的编译器。

首先，获取yock的源代码
```cmd
git clone https://github.com/Ansurfen/yock.git
```

执行go命令，调度yock构建脚本去构建yock
```cmd
cd ctl
./build.bat ffi 
```

#### 嵌入Go语言

安装yock库 (Yock支持版本 >= Go1.20)
```cmd
go get "github.com/ansurfen/yock"
```

yockr: 基于`gopher-lua`封装的lua语言解释器，上手起来更加简单。
```go
import (
	"fmt"

	yockr "github.com/ansurfen/yock/runtime"
)

type user struct {
	Name string
	Pwd  string
}

func main() {
	u := user{}
	r := yockr.New()
	if err := r.Eval(`return {name = "root", pwd = "123456"}`); err != nil {
		panic(err)
	}
	if err := r.State().CheckTable(1).Bind(&u); err != nil {
		panic(err)
	}
	fmt.Println(u)
}
```

yockc: 为yock提供了一系列简单的GNU命令，例如echo, rm, ls, curl等等。你能够直接调用它们，具体的Opt字段可以前往`/docs/yockc`查看详情。
```go
import yockc "github.com/ansurfen/yock/cmd"

func main() {
	yockc.Curl(yockc.HttpOpt{
		Method: "GET",
		Save:   true,
		Debug:  true,
		Dir:    ".",
		FilenameHandle: func(s string) string {
			return s
		},
	}, []string{"https://www.github.com"})
}
```

yocks: 基于yockr封装的调度器，它是yock的核心。yocks不仅包含了解释器以及标准库，还包含了简单的协程池，便于异步任务的调度。此外，yocks还配备了基于yocki协议的protobuf实现跨语言调用服务。
```go
import (
	"fmt"
	"log"
	"sync"

	"github.com/ansurfen/yock/scheduler"
)

func main() {
	var wg sync.WaitGroup
	ys := yocks.New()
	lib := ys.CreateLib("log")
	lib.SetField(map[string]any{
		"Info": func(msg string) {
			log.Println(msg)
		},
		"Infof": func(msg string, a ...any) {
			log.Printf(msg, a...)
		},
	})
	if err := ys.Eval(`log.Info("Hello World!")`); err != nil {
		panic(err)
	}
	go ys.EventLoop()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ys.Do(func() {
				ys.Eval(fmt.Sprintf(`log.Info("%d")`, i))
				wg.Done()
			})
		}(i)
	}
	wg.Wait()
}
```

yocku: 为yock提供crypto、ssh、random等基础功能。以下为util包实现的基于权重的负载均衡器例子。
```go
import (
	"fmt"

	"github.com/ansurfen/yock/util"
)

func main() {
	testset := map[string]int{
		"apple": 0,
		"banana": 0,
		"coconut": 0,
	}
	elements := []string{}
	for k := range testset {
		elements = append(elements, k)
	}
	lb := util.NewWeightedRandom(elements)
	for i := 0; i < 100; i++ {
		e, idx := lb.Next()
		testset[e]++
		lb.Up(idx)
	}
	fmt.Println(testset, lb.Weights())
}
```

yocke: 开箱即用的配置文件模块，能够在用户目录下快速创建环境。此外，yocke还提供了环境变量的操作以及对registry(windows)和plist(darwin)的抽象。你能够在`docs/yocke`下面查看详细信息。
```go
import (
	"fmt"

	yocke "github.com/ansurfen/yock/env"
)

type YockeTestConf struct {
	Author  string `yaml:"author"`
	Version string `yaml:"version"`
}

func main() {
	// defer yocke.FreeEnv[YockeTestConf]()
	yocke.InitEnv(&yocke.EnvOpt[YockeTestConf]{
		Workdir: ".yocke",
		Subdirs: []string{"tmp", "log", "mnt", "unmnt"},
		Conf:    YockeTestConf{},
	})
	env := yocke.GetEnv[YockeTestConf]()
	defer env.Save()
	env.SetValue("author", "a")
	fmt.Println(env.Viper())
}
```

ycho: yock的日志模块，基于`zap`实现了zlog(支持日志分割)以及`bubbletea`实现了tlog(tui日志)等ycho实例。
```go
import "github.com/ansurfen/yock/ycho"

func init() {
	zlog, err := ycho.NewZLog(ycho.YchoOpt{
		Stdout: true,
	})
	if err != nil {
		panic(err)
	}
	ycho.SetYcho(zlog)
}

func main() {
	ycho.Info("Hello World!")
	ycho.Fatalf("1 == 2 -> %v", 1 == 2)
}
```

yockf: 基于[libffi](https://github.com/libffi/libffi)封装的跨语言调用服务，你能够在Go语言中方便的调用动态库。
```go
import (
	"fmt"

	"github.com/ansurfen/yock/ffi"
)

func main() {
	mylib, err := ffi.NewLibray("libmylib.dll")
	if err != nil {
		panic(err)
	}
	hello := mylib.Func("hello", "void", []string{})
	hello()
	hello2 := mylib.Func("hello2", "str", []string{"str", "int"})
	fmt.Println(hello2("ansurfen", int32(60)))
}
```

## 文档

你可以在[这里](/docs/)查看关于模块开发以及yock开发相关的信息。

## 中央仓库

如果你想要将模块登记到yock，以便他能够使用标识符索引取代url，可以看这里。

## 未来计划

- [ ] 实现MOCK服务以及中间件编排（YockCloud）
- [ ] 实现预处理以及dsl增强lua语法

## 协议

这个软件被构建在MIT协议之下，详情请查看 [LICENSE](../../LICENSE) 。