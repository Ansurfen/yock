#  <img src="../static/yock.ico" width = "60" height = "60" alt="logo" align=center />Yock

[![Go Report Card](https://goreportcard.com/badge/github.com/ansurfen/cushion)](https://goreportcard.com/report/github.com/ansurfen/yock)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/ansurfen/yock?status.svg)](https://pkg.go.dev/github.com/ansurfen/yock)
[![codecov](https://codecov.io/gh/Ansurfen/yock/branch/main/graph/badge.svg?token=UHYKJTT80P)](https://codecov.io/gh/Ansurfen/yock)
[![Discord](https://img.shields.io/badge/chat-on_discord-7289da)](https://discord.gg/vdybzz8RJn)

[English](../../README.md) | 简体中文

Yock 是一个跨平台的分布式构建流编排解决方案。它能够作为软件包使用，就像Homebrew, rpm, winget等等。同时它还能充当编程语言的依赖管理的角色（pip，npm，maven等等）。在此基础上，yock还基于grpc和协程实现分布式构建任务（甚至可以为此搭建集群）。你可以将他视作nodejs框架的lua语言版本，不同的是他专注于编排，更加轻量。

## 架构
![arch](../static/arch.png)

* Yctl: 负责调度yock的命令。
* Yockp: 主要用于对lua文件的预处理，例如模式分解，将一份lua代码根据给定的模式分解成多份lua文件供分布式运行。
* Yocks: 调度器负责运行lua代码，以task为单位起协程执行。
* YPM: yock包管理，负责补全和装载依赖。
* Yockd: yock的守护进程，负责跨进程和跨端通信，构建P2P或中心化集群。
* Yockr: yock的运行时。
* Yockw: yock的监控，用于日志查询、指标监控。
* Ycho: yock的日志模块，用于呈现运行时信息。

## 安装

你能够在[这里](https://github.com/Ansurfen/yock/releases)下载二进制版本，或者尝试以下几种方式。
`注意`: 下载完后还需要将yock挂载到本地环境中。在解压压缩包后，进入可执行文件的目录运行`yock run install.lua`完成这个过程。

#### 包管理 (版本更新存在滞后)
npm: npm install @ansurfen/yock -g

pip: pip install yock

#### 使用yock构建

Yock实现了类似"自举"的操作，这意味着它能够自己构建自己。当然，这一切的前提还需要go语言的编译器。

```cmd
git clone https://github.com/Ansurfen/yock.git

cd ctl

./build.bat/sh //正常构建
./build.bat/sh ffi //带 libffi 构建 (需要 gcc 或 mingw)
./build.bat/sh dev //构建开发版本
./build.bat/sh oslinux //交叉编译到linux平台

// 自动构建出带libffi版本的项目，当上一步完成后
yock run install.lua
yock run ../auto/build-ffi.lua
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

yockc: 为yock提供了一系列简单的GNU命令，例如iptabls, curl, crontab等等。你能够直接调用它们，具体的Opt字段可以前往`/docs/yockc`查看详情。
```go
import yockc "github.com/ansurfen/yock/cmd"

func main() {
	yockc.Curl(yockc.CurlOpt{
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

## 快速开始
下面是一些简单的例子，你也能够在`lib/test`中查看测试内容以了解函数的使用。

GNU命令: yock首要的目标就是实现跨平台的构建方案以代替bat和bash脚本。因此，你能够在脚本中以函数的形式使用gnu命令。除了这个例子外，你也可以参考`auto/build.lua`（yock的构建文件）以及上文所说的`lib/test/gnu_test.lua`。
```lua
mkdir("a")
cd("a")
write("test.txt", "Hello World!")
local data = cat("test.txt")
print(data)
cd("..")
rm({safe = false}, "/a")
```

任务调度: yock能够并发执行异步命令，你能够使用`yock run test.lua echo echo2`或者`yock run test.lua all`的形式调用。
```lua
-- test.lua
job_option({
    strict = true,
    debug = true,
})

job("echo", function(ctx)
	print("echo")
end)

job("echo2", function(ctx)
	print("echo2")
end)

jobs("all", "echo", "echo2")
```

协程与信号量: 得益于go语言的特性，yock也继承了go协程的能力，也就意味着lua不再是单线程的语言，它能够实现真正的异步任务。同时，yock还提供了notify和wait函数应对异步转同步的需求。
```lua
go(function()
    local i = 0
    while true do
        time.Sleep(1 * time.Second)
        if i > 3 then
            notify("x")
        end
        print("do task1")
        i = i + 1
    end
end)

go(function()
    local i = 0
    while true do
        if i == 5 then
            wait("x")
        end
        print("do task2")
        i = i + 1
    end
end)

time.Sleep(8 * time.Second)
```

HTTP服务: yock继承了go语言大部分标准库，因此你能够用类似go的语法快速开始。
```lua
http.HandleFunc("/", function(w, req)
    fmt.Fprintf(w, "Hello World!\n")
end)
http.ListenAndServe(":8080", nil)
```

YPM: 当执行完`yock run install.lua`后，便会全局注册包管理工具。你能够用他安装yock的模块。
```cmd
<!-- 列出全部命令 -->
ypm

<!-- 全局安装模块 -->
ypm install ark -g
```

## 文档

你可以在[这里](https://ansurfen.github.io/YockNav/)查看关于模块开发以及yock开发相关的信息。

## 中央仓库

如果你想要将模块登记到yock，以便他能够使用标识符索引取代url，可以看[这里](https://github.com/Ansurfen/yock-todo)。

## 未来计划

- [ ] 实现MOCK服务以及中间件编排（YockCloud）
- [ ] 实现预处理以及dsl增强lua语法

## 协议

这个软件被构建在MIT协议之下，详情请查看 [LICENSE](../../LICENSE) 。