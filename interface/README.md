# Yock Interface

English | [简体中文](../docs/zh_cn/yocki/README.md)

Yock Interface provides a mesh interface based on grpc and protobuf to help developers build Yock extensions quickly and easily. It's somewhat similar to microservices and requires constant local loopbacks. But the benefits are self-evident, and the use of networking supports more languages and is more convenient than intrusive (dynamic libraries).

## API

#### Ping

Parameter: empty
Return value: empty
Function: Used to test whether the connection is reachable

#### Call

Parameters: Fn (string), Args (string)
Return value: Buf (string)
Function: The Call function is the core of the Interface, which receives the function name (Fn) and parameters (Args) to return Buf. This means that it can meet a large number of needs, through the free combination of function names and parameters, just like calling functions within the program, the return value can be in the serialized JSON, XML and other formats.