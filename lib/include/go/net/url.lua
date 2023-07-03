-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

url = {}

---{{.urlPathUnescape}}
---@param s string
---@return string, err
function url.PathUnescape(s)
end

---{{.urlPathEscape}}
---@param s string
---@return string
function url.PathEscape(s)
end

---{{.urlQueryEscape}}
---@param s string
---@return string
function url.QueryEscape(s)
end

---{{.urlJoinPath}}
---@param base string
---@vararg string
---@return string, err
function url.JoinPath(base, ...)
end

---{{.urlParseQuery}}
---@param query string
---@return urlValues, err
function url.ParseQuery(query)
end

---{{.urlParseRequestURI}}
---@param rawURL string
---@return urlURL, err
function url.ParseRequestURI(rawURL)
end

---{{.urlUserPassword}}
---@param username string
---@param password string
---@return urlUserinfo
function url.UserPassword(username, password)
end

---{{.urlUser}}
---@param username string
---@return urlUserinfo
function url.User(username)
end

---{{.urlParse}}
---@param rawURL string
---@return urlURL, err
function url.Parse(rawURL)
end

---{{.urlQueryUnescape}}
---@param s string
---@return string, err
function url.QueryUnescape(s)
end

---@class urlValues
local urlValues = {}

---{{.urlValuesSet}}
---@param key string
---@param value string
function urlValues:Set(key, value)
end

---{{.urlValuesAdd}}
---@param key string
---@param value string
function urlValues:Add(key, value)
end

---{{.urlValuesDel}}
---@param key string
function urlValues:Del(key)
end

---{{.urlValuesHas}}
---@param key string
---@return boolean
function urlValues:Has(key)
end

---{{.urlValuesEncode}}
---@return string
function urlValues:Encode()
end

---{{.urlValuesGet}}
---@param key string
---@return string
function urlValues:Get(key)
end

---@class netError
---@field Op string
---@field URL string
---@field Err err
local netError = {}

---{{.netErrorTemporary}}
---@return boolean
function netError:Temporary()
end

---{{.netErrorUnwrap}}
---@return err
function netError:Unwrap()
end

---{{.netErrorError}}
---@return string
function netError:Error()
end

---{{.netErrorTimeout}}
---@return boolean
function netError:Timeout()
end

---@class urlEscapeError
local urlEscapeError = {}

---{{.urlEscapeErrorError}}
---@return string
function urlEscapeError:Error()
end

---@class urlInvalidHostError
local urlInvalidHostError = {}

---{{.urlInvalidHostErrorError}}
---@return string
function urlInvalidHostError:Error()
end

---@class urlURL
---@field Scheme string
---@field Opaque string
---@field User urlUserinfo
---@field Host string
---@field Path string
---@field RawPath string
---@field OmitHost boolean
---@field ForceQuery boolean
---@field RawQuery string
---@field Fragment string
---@field RawFragment string
local urlURL = {}

---{{.urlURLEscapedPath}}
---@return string
function urlURL:EscapedPath()
end

---{{.urlURLRedacted}}
---@return string
function urlURL:Redacted()
end

---{{.urlURLIsAbs}}
---@return boolean
function urlURL:IsAbs()
end

---{{.urlURLParse}}
---@param ref string
---@return urlURL, err
function urlURL:Parse(ref)
end

---{{.urlURLString}}
---@return string
function urlURL:String()
end

---{{.urlURLQuery}}
---@return urlValues
function urlURL:Query()
end

---{{.urlURLEscapedFragment}}
---@return string
function urlURL:EscapedFragment()
end

---{{.urlURLRequestURI}}
---@return string
function urlURL:RequestURI()
end

---{{.urlURLPort}}
---@return string
function urlURL:Port()
end

---{{.urlURLJoinPath}}
---@vararg string
---@return urlURL
function urlURL:JoinPath(...)
end

---{{.urlURLResolveReference}}
---@param ref urlURL
---@return urlURL
function urlURL:ResolveReference(ref)
end

---{{.urlURLHostname}}
---@return string
function urlURL:Hostname()
end

---{{.urlURLMarshalBinary}}
---@return byte[], err
function urlURL:MarshalBinary()
end

---{{.urlURLUnmarshalBinary}}
---@param text byte[]
---@return err
function urlURL:UnmarshalBinary(text)
end

---@class urlUserinfo
local urlUserinfo = {}

---{{.urlUserinfoUsername}}
---@return string
function urlUserinfo:Username()
end

---{{.urlUserinfoPassword}}
---@return string, boolean
function urlUserinfo:Password()
end

---{{.urlUserinfoString}}
---@return string
function urlUserinfo:String()
end
