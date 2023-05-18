# Include

[English](../../../lib/include/README.md) | 简体中文

include目录包含着yock提供的lib定义，他会在`yock new`创建新模块的时候引入，实现类型提示。同时，为了支持国际化，代码中`{{.}}`的部分都会和/lang路径下面json文件通过模板替换实现I18N注释。

## 未来计划

将yock.lua拆分成各个模块，实现解耦。同时利用`auto`包下面的generator生成注释，针对go语言导给lua的函数。