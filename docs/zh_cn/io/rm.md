```lua
rm({
    safe = false,
    pattern = ".*"
}, "aaa")
```

参数
* safe 默认值为 true，表示不允许递归删除文件
* pattern 根据正则表达式规则匹配文件删除