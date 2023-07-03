-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class mail
---@field ErrHeaderNotPresent any
mail = {}

---{{.mailReadMessage}}
---@param r ioReader
---@return mailMessage, err
function mail.ReadMessage(r)
end

---{{.mailParseAddress}}
---@param address string
---@return mailAddress, err
function mail.ParseAddress(address)
end

---{{.mailParseAddressList}}
---@param list string
---@return any, err
function mail.ParseAddressList(list)
end

---{{.mailParseDate}}
---@param date string
---@return timeTime, err
function mail.ParseDate(date)
end

---@class mailMessage
---@field Header httpHeader
---@field Body any
local mailMessage = {}

---@class httpHeader
local httpHeader = {}

---{{.httpHeaderGet}}
---@param key string
---@return string
function httpHeader:Get(key)
end

---{{.httpHeaderDate}}
---@return timeTime, err
function httpHeader:Date()
end

---{{.httpHeaderAddressList}}
---@param key string
---@return any, err
function httpHeader:AddressList(key)
end

---@class mailAddress
---@field Name string
---@field Address string
local mailAddress = {}

---{{.mailAddressString}}
---@return string
function mailAddress:String()
end

---@class mailAddressParser
---@field WordDecoder any
local mailAddressParser = {}

---{{.mailAddressParserParse}}
---@param address string
---@return mailAddress, err
function mailAddressParser:Parse(address)
end

---{{.mailAddressParserParseList}}
---@param list string
---@return any, err
function mailAddressParser:ParseList(list)
end
