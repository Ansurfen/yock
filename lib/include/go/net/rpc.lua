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

---{{.rpcServeRequest}}
---@param codec rpcServerCodec
---@return err
function rpc.ServeRequest(codec)
end

---{{.rpcServeConn}}
---@param conn ioReadWriteCloser
function rpc.ServeConn(conn)
end

---{{.rpcNewClientWithCodec}}
---@param codec rpcClientCodec
---@return httpClient
function rpc.NewClientWithCodec(codec)
end

---{{.rpcDialHTTP}}
---@param network string
---@param address string
---@return httpClient, err
function rpc.DialHTTP(network, address)
end

---{{.rpcHandleHTTP}}
function rpc.HandleHTTP()
end

---{{.rpcNewServer}}
---@return httpServer
function rpc.NewServer()
end

---{{.rpcServeCodec}}
---@param codec rpcServerCodec
function rpc.ServeCodec(codec)
end

---{{.rpcAccept}}
---@param lis netListener
function rpc.Accept(lis)
end

---{{.rpcRegister}}
---@param rcvr any
---@return err
function rpc.Register(rcvr)
end

---{{.rpcDialHTTPPath}}
---@param network string
---@param address string
---@param path string
---@return httpClient, err
function rpc.DialHTTPPath(network, address, path)
end

---{{.rpcDial}}
---@param network string
---@param address string
---@return httpClient, err
function rpc.Dial(network, address)
end

---{{.rpcNewClient}}
---@param conn ioReadWriteCloser
---@return httpClient
function rpc.NewClient(conn)
end

---{{.rpcRegisterName}}
---@param name string
---@param rcvr any
---@return err
function rpc.RegisterName(name, rcvr)
end

---@class httpRequest
---@field ServiceMethod string
---@field Seq number
local httpRequest = {}

---@class httpResponse
---@field ServiceMethod string
---@field Seq number
---@field Error string
local httpResponse = {}

---@class httpServer
local httpServer = {}

---{{.httpServerServeHTTP}}
---@param w httpResponseWriter
---@param req httpRequest
function httpServer:ServeHTTP(w, req)
end

---{{.httpServerHandleHTTP}}
---@param rpcPath string
---@param debugPath string
function httpServer:HandleHTTP(rpcPath, debugPath)
end

---{{.httpServerServeConn}}
---@param conn ioReadWriteCloser
function httpServer:ServeConn(conn)
end

---{{.httpServerServeCodec}}
---@param codec rpcServerCodec
function httpServer:ServeCodec(codec)
end

---{{.httpServerServeRequest}}
---@param codec rpcServerCodec
---@return err
function httpServer:ServeRequest(codec)
end

---{{.httpServerRegister}}
---@param rcvr any
---@return err
function httpServer:Register(rcvr)
end

---{{.httpServerRegisterName}}
---@param name string
---@param rcvr any
---@return err
function httpServer:RegisterName(name, rcvr)
end

---{{.httpServerAccept}}
---@param lis netListener
function httpServer:Accept(lis)
end

---@class rpcServerError
local rpcServerError = {}

---{{.rpcServerErrorError}}
---@return string
function rpcServerError:Error()
end

---@class rpcCall
---@field ServiceMethod string
---@field Args any
---@field Reply any
---@field Error err
---@field Done any
local rpcCall = {}

---@class httpClient
local httpClient = {}

---{{.httpClientClose}}
---@return err
function httpClient:Close()
end

---{{.httpClientGo}}
---@param serviceMethod string
---@param args any
---@param reply any
---@param done any
---@return rpcCall
function httpClient:Go(serviceMethod, args, reply, done)
end

---{{.httpClientCall}}
---@param serviceMethod string
---@param args any
---@param reply any
---@return err
function httpClient:Call(serviceMethod, args, reply)
end

---@class rpcClientCodec
local rpcClientCodec = {}

---@class rpcServerCodec
local rpcServerCodec = {}
