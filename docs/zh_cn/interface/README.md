# Yock Interface

[English](../../../interface/README.md) | 简体中文

Yock Interface基于grpc和protobuf提供网格化的接口，帮助开发者快速更方便的构建Yock拓展。它有点类似于微服务，需要不断的本地回环。但是这样的好处也是不言而喻的，利用网络化，相比侵入式（动态库）支持的语言也更多，也更方便。

## API

#### Ping

参数：空
返回值：空
功能：用于测试连接是否可达

#### Call

参数：Fn (string), Args (string)
返回值: Buf (string)
功能：Call函数是Interface的核心，他接收函数名（Fn）和参数（Args）返回Buf。这意味着他能够满足大量的需求，通过函数名和参数的自由组合，就像调用程序内的函数一样，返回值可以是序列化后的JSON、XML等格式。