-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class rpc
---@field DefaultRPCPath any
---@field DefaultDebugPath any
---@field ErrShutdown any
---@field DefaultServer any
rpc = {}

--- NewClient returns a new Client to handle requests to the
--- set of services at the other end of the connection.
--- It adds a buffer to the write side of the connection so
--- the header and payload are sent as a unit.
---
--- The read and write halves of the connection are serialized independently,
--- so no interlocking is required. However each half may be accessed
--- concurrently so the implementation of conn should protect against
--- concurrent reads or concurrent writes.
---@param conn ioReadWriteCloser
---@return httpClient
function rpc.NewClient(conn) end

--- NewClientWithCodec is like NewClient but uses the specified
--- codec to encode requests and decode responses.
---@param codec rpcClientCodec
---@return httpClient
function rpc.NewClientWithCodec(codec) end

--- DialHTTPPath connects to an HTTP RPC server
--- at the specified network address and path.
---@param network string
---@param address string
---@param path string
---@return httpClient, err
function rpc.DialHTTPPath(network, address, path) end

--- HandleHTTP registers an HTTP handler for RPC messages to DefaultServer
--- on DefaultRPCPath and a debugging handler on DefaultDebugPath.
--- It is still necessary to invoke http.Serve(), typically in a go statement.
function rpc.HandleHTTP() end

--- NewServer returns a new Server.
---@return httpServer
function rpc.NewServer() end

--- ServeRequest is like ServeCodec but synchronously serves a single request.
--- It does not close the codec upon completion.
---@param codec rpcServerCodec
---@return err
function rpc.ServeRequest(codec) end

--- Accept accepts connections on the listener and serves requests
--- to DefaultServer for each incoming connection.
--- Accept blocks; the caller typically invokes it in a go statement.
---@param lis netListener
function rpc.Accept(lis) end

--- DialHTTP connects to an HTTP RPC server at the specified network address
--- listening on the default HTTP RPC path.
---@param network string
---@param address string
---@return httpClient, err
function rpc.DialHTTP(network, address) end

--- Dial connects to an RPC server at the specified network address.
---@param network string
---@param address string
---@return httpClient, err
function rpc.Dial(network, address) end

--- RegisterName is like Register but uses the provided name for the type
--- instead of the receiver's concrete type.
---@param name string
---@param rcvr any
---@return err
function rpc.RegisterName(name, rcvr) end

--- ServeConn runs the DefaultServer on a single connection.
--- ServeConn blocks, serving the connection until the client hangs up.
--- The caller typically invokes ServeConn in a go statement.
--- ServeConn uses the gob wire format (see package gob) on the
--- connection. To use an alternate codec, use ServeCodec.
--- See NewClient's comment for information about concurrent access.
---@param conn ioReadWriteCloser
function rpc.ServeConn(conn) end

--- Register publishes the receiver's methods in the DefaultServer.
---@param rcvr any
---@return err
function rpc.Register(rcvr) end

--- ServeCodec is like ServeConn but uses the specified codec to
--- decode requests and encode responses.
---@param codec rpcServerCodec
function rpc.ServeCodec(codec) end

--- Request is a header written before every RPC call. It is used internally
--- but documented here as an aid to debugging, such as when analyzing
--- network traffic.
---@class httpRequest
---@field ServiceMethod string
---@field Seq number
local httpRequest = {}

--- Response is a header written before every RPC return. It is used internally
--- but documented here as an aid to debugging, such as when analyzing
--- network traffic.
---@class httpResponse
---@field ServiceMethod string
---@field Seq number
---@field Error string
local httpResponse = {}

--- ServerError represents an error that has been returned from
--- the remote side of the RPC connection.
---@class rpcServerError
local rpcServerError = {}


---@return string
function rpcServerError:Error() end

--- Call represents an active RPC.
---@class rpcCall
---@field ServiceMethod string
---@field Args any
---@field Reply any
---@field Error err
---@field Done any
local rpcCall = {}

--- Client represents an RPC Client.
--- There may be multiple outstanding Calls associated
--- with a single Client, and a Client may be used by
--- multiple goroutines simultaneously.
---@class httpClient
local httpClient = {}

--- Close calls the underlying codec's Close method. If the connection is already
--- shutting down, ErrShutdown is returned.
---@return err
function httpClient:Close() end

--- Go invokes the function asynchronously. It returns the Call structure representing
--- the invocation. The done channel will signal when the call is complete by returning
--- the same Call object. If done is nil, Go will allocate a new channel.
--- If non-nil, done must be buffered or Go will deliberately crash.
---@param serviceMethod string
---@param args any
---@param reply any
---@param done any
---@return rpcCall
function httpClient:Go(serviceMethod, args, reply, done) end

--- Call invokes the named function, waits for it to complete, and returns its error status.
---@param serviceMethod string
---@param args any
---@param reply any
---@return err
function httpClient:Call(serviceMethod, args, reply) end

--- A ClientCodec implements writing of RPC requests and
--- reading of RPC responses for the client side of an RPC session.
--- The client calls WriteRequest to write a request to the connection
--- and calls ReadResponseHeader and ReadResponseBody in pairs
--- to read responses. The client calls Close when finished with the
--- connection. ReadResponseBody may be called with a nil
--- argument to force the body of the response to be read and then
--- discarded.
--- See NewClient's comment for information about concurrent access.
---@class rpcClientCodec
local rpcClientCodec = {}

--- Server represents an RPC Server.
---@class httpServer
local httpServer = {}

--- RegisterName is like Register but uses the provided name for the type
--- instead of the receiver's concrete type.
---@param name string
---@param rcvr any
---@return err
function httpServer:RegisterName(name, rcvr) end

--- ServeRequest is like ServeCodec but synchronously serves a single request.
--- It does not close the codec upon completion.
---@param codec rpcServerCodec
---@return err
function httpServer:ServeRequest(codec) end

--- Accept accepts connections on the listener and serves requests
--- for each incoming connection. Accept blocks until the listener
--- returns a non-nil error. The caller typically invokes Accept in a
--- go statement.
---@param lis netListener
function httpServer:Accept(lis) end

--- Register publishes in the server the set of methods of the
--- receiver value that satisfy the following conditions:
---   - exported method of exported type
---   - two arguments, both of exported type
---   - the second argument is a pointer
---   - one return value, of type error
---
--- It returns an error if the receiver is not an exported type or has
--- no suitable methods. It also logs the error using package log.
--- The client accesses each method using a string of the form "Type.Method",
--- where Type is the receiver's concrete type.
---@param rcvr any
---@return err
function httpServer:Register(rcvr) end

--- ServeConn runs the server on a single connection.
--- ServeConn blocks, serving the connection until the client hangs up.
--- The caller typically invokes ServeConn in a go statement.
--- ServeConn uses the gob wire format (see package gob) on the
--- connection. To use an alternate codec, use ServeCodec.
--- See NewClient's comment for information about concurrent access.
---@param conn ioReadWriteCloser
function httpServer:ServeConn(conn) end

--- ServeCodec is like ServeConn but uses the specified codec to
--- decode requests and encode responses.
---@param codec rpcServerCodec
function httpServer:ServeCodec(codec) end

--- ServeHTTP implements an http.Handler that answers RPC requests.
---@param w httpResponseWriter
---@param req httpRequest
function httpServer:ServeHTTP(w, req) end

--- HandleHTTP registers an HTTP handler for RPC messages on rpcPath,
--- and a debugging handler on debugPath.
--- It is still necessary to invoke http.Serve(), typically in a go statement.
---@param rpcPath string
---@param debugPath string
function httpServer:HandleHTTP(rpcPath, debugPath) end

--- A ServerCodec implements reading of RPC requests and writing of
--- RPC responses for the server side of an RPC session.
--- The server calls ReadRequestHeader and ReadRequestBody in pairs
--- to read requests from the connection, and it calls WriteResponse to
--- write a response back. The server calls Close when finished with the
--- connection. ReadRequestBody may be called with a nil
--- argument to force the body of the request to be read and discarded.
--- See NewClient's comment for information about concurrent access.
---@class rpcServerCodec
local rpcServerCodec = {}
