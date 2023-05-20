# Ylang

[English](../../../ylang/README.md) | 简体中文

Ylang 是对lua语法的增强，主要通过预处理实现，最终ylang会被翻译成lua语法。

## 语法

#### 定义/获取变量

```
a = 10 // a = 10
_a = 10 // local a = 10
echo a, _a // print(a, _a)
```

#### 类

```
class user { a = "", b = "" } // 等价于 local user = {a = "", b = ""}

class admin extends user {
    c = 10
} // local admin = {a = "", b = "", c = 10}
```

#### 流式传递

```
class user {
    id: number
    name: string
    pwd: string
}
_user{name: "ansurfen"} >> mysql.query >> (user, ok) => {
    if #user.name > 6 {
        return "user_" + user.id, user.name
    }
    return nil
} >> redis.zset.put
// 每个 >> 就相当于一个函数调用， mysql.query(user{name:"ansurfen"})
// () 用于接受上一层函数的返回值 local user, ok = mysql.query(user{name:"ansurfen"})
// => 是lambda表达式，可以把他理解成立即调用的函数
// 最终会被翻译成:
// local _user = {name = "ansurfen"}
// local user, ok = mysql.query(_user)
// local _a = function(user, ok) if #user.name > 6 then return "user_" .. user.id, user.name end return nil end
// redis.zset.put(_a(user, ok))
```