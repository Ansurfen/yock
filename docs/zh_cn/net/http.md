```lua
http({
    method = "GET",
    data = "Hello World",
    save = true,
    debug = true,
    header = {
        ["Content-Type"] = "application/json;charset=utf-8",
        ["jwt"] = "..."
    },
    cookie = {
        name = "aaa",
        value = "aaa-value"
    },
})
```

参数：
* method 设置HTTP请求的类型，GET, POST，DEL ....
* data 设置body内容
* save 为 true 时保存报文
* debug 打印输出
* header 设置数据包头部