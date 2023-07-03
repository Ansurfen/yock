```lua
curl({
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

Parameter
* method set type of http request，GET, POST，DEL ....
* data set content of body
* save set true to save package
* debug print output
* header set headers of request