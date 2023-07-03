-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class flate
---@field NoCompression any
---@field BestSpeed any
---@field BestCompression any
---@field DefaultCompression any
---@field HuffmanOnly any
flate = {}

---{{.flateNewWriter}}
---@param w ioWriter
---@param level number
---@return flateWriter, err
function flate.NewWriter(w, level)
end

---{{.flateNewWriterDict}}
---@param w ioWriter
---@param level number
---@param dict byte[]
---@return flateWriter, err
function flate.NewWriterDict(w, level, dict)
end

---{{.flateNewReader}}
---@param r ioReader
---@return any
function flate.NewReader(r)
end

---{{.flateNewReaderDict}}
---@param r ioReader
---@param dict byte[]
---@return any
function flate.NewReaderDict(r, dict)
end

---@class flateWriter
local flateWriter = {}

---{{.flateWriterReset}}
---@param dst ioWriter
function flateWriter:Reset(dst)
end

---{{.flateWriterWrite}}
---@param data byte[]
---@return number, err
function flateWriter:Write(data)
end

---{{.flateWriterFlush}}
---@return err
function flateWriter:Flush()
end

---{{.flateWriterClose}}
---@return err
function flateWriter:Close()
end

---@class flateReader
local flateReader = {}

---@class flateCorruptInputError
local flateCorruptInputError = {}

---{{.flateCorruptInputErrorError}}
---@return string
function flateCorruptInputError:Error()
end

---@class flateInternalError
local flateInternalError = {}

---{{.flateInternalErrorError}}
---@return string
function flateInternalError:Error()
end

---@class flateReadError
---@field Offset number
---@field Err err
local flateReadError = {}

---{{.flateReadErrorError}}
---@return string
function flateReadError:Error()
end

---@class flateWriteError
---@field Offset number
---@field Err err
local flateWriteError = {}

---{{.flateWriteErrorError}}
---@return string
function flateWriteError:Error()
end

---@class flateResetter
local flateResetter = {}
