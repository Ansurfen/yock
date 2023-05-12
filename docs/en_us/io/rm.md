```lua
rm({
    safe = false,
    pattern = ".*"
}, "aaa")
```

Parameter
* safe set true in default，and no allow to delete file with recurse
* pattern match file to delete according to regexp