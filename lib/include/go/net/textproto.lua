-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

textproto = {}

--- TrimBytes returns b without leading and trailing ASCII space.
---@param b byte[]
---@return byte[]
function textproto.TrimBytes(b) end

--- NewWriter returns a new Writer writing to w.
---@param w bufioWriter
---@return flateWriter
function textproto.NewWriter(w) end

--- NewReader returns a new Reader reading from r.
---
--- To avoid denial of service attacks, the provided bufio.Reader
--- should be reading from an io.LimitReader or similar Reader to bound
--- the size of responses.
---@param r bufioReader
---@return flateReader
function textproto.NewReader(r) end

--- CanonicalMIMEHeaderKey returns the canonical format of the
--- MIME header key s. The canonicalization converts the first
--- letter and any letter following a hyphen to upper case;
--- the rest are converted to lowercase. For example, the
--- canonical key for "accept-encoding" is "Accept-Encoding".
--- MIME header keys are assumed to be ASCII only.
--- If s contains a space or invalid header field bytes, it is
--- returned without modifications.
---@param s string
---@return string
function textproto.CanonicalMIMEHeaderKey(s) end

--- NewConn returns a new Conn using conn for I/O.
---@param conn ioReadWriteCloser
---@return netConn
function textproto.NewConn(conn) end

--- Dial connects to the given address on the given network using net.Dial
--- and then returns a new Conn for the connection.
---@param network string
---@param addr string
---@return netConn, err
function textproto.Dial(network, addr) end

--- TrimString returns s without leading and trailing ASCII space.
---@param s string
---@return string
function textproto.TrimString(s) end

--- A Writer implements convenience methods for writing
--- requests or responses to a text protocol network connection.
---@class flateWriter
---@field W any
local flateWriter = {}

--- PrintfLine writes the formatted output followed by \r\n.
---@param format string
---@vararg any
---@return err
function flateWriter:PrintfLine(format, ...) end

--- DotWriter returns a writer that can be used to write a dot-encoding to w.
--- It takes care of inserting leading dots when necessary,
--- translating line-ending \n into \r\n, and adding the final .\r\n line
--- when the DotWriter is closed. The caller should close the
--- DotWriter before the next call to a method on w.
---
--- See the documentation for Reader's DotReader method for details about dot-encoding.
---@return any
function flateWriter:DotWriter() end

--- A MIMEHeader represents a MIME-style header mapping
--- keys to sets of values.
---@class textprotoMIMEHeader
local textprotoMIMEHeader = {}

--- Values returns all values associated with the given key.
--- It is case insensitive; CanonicalMIMEHeaderKey is
--- used to canonicalize the provided key. To use non-canonical
--- keys, access the map directly.
--- The returned slice is not a copy.
---@param key string
---@return string[]
function textprotoMIMEHeader:Values(key) end

--- Del deletes the values associated with key.
---@param key string
function textprotoMIMEHeader:Del(key) end

--- Add adds the key, value pair to the header.
--- It appends to any existing values associated with key.
---@param key string
---@param value string
function textprotoMIMEHeader:Add(key, value) end

--- Set sets the header entries associated with key to
--- the single element value. It replaces any existing
--- values associated with key.
---@param key string
---@param value string
function textprotoMIMEHeader:Set(key, value) end

--- Get gets the first value associated with the given key.
--- It is case insensitive; CanonicalMIMEHeaderKey is used
--- to canonicalize the provided key.
--- If there are no values associated with the key, Get returns "".
--- To use non-canonical keys, access the map directly.
---@param key string
---@return string
function textprotoMIMEHeader:Get(key) end

--- A Pipeline manages a pipelined in-order request/response sequence.
---
--- To use a Pipeline p to manage multiple clients on a connection,
--- each client should run:
---
---	id := p.Next()	// take a number
---
---	p.StartRequest(id)	// wait for turn to send request
---	«send request»
---	p.EndRequest(id)	// notify Pipeline that request is sent
---
---	p.StartResponse(id)	// wait for turn to read response
---	«read response»
---	p.EndResponse(id)	// notify Pipeline that response is read
---
--- A pipelined server can use the same calls to ensure that
--- responses computed in parallel are written in the correct order.
---@class textprotoPipeline
local textprotoPipeline = {}

--- Next returns the next id for a request/response pair.
---@return any
function textprotoPipeline:Next() end

--- StartRequest blocks until it is time to send (or, if this is a server, receive)
--- the request with the given id.
---@param id any
function textprotoPipeline:StartRequest(id) end

--- EndRequest notifies p that the request with the given id has been sent
--- (or, if this is a server, received).
---@param id any
function textprotoPipeline:EndRequest(id) end

--- StartResponse blocks until it is time to receive (or, if this is a server, send)
--- the request with the given id.
---@param id any
function textprotoPipeline:StartResponse(id) end

--- EndResponse notifies p that the response with the given id has been received
--- (or, if this is a server, sent).
---@param id any
function textprotoPipeline:EndResponse(id) end

--- A Reader implements convenience methods for reading requests
--- or responses from a text protocol network connection.
---@class flateReader
---@field R any
local flateReader = {}

--- ReadContinuedLine reads a possibly continued line from r,
--- eliding the final trailing ASCII white space.
--- Lines after the first are considered continuations if they
--- begin with a space or tab character. In the returned data,
--- continuation lines are separated from the previous line
--- only by a single space: the newline and leading white space
--- are removed.
---
--- For example, consider this input:
---
---	Line 1
---	  continued...
---	Line 2
---
--- The first call to ReadContinuedLine will return "Line 1 continued..."
--- and the second will return "Line 2".
---
--- Empty lines are never continued.
---@return string, err
function flateReader:ReadContinuedLine() end

