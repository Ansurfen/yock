# bash

[English](../../../lib/bash/README.md) | 简体中文

这是一个实验性的项目，皆在提供跨平台的bash解释器。值得注意的是，它不会被打包进yock里面，而是作为module分发。

在yock脚本里面引入
```lua
local sh = import("bash")
-- 以字符串的形式载入
sh([[
    url="github.com/ansurfen/cushion/yock.load.sh"
    mkdir -p e/b
    cp ./a ./b
    echo $url
    echo $GOPATH | echo
    rmdir c
    rmdir d
    rm -r a
    rm -r b
    rm -r d
    curl https://www.github.com | echo
    whoami | echo
    pwd | echo
    touch a.txt
    curl -o . -O https://repo.spring.io/ui/native/release/org/springframework/boot/spring-boot-cli/1.4.3.RELEASE/spring-boot-cli-1.4.3.RELEASE-bin.zip
]])
-- 以文件的形式载入
sh("#load yock.load.sh")
```