-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

textproto = {}

---{{.textprotoCanonicalMIMEHeaderKey}}
---@param s string
---@return string
function textproto.CanonicalMIMEHeaderKey(s)
end

---{{.textprotoTrimBytes}}
---@param b byte[]
---@return byte[]
function textproto.TrimBytes(b)
end

---{{.textprotoNewConn}}
---@param conn ioReadWriteCloser
---@return netConn
function textproto.NewConn(conn)
end

---{{.textprotoDial}}
---@param network string
---@param addr string
---@return netConn, err
function textproto.Dial(network, addr)
end

---{{.textprotoTrimString}}
---@param s string
---@return string
function textproto.TrimString(s)
end

---{{.textprotoNewWriter}}
---@param w bufioWriter
---@return textprotoWriter
function textproto.NewWriter(w)
end

---{{.textprotoNewReader}}
---@param r bufioReader
---@return textprotoReader
function textproto.NewReader(r)
end

---@class textprotoWriter
---@field W any
local textprotoWriter = {}

---{{.textprotoWriterPrintfLine}}
---@param format string
---@vararg any
---@return err
function textprotoWriter:PrintfLine(format, ...)
end

---{{.textprotoWriterDotWriter}}
---@return any
function textprotoWriter:DotWriter()
end

---@class textprotoMIMEHeader
local textprotoMIMEHeader = {}

---{{.textprotoMIMEHeaderAdd}}
---@param key string
---@param value string
function textprotoMIMEHeader:Add(key, value)
end

---{{.textprotoMIMEHeaderSet}}
---@param key string
---@param value string
function textprotoMIMEHeader:Set(key, value)
end

---{{.textprotoMIMEHeaderGet}}
---@param key string
---@return string
function textprotoMIMEHeader:Get(key)
end

---{{.textprotoMIMEHeaderValues}}
---@param key string
---@return string[]
function textprotoMIMEHeader:Values(key)
end

---{{.textprotoMIMEHeaderDel}}
---@param key string
function textprotoMIMEHeader:Del(key)
end

---@class textprotoPipeline
local textprotoPipeline = {}

---{{.textprotoPipelineStartResponse}}
---@param id any
function textprotoPipeline:StartResponse(id)
end

---{{.textprotoPipelineEndResponse}}
---@param id any
function textprotoPipeline:EndResponse(id)
end

---{{.textprotoPipelineNext}}
---@return any
function textprotoPipeline:Next()
end

---{{.textprotoPipelineStartRequest}}
---@param id any
function textprotoPipeline:StartRequest(id)
end

---{{.textprotoPipelineEndRequest}}
---@param id any
function textprotoPipeline:EndRequest(id)
end

---@class textprotoReader
---@field R any
local textprotoReader = {}

---{{.textprotoReaderReadContinuedLineBytes}}
---@return byte[], err
function textprotoReader:ReadContinuedLineBytes()
end

---{{.textprotoReaderReadCodeLine}}
---@param expectCode number
---@return number, string, err
function textprotoReader:ReadCodeLine(expectCode)
end

---{{.textprotoReaderReadDotBytes}}
---@return byte[], err
function textprotoReader:ReadDotBytes()
end

---{{.textprotoReaderReadDotLines}}
---@return string[], err
function textprotoReader:ReadDotLines()
end

---{{.textprotoReaderReadLineBytes}}
---@return byte[], err
function textprotoReader:ReadLineBytes()
end

---{{.textprotoReaderDotReader}}
---@return any
function textprotoReader:DotReader()
end

---{{.textprotoReaderReadLine}}
---@return string, err
function textprotoReader:ReadLine()
end

---{{.textprotoReaderReadResponse}}
---@param expectCode number
---@return number, string, err
function textprotoReader:ReadResponse(expectCode)
end

---{{.textprotoReaderReadContinuedLine}}
---@return string, err
function textprotoReader:ReadContinuedLine()
end

---{{.textprotoReaderReadMIMEHeader}}
---@return textprotoMIMEHeader, err
function textprotoReader:ReadMIMEHeader()
end

---@class netError
---@field Code number
---@field Msg string
local netError = {}

---{{.netErrorError}}
---@return string
function netError:Error()
end

---@class httpProtocolError
local httpProtocolError = {}

---{{.httpProtocolErrorError}}
---@return string
function httpProtocolError:Error()
end

---@class netConn
local netConn = {}

---{{.netConnClose}}
---@return err
function netConn:Close()
end

---{{.netConnCmd}}
---@param format string
---@vararg any
---@return any, err
function netConn:Cmd(format, ...)
end
