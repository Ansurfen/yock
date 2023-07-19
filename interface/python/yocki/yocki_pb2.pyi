from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor


class CallRequest(_message.Message):
    __slots__ = ["Arg", "Fn"]
    ARG_FIELD_NUMBER: _ClassVar[int]
    Arg: str
    FN_FIELD_NUMBER: _ClassVar[int]
    Fn: str
    def __init__(self, Fn: _Optional[str] = ...,
                 Arg: _Optional[str] = ...) -> None: ...


class CallResponse(_message.Message):
    __slots__ = ["Buf"]
    BUF_FIELD_NUMBER: _ClassVar[int]
    Buf: str
    def __init__(self, Buf: _Optional[str] = ...) -> None: ...


class InfoRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...


class InfoResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...


class PingRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...


class PingResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...
