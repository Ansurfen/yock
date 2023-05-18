# Include

English | [简体中文](../../docs/zh_cn/lib/include.md)

The include directory contains the lib definition provided by yock, which it'll introduce when `yock new` create a new module, implementing type tips. At the same time, in order to support internationalization, the code `{{.}}`'s section and the json file below the /lang path implement I18N comments through template substitution.

## TODO

Split the yock.lua into modules to achieve decoupling. Meanwhile, use the generator under the `auto` package to generate comments for the functions introduced to lua in golang.