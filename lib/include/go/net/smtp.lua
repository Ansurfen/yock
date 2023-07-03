-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

smtp = {}

---{{.smtpSendMail}}
---@param addr string
---@param a smtpAuth
---@param from string
---@param to string[]
---@param msg byte[]
---@return err
function smtp.SendMail(addr, a, from, to, msg)
end

---{{.smtpPlainAuth}}
---@param identity string
---@param username string
---@param password string
---@param host string
---@return smtpAuth
function smtp.PlainAuth(identity, username, password, host)
end

---{{.smtpCRAMMD5Auth}}
---@param username string
---@param secret string
---@return smtpAuth
function smtp.CRAMMD5Auth(username, secret)
end

---{{.smtpDial}}
---@param addr string
---@return httpClient, err
function smtp.Dial(addr)
end

---{{.smtpNewClient}}
---@param conn netConn
---@param host string
---@return httpClient, err
function smtp.NewClient(conn, host)
end

---@class smtpServerInfo
---@field Name string
---@field TLS boolean
---@field Auth any
local smtpServerInfo = {}

---@class smtpAuth
local smtpAuth = {}

---@class httpClient
---@field Text any
local httpClient = {}

---{{.httpClientVerify}}
---@param addr string
---@return err
function httpClient:Verify(addr)
end

---{{.httpClientReset}}
---@return err
function httpClient:Reset()
end

---{{.httpClientNoop}}
---@return err
function httpClient:Noop()
end

---{{.httpClientClose}}
---@return err
function httpClient:Close()
end

---{{.httpClientStartTLS}}
---@param config tlsConfig
---@return err
function httpClient:StartTLS(config)
end

---{{.httpClientTLSConnectionState}}
---@return any, boolean
function httpClient:TLSConnectionState()
end

---{{.httpClientRcpt}}
---@param to string
---@return err
function httpClient:Rcpt(to)
end

---{{.httpClientData}}
---@return any, err
function httpClient:Data()
end

---{{.httpClientExtension}}
---@param ext string
---@return boolean, string
function httpClient:Extension(ext)
end

---{{.httpClientQuit}}
---@return err
function httpClient:Quit()
end

---{{.httpClientHello}}
---@param localName string
---@return err
function httpClient:Hello(localName)
end

---{{.httpClientMail}}
---@param from string
---@return err
function httpClient:Mail(from)
end

---{{.httpClientAuth}}
---@param a smtpAuth
---@return err
function httpClient:Auth(a)
end
