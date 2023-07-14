-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

smtp = {}

--- PlainAuth returns an Auth that implements the PLAIN authentication
--- mechanism as defined in RFC 4616. The returned Auth uses the given
--- username and password to authenticate to host and act as identity.
--- Usually identity should be the empty string, to act as username.
---
--- PlainAuth will only send the credentials if the connection is using TLS
--- or is connected to localhost. Otherwise authentication will fail with an
--- error, without sending the credentials.
---@param identity string
---@param username string
---@param password string
---@param host string
---@return smtpAuth
function smtp.PlainAuth(identity, username, password, host) end

--- CRAMMD5Auth returns an Auth that implements the CRAM-MD5 authentication
--- mechanism as defined in RFC 2195.
--- The returned Auth uses the given username and secret to authenticate
--- to the server using the challenge-response mechanism.
---@param username string
---@param secret string
---@return smtpAuth
function smtp.CRAMMD5Auth(username, secret) end

--- Dial returns a new Client connected to an SMTP server at addr.
--- The addr must include a port, as in "mail.example.com:smtp".
---@param addr string
---@return httpClient, err
function smtp.Dial(addr) end

--- NewClient returns a new Client using an existing connection and host as a
--- server name to be used when authenticating.
---@param conn netConn
---@param host string
---@return httpClient, err
function smtp.NewClient(conn, host) end

--- SendMail connects to the server at addr, switches to TLS if
--- possible, authenticates with the optional mechanism a if possible,
--- and then sends an email from address from, to addresses to, with
--- message msg.
--- The addr must include a port, as in "mail.example.com:smtp".
---
--- The addresses in the to parameter are the SMTP RCPT addresses.
---
--- The msg parameter should be an RFC 822-style email with headers
--- first, a blank line, and then the message body. The lines of msg
--- should be CRLF terminated. The msg headers should usually include
--- fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
--- messages is accomplished by including an email address in the to
--- parameter but not including it in the msg headers.
---
--- The SendMail function and the net/smtp package are low-level
--- mechanisms and provide no support for DKIM signing, MIME
--- attachments (see the mime/multipart package), or other mail
--- functionality. Higher-level packages exist outside of the standard
--- library.
---@param addr string
---@param a smtpAuth
---@param from string
---@param to string[]
---@param msg byte[]
---@return err
function smtp.SendMail(addr, a, from, to, msg) end

--- Auth is implemented by an SMTP authentication mechanism.
---@class smtpAuth
local smtpAuth = {}

--- ServerInfo records information about an SMTP server.
---@class smtpServerInfo
---@field Name string
---@field TLS boolean
---@field Auth any
local smtpServerInfo = {}

--- A Client represents a client connection to an SMTP server.
---@class httpClient
---@field Text any
local httpClient = {}

--- StartTLS sends the STARTTLS command and encrypts all further communication.
--- Only servers that advertise the STARTTLS extension support this function.
---@param config tlsConfig
---@return err
function httpClient:StartTLS(config) end

--- Rcpt issues a RCPT command to the server using the provided email address.
--- A call to Rcpt must be preceded by a call to Mail and may be followed by
--- a Data call or another Rcpt call.
---@param to string
---@return err
function httpClient:Rcpt(to) end

--- Verify checks the validity of an email address on the server.
--- If Verify returns nil, the address is valid. A non-nil return
--- does not necessarily indicate an invalid address. Many servers
--- will not verify addresses for security reasons.
---@param addr string
---@return err
function httpClient:Verify(addr) end

--- Mail issues a MAIL command to the server using the provided email address.
--- If the server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME
--- parameter. If the server supports the SMTPUTF8 extension, Mail adds the
--- SMTPUTF8 parameter.
--- This initiates a mail transaction and is followed by one or more Rcpt calls.
---@param from string
---@return err
function httpClient:Mail(from) end

--- Close closes the connection.
---@return err
function httpClient:Close() end

--- Auth authenticates a client using the provided authentication mechanism.
--- A failed authentication closes the connection.
--- Only servers that advertise the AUTH extension support this function.
---@param a smtpAuth
---@return err
function httpClient:Auth(a) end

--- Data issues a DATA command to the server and returns a writer that
--- can be used to write the mail headers and body. The caller should
--- close the writer before calling any more methods on c. A call to
--- Data must be preceded by one or more calls to Rcpt.
---@return any, err
function httpClient:Data() end

--- Reset sends the RSET command to the server, aborting the current mail
--- transaction.
---@return err
function httpClient:Reset() end

--- Noop sends the NOOP command to the server. It does nothing but check
--- that the connection to the server is okay.
---@return err
function httpClient:Noop() end

--- Hello sends a HELO or EHLO to the server as the given host name.
--- Calling this method is only necessary if the client needs control
--- over the host name used. The client will introduce itself as "localhost"
--- automatically otherwise. If Hello is called, it must be called before
--- any of the other methods.
---@param localName string
---@return err
function httpClient:Hello(localName) end

--- TLSConnectionState returns the client's TLS connection state.
--- The return values are their zero values if StartTLS did
--- not succeed.
---@return any, boolean
function httpClient:TLSConnectionState() end

--- Extension reports whether an extension is support by the server.
--- The extension name is case-insensitive. If the extension is supported,
--- Extension also returns a string that contains any parameters the
--- server specifies for the extension.
---@param ext string
---@return boolean, string
function httpClient:Extension(ext) end

--- Quit sends the QUIT command and closes the connection to the server.
---@return err
function httpClient:Quit() end