--- ReadResponse reads a multi-line response of the form:
---
---	code-message line 1
---	code-message line 2
---	...
---	code message line n
---
--- where code is a three-digit status code. The first line starts with the
--- code and a hyphen. The response is terminated by a line that starts
--- with the same code followed by a space. Each line in message is
--- separated by a newline (\n).
---
--- See page 36 of RFC 959 (https://www.ietf.org/rfc/rfc959.txt) for
--- details of another form of response accepted:
---
---	code-message line 1
---	message line 2
---	...
---	code message line n
---
--- If the prefix of the status does not match the digits in expectCode,
--- ReadResponse returns with err set to &Error{code, message}.
--- For example, if expectCode is 31, an error will be returned if
--- the status is not in the range [310,319].
---
--- An expectCode <= 0 disables the check of the status code.
---@param expectCode number
---@return number, string, err
function flateReader:ReadResponse(expectCode) end

--- ReadDotLines reads a dot-encoding and returns a slice
--- containing the decoded lines, with the final \r\n or \n elided from each.
---
--- See the documentation for the DotReader method for details about dot-encoding.
---@return string[], err
function flateReader:ReadDotLines() end

--- ReadLine reads a single line from r,
--- eliding the final \n or \r\n from the returned string.
---@return string, err
function flateReader:ReadLine() end

--- ReadContinuedLineBytes is like ReadContinuedLine but
--- returns a []byte instead of a string.
---@return byte[], err
function flateReader:ReadContinuedLineBytes() end

--- ReadCodeLine reads a response code line of the form
---
---	code message
---
--- where code is a three-digit status code and the message
--- extends to the rest of the line. An example of such a line is:
---
---	220 plan9.bell-labs.com ESMTP
---
--- If the prefix of the status does not match the digits in expectCode,
--- ReadCodeLine returns with err set to &Error{code, message}.
--- For example, if expectCode is 31, an error will be returned if
--- the status is not in the range [310,319].
---
--- If the response is multi-line, ReadCodeLine returns an error.
---
--- An expectCode <= 0 disables the check of the status code.
---@param expectCode number
---@return number, string, err
function flateReader:ReadCodeLine(expectCode) end

--- ReadDotBytes reads a dot-encoding and returns the decoded data.
---
--- See the documentation for the DotReader method for details about dot-encoding.
---@return byte[], err
function flateReader:ReadDotBytes() end

--- ReadLineBytes is like ReadLine but returns a []byte instead of a string.
---@return byte[], err
function flateReader:ReadLineBytes() end

--- ReadMIMEHeader reads a MIME-style header from r.
--- The header is a sequence of possibly continued Key: Value lines
--- ending in a blank line.
--- The returned map m maps CanonicalMIMEHeaderKey(key) to a
--- sequence of values in the same order encountered in the input.
---
--- For example, consider this input:
---
---	My-Key: Value 1
---	Long-Key: Even
---	       Longer Value
---	My-Key: Value 2
---
--- Given that input, ReadMIMEHeader returns the map:
---
---	map[string][]string{
---		"My-Key": {"Value 1", "Value 2"},
---		"Long-Key": {"Even Longer Value"},
---	}
---@return textprotoMIMEHeader, err
function flateReader:ReadMIMEHeader() end

--- DotReader returns a new Reader that satisfies Reads using the
--- decoded text of a dot-encoded block read from r.
--- The returned Reader is only valid until the next call
--- to a method on r.
---
--- Dot encoding is a common framing used for data blocks
--- in text protocols such as SMTP.  The data consists of a sequence
--- of lines, each of which ends in "\r\n".  The sequence itself
--- ends at a line containing just a dot: ".\r\n".  Lines beginning
--- with a dot are escaped with an additional dot to avoid
--- looking like the end of the sequence.
---
--- The decoded form returned by the Reader's Read method
--- rewrites the "\r\n" line endings into the simpler "\n",
--- removes leading dot escapes if present, and stops with error io.EOF
--- after consuming (and discarding) the end-of-sequence line.
---@return ioReader
function flateReader:DotReader() end

--- An Error represents a numeric error response from a server.
---@class execError
---@field Code number
---@field Msg string
local execError = {}


---@return string
function execError:Error() end

--- A ProtocolError describes a protocol violation such
--- as an invalid response or a hung-up connection.
---@class httpProtocolError
local httpProtocolError = {}


---@return string
function httpProtocolError:Error() end

--- A Conn represents a textual network protocol connection.
--- It consists of a Reader and Writer to manage I/O
--- and a Pipeline to sequence concurrent requests on the connection.
--- These embedded types carry methods with them;
--- see the documentation of those types for details.
---@class netConn
local netConn = {}

--- Close closes the connection.
---@return err
function netConn:Close() end

--- Cmd is a convenience method that sends a command after
--- waiting its turn in the pipeline. The command text is the
--- result of formatting format with args and appending \r\n.
--- Cmd returns the id of the command, for use with StartResponse and EndResponse.
---
--- For example, a client might run a HELP command that returns a dot-body
--- by using:
---
---	id, err := c.Cmd("HELP")
---	if err != nil {
---		return nil, err
---	}
---
---	c.StartResponse(id)
---	defer c.EndResponse(id)
---
---	if _, _, err = c.ReadCodeLine(110); err != nil {
---		return nil, err
---	}
---	text, err := c.ReadDotBytes()
---	if err != nil {
---		return nil, err
---	}
---	return c.ReadCodeLine(250)
---@param format string
---@vararg any
---@return any, err
function netConn:Cmd(format, ...) end
