# bash

English | [简体中文](../../docs/zh_cn/lib/bash.md)

This is an experimental project that provides a cross-platform bash interpreter. It's worth nothing that it'll not be packaged into yock, but distributed as a module.

Import in yock script
```lua
local sh = import("bash")
-- load by string
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
-- load by file
sh("#load yock.load.sh")
```