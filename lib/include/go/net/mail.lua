-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class mail
---@field ErrHeaderNotPresent any
mail = {}

--- ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
---@param address string
---@return mailAddress, err
function mail.ParseAddress(address) end

--- ReadMessage reads a message from r.
--- The headers are parsed, and the body of the message will be available
--- for reading from msg.Body.
---@param r ioReader
---@return mailMessage, err
function mail.ReadMessage(r) end

--- ParseDate parses an RFC 5322 date string.
---@param date string
---@return any, err
function mail.ParseDate(date) end

--- ParseAddressList parses the given string as a list of addresses.
---@param list string
---@return any, err
function mail.ParseAddressList(list) end

--- A Header represents the key-value pairs in a mail message header.
---@class gzipHeader
local gzipHeader = {}

--- Get gets the first value associated with the given key.
--- It is case insensitive; CanonicalMIMEHeaderKey is used
--- to canonicalize the provided key.
--- If there are no values associated with the key, Get returns "".
--- To access multiple values of a key, or to use non-canonical keys,
--- access the map directly.
---@param key string
---@return string
function gzipHeader:Get(key) end

--- Date parses the Date header field.
---@return any, err
function gzipHeader:Date() end

--- AddressList parses the named header field as a list of addresses.
---@param key string
---@return any, err
function gzipHeader:AddressList(key) end

--- Address represents a single mail address.
--- An address such as "Barry Gibbs <bg@example.com>" is represented
--- as Address{Name: "Barry Gibbs", Address: "bg@example.com"}.
---@class mailAddress
---@field Name string
---@field Address string
local mailAddress = {}

--- String formats the address as a valid RFC 5322 address.
--- If the address's name contains non-ASCII characters
--- the name will be rendered according to RFC 2047.
---@return string
function mailAddress:String() end

--- An AddressParser is an RFC 5322 address parser.
---@class mailAddressParser
---@field WordDecoder any
local mailAddressParser = {}

--- Parse parses a single RFC 5322 address of the
--- form "Gogh Fir <gf@example.com>" or "foo@example.com".
---@param address string
---@return mailAddress, err
function mailAddressParser:Parse(address) end

--- ParseList parses the given string as a list of comma-separated addresses
--- of the form "Gogh Fir <gf@example.com>" or "foo@example.com".
---@param list string
---@return any, err
function mailAddressParser:ParseList(list) end

--- A Message represents a parsed mail message.
---@class mailMessage
---@field Header gzipHeader
---@field Body any
local mailMessage = {}
