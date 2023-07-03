-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class http
---@field SameSiteDefaultMode any
---@field SameSiteLaxMode any
---@field SameSiteStrictMode any
---@field SameSiteNoneMode any
---@field MethodGet any
---@field MethodHead any
---@field MethodPost any
---@field MethodPut any
---@field MethodPatch any
---@field MethodDelete any
---@field MethodConnect any
---@field MethodOptions any
---@field MethodTrace any
---@field TrailerPrefix any
---@field DefaultMaxHeaderBytes any
---@field TimeFormat any
---@field StateNew any
---@field StateActive any
---@field StateIdle any
---@field StateHijacked any
---@field StateClosed any
---@field StatusContinue any
---@field StatusSwitchingProtocols any
---@field StatusProcessing any
---@field StatusEarlyHints any
---@field StatusOK any
---@field StatusCreated any
---@field StatusAccepted any
---@field StatusNonAuthoritativeInfo any
---@field StatusNoContent any
---@field StatusResetContent any
---@field StatusPartialContent any
---@field StatusMultiStatus any
---@field StatusAlreadyReported any
---@field StatusIMUsed any
---@field StatusMultipleChoices any
---@field StatusMovedPermanently any
---@field StatusFound any
---@field StatusSeeOther any
---@field StatusNotModified any
---@field StatusUseProxy any
---@field StatusTemporaryRedirect any
---@field StatusPermanentRedirect any
---@field StatusBadRequest any
---@field StatusUnauthorized any
---@field StatusPaymentRequired any
---@field StatusForbidden any
---@field StatusNotFound any
---@field StatusMethodNotAllowed any
---@field StatusNotAcceptable any
---@field StatusProxyAuthRequired any
---@field StatusRequestTimeout any
---@field StatusConflict any
---@field StatusGone any
---@field StatusLengthRequired any
---@field StatusPreconditionFailed any
---@field StatusRequestEntityTooLarge any
---@field StatusRequestURITooLong any
---@field StatusUnsupportedMediaType any
---@field StatusRequestedRangeNotSatisfiable any
---@field StatusExpectationFailed any
---@field StatusTeapot any
---@field StatusMisdirectedRequest any
---@field StatusUnprocessableEntity any
---@field StatusLocked any
---@field StatusFailedDependency any
---@field StatusTooEarly any
---@field StatusUpgradeRequired any
---@field StatusPreconditionRequired any
---@field StatusTooManyRequests any
---@field StatusRequestHeaderFieldsTooLarge any
---@field StatusUnavailableForLegalReasons any
---@field StatusInternalServerError any
---@field StatusNotImplemented any
---@field StatusBadGateway any
---@field StatusServiceUnavailable any
---@field StatusGatewayTimeout any
---@field StatusHTTPVersionNotSupported any
---@field StatusVariantAlsoNegotiates any
---@field StatusInsufficientStorage any
---@field StatusLoopDetected any
---@field StatusNotExtended any
---@field StatusNetworkAuthenticationRequired any
---@field DefaultMaxIdleConnsPerHost any
---@field DefaultClient any
---@field ErrUseLastResponse any
---@field NoBody any
---@field ErrMissingFile any
---@field ErrNotSupported any
---@field ErrUnexpectedTrailer any
---@field ErrMissingBoundary any
---@field ErrNotMultipart any
---@field ErrHeaderTooLong any
---@field ErrShortBody any
---@field ErrMissingContentLength any
---@field ErrNoCookie any
---@field ErrNoLocation any
---@field ErrBodyNotAllowed any
---@field ErrHijacked any
---@field ErrContentLength any
---@field ErrWriteAfterFlush any
---@field ServerContextKey any
---@field LocalAddrContextKey any
---@field ErrAbortHandler any
---@field DefaultServeMux any
---@field ErrServerClosed any
---@field ErrHandlerTimeout any
---@field ErrLineTooLong any
---@field ErrBodyReadAfterClose any
---@field DefaultTransport any
---@field ErrSkipAltProtocol any
http = {}

---@return userdata
function http.Client()
end

---{{.httpProxyFromEnvironment}}
---@param req httpRequest
---@return any, err
function http.ProxyFromEnvironment(req)
end

---{{.httpPost}}
---@param url string
---@param contentType string
---@param body ioReader
---@return httpResponse, err
function http.Post(url, contentType, body)
end

---{{.httpServeContent}}
---@param w httpResponseWriter
---@param req httpRequest
---@param name string
---@param modtime timeTime
---@param content ioReadSeeker
function http.ServeContent(w, req, name, modtime, content)
end

---{{.httpNewRequestWithContext}}
---@param ctx contextContext
---@param method string
---@param url string
---@param body ioReader
---@return httpRequest, err
function http.NewRequestWithContext(ctx, method, url, body)
end

---{{.httpNewResponseController}}
---@param rw httpResponseWriter
---@return httpResponseController
function http.NewResponseController(rw)
end

---{{.httpServeTLS}}
---@param l netListener
---@param handler httpHandler
---@param certFile string
---@param keyFile string
---@return err
function http.ServeTLS(l, handler, certFile, keyFile)
end

---{{.httpNewServeMux}}
---@return httpServeMux
function http.NewServeMux()
end

---{{.httpServe}}
---@param l netListener
---@param handler httpHandler
---@return err
function http.Serve(l, handler)
end

---{{.httpDetectContentType}}
---@param data byte[]
---@return string
function http.DetectContentType(data)
end

---{{.httpHead}}
---@param url string
---@return httpResponse, err
function http.Head(url)
end

---{{.httpServeFile}}
---@param w httpResponseWriter
---@param r httpRequest
---@param name string
function http.ServeFile(w, r, name)
end

---{{.httpReadRequest}}
---@param b bufioReader
---@return httpRequest, err
function http.ReadRequest(b)
end

---{{.httpParseHTTPVersion}}
---@param vers string
---@return number, boolean
function http.ParseHTTPVersion(vers)
end

---{{.httpAllowQuerySemicolons}}
---@param h httpHandler
---@return httpHandler
function http.AllowQuerySemicolons(h)
end

---{{.httpHandle}}
---@param pattern string
---@param handler httpHandler
function http.Handle(pattern, handler)
end

---{{.httpProxyURL}}
---@param fixedURL urlURL
---@return any
function http.ProxyURL(fixedURL)
end

---{{.httpFS}}
---@param fsys fsFS
---@return httpFileSystem
function http.FS(fsys)
end

---{{.httpCanonicalHeaderKey}}
---@param s string
---@return string
function http.CanonicalHeaderKey(s)
end

---{{.httpLogger}}
---@param w httpResponseWriter
---@param req httpRequest
function http.Logger(w, req)
end

---{{.httpNotFound}}
---@param w httpResponseWriter
---@param r httpRequest
function http.NotFound(w, r)
end

---{{.httpStripPrefix}}
---@param prefix string
---@param h httpHandler
---@return httpHandler
function http.StripPrefix(prefix, h)
end

---{{.httpListenAndServe}}
---@param addr string
---@param handler httpHandler
---@return err
function http.ListenAndServe(addr, handler)
end

---{{.httpTimeoutHandler}}
---@param h httpHandler
---@param dt timeDuration
---@param msg string
---@return httpHandler
function http.TimeoutHandler(h, dt, msg)
end

---{{.httpChanCreate}}
---@return httpChan
function http.ChanCreate()
end

---{{.httpNewFileTransport}}
---@param fs httpFileSystem
---@return httpRoundTripper
function http.NewFileTransport(fs)
end

---{{.httpFileServer}}
---@param root httpFileSystem
---@return httpHandler
function http.FileServer(root)
end

---{{.httpParseTime}}
---@param text string
---@return timeTime, err
function http.ParseTime(text)
end

---{{.httpMaxBytesHandler}}
---@param h httpHandler
---@param n number
---@return httpHandler
function http.MaxBytesHandler(h, n)
end

---{{.httpError}}
---@param w httpResponseWriter
---@param error string
---@param code number
function http.Error(w, error, code)
end

---{{.httpHelloServer}}
---@param w httpResponseWriter
---@param req httpRequest
function http.HelloServer(w, req)
end

---{{.httpGet}}
---@param url string
---@return httpResponse, err
function http.Get(url)
end

---{{.httpSetCookie}}
---@param w httpResponseWriter
---@param cookie httpCookie
function http.SetCookie(w, cookie)
end

---{{.httpListenAndServeTLS}}
---@param addr string
---@param certFile string
---@param keyFile string
---@param handler httpHandler
---@return err
function http.ListenAndServeTLS(addr, certFile, keyFile, handler)
end

---{{.httpHandleFunc}}
---@param pattern string
---@param handler function
function http.HandleFunc(pattern, handler)
end

---{{.httpFlagServer}}
---@param w httpResponseWriter
---@param req httpRequest
function http.FlagServer(w, req)
end

---{{.httpMaxBytesReader}}
---@param w httpResponseWriter
---@param r ioReadCloser
---@param n number
---@return ioReadCloser
function http.MaxBytesReader(w, r, n)
end

---{{.httpNewRequest}}
---@param method string
---@param url string
---@param body ioReader
---@return httpRequest, err
function http.NewRequest(method, url, body)
end

---{{.httpNotFoundHandler}}
---@return httpHandler
function http.NotFoundHandler()
end

---{{.httpStatusText}}
---@param code number
---@return string
function http.StatusText(code)
end

---{{.httpPostForm}}
---@param url string
---@param data urlValues
---@return httpResponse, err
function http.PostForm(url, data)
end

---{{.httpReadResponse}}
---@param r bufioReader
---@param req httpRequest
---@return httpResponse, err
function http.ReadResponse(r, req)
end

---{{.httpRedirectHandler}}
---@param url string
---@param code number
---@return httpHandler
function http.RedirectHandler(url, code)
end

---{{.httpRedirect}}
---@param w httpResponseWriter
---@param r httpRequest
---@param url string
---@param code number
function http.Redirect(w, r, url, code)
end

---{{.httpArgServer}}
---@param w httpResponseWriter
---@param req httpRequest
function http.ArgServer(w, req)
end

---{{.httpDateServer}}
---@param rw httpResponseWriter
---@param req httpRequest
function http.DateServer(rw, req)
end

---@class httpFlusher
local httpFlusher = {}

---@class httpResponse
---@field Status string
---@field StatusCode number
---@field Proto string
---@field ProtoMajor number
---@field ProtoMinor number
---@field Header httpHeader
---@field Body any
---@field ContentLength number
---@field TransferEncoding any
---@field Close boolean
---@field Uncompressed boolean
---@field Trailer httpHeader
---@field Request httpRequest
---@field TLS any
local httpResponse = {}

---{{.httpResponseCookies}}
---@return any
function httpResponse:Cookies()
end

---{{.httpResponseLocation}}
---@return urlURL, err
function httpResponse:Location()
end

---{{.httpResponseProtoAtLeast}}
---@param major number
---@param minor number
---@return boolean
function httpResponse:ProtoAtLeast(major, minor)
end

---{{.httpResponseWrite}}
---@param w ioWriter
---@return err
function httpResponse:Write(w)
end

---@class httpDir
local httpDir = {}

---{{.httpDirOpen}}
---@param name string
---@return httpFile, err
function httpDir:Open(name)
end

---@class httpMaxBytesError
---@field Limit number
local httpMaxBytesError = {}

---{{.httpMaxBytesErrorError}}
---@return string
function httpMaxBytesError:Error()
end

---@class httpCounter
local httpCounter = {}

---{{.httpCounterString}}
---@return string
function httpCounter:String()
end

---{{.httpCounterServeHTTP}}
---@param w httpResponseWriter
---@param req httpRequest
function httpCounter:ServeHTTP(w, req)
end

---@class httpFile
local httpFile = {}

---@class httpPushOptions
---@field Method string
---@field Header httpHeader
local httpPushOptions = {}

---@class httpTransport
---@field Proxy any
---@field OnProxyConnectResponse any
---@field DialContext any
---@field Dial any
---@field DialTLSContext any
---@field DialTLS any
---@field TLSClientConfig any
---@field TLSHandshakeTimeout any
---@field DisableKeepAlives boolean
---@field DisableCompression boolean
---@field MaxIdleConns number
---@field MaxIdleConnsPerHost number
---@field MaxConnsPerHost number
---@field IdleConnTimeout any
---@field ResponseHeaderTimeout any
---@field ExpectContinueTimeout any
---@field TLSNextProto any
---@field ProxyConnectHeader httpHeader
---@field GetProxyConnectHeader any
---@field MaxResponseHeaderBytes number
---@field WriteBufferSize number
---@field ReadBufferSize number
---@field ForceAttemptHTTP2 boolean
local httpTransport = {}

---{{.httpTransportClone}}
---@return httpTransport
function httpTransport:Clone()
end

---{{.httpTransportCancelRequest}}
---@param req httpRequest
function httpTransport:CancelRequest(req)
end

---{{.httpTransportCloseIdleConnections}}
function httpTransport:CloseIdleConnections()
end

---{{.httpTransportRegisterProtocol}}
---@param scheme string
---@param rt httpRoundTripper
function httpTransport:RegisterProtocol(scheme, rt)
end

---@class httpHandler
local httpHandler = {}

---@class httpSameSite
local httpSameSite = {}

---@class httpCookie
---@field Name string
---@field Value string
---@field Path string
---@field Domain string
---@field Expires any
---@field RawExpires string
---@field MaxAge number
---@field Secure boolean
---@field HttpOnly boolean
---@field SameSite httpSameSite
---@field Raw string
---@field Unparsed any
local httpCookie = {}

---{{.httpCookieString}}
---@return string
function httpCookie:String()
end

---{{.httpCookieValid}}
---@return err
function httpCookie:Valid()
end

---@class httpConnState
local httpConnState = {}

---{{.httpConnStateString}}
---@return string
function httpConnState:String()
end

---@class httpChan
local httpChan = {}

---{{.httpChanServeHTTP}}
---@param w httpResponseWriter
---@param req httpRequest
function httpChan:ServeHTTP(w, req)
end

---@class httpClient
---@field Transport httpRoundTripper
---@field CheckRedirect any
---@field Jar httpCookieJar
---@field Timeout any
local httpClient = {}

---{{.httpClientDo}}
---@param req httpRequest
---@return httpResponse, err
function httpClient:Do(req)
end

---{{.httpClientGet}}
---@param url string
---@return httpResponse, err
function httpClient:Get(url)
end

---{{.httpClientPost}}
---@param url string
---@param contentType string
---@param body ioReader
---@return httpResponse, err
function httpClient:Post(url, contentType, body)
end

---{{.httpClientPostForm}}
---@param url string
---@param data urlValues
---@return httpResponse, err
function httpClient:PostForm(url, data)
end

---{{.httpClientHead}}
---@param url string
---@return httpResponse, err
function httpClient:Head(url)
end

---{{.httpClientCloseIdleConnections}}
function httpClient:CloseIdleConnections()
end

---@class httpProtocolError
---@field ErrorString string
local httpProtocolError = {}

---{{.httpProtocolErrorError}}
---@return string
function httpProtocolError:Error()
end

---@class httpResponseController
local httpResponseController = {}

---{{.httpResponseControllerSetWriteDeadline}}
---@param deadline timeTime
---@return err
function httpResponseController:SetWriteDeadline(deadline)
end

---{{.httpResponseControllerFlush}}
---@return err
function httpResponseController:Flush()
end

---{{.httpResponseControllerHijack}}
---@return netConn, any, err
function httpResponseController:Hijack()
end

---{{.httpResponseControllerSetReadDeadline}}
---@param deadline timeTime
---@return err
function httpResponseController:SetReadDeadline(deadline)
end

---@class httpRequest
---@field Method string
---@field URL any
---@field Proto string
---@field ProtoMajor number
---@field ProtoMinor number
---@field Header httpHeader
---@field Body any
---@field GetBody any
---@field ContentLength number
---@field TransferEncoding any
---@field Close boolean
---@field Host string
---@field Form any
---@field PostForm any
---@field MultipartForm any
---@field Trailer httpHeader
---@field RemoteAddr string
---@field RequestURI string
---@field TLS any
---@field Cancel any
---@field Response httpResponse
local httpRequest = {}

---{{.httpRequestContext}}
---@return contextContext
function httpRequest:Context()
end

---{{.httpRequestParseMultipartForm}}
---@param maxMemory number
---@return err
function httpRequest:ParseMultipartForm(maxMemory)
end

---{{.httpRequestWriteProxy}}
---@param w ioWriter
---@return err
function httpRequest:WriteProxy(w)
end

---{{.httpRequestAddCookie}}
---@param c httpCookie
function httpRequest:AddCookie(c)
end

---{{.httpRequestUserAgent}}
---@return string
function httpRequest:UserAgent()
end

---{{.httpRequestMultipartReader}}
---@return any, err
function httpRequest:MultipartReader()
end

---{{.httpRequestFormFile}}
---@param key string
---@return any, any, err
function httpRequest:FormFile(key)
end

---{{.httpRequestCookies}}
---@return any
function httpRequest:Cookies()
end

---{{.httpRequestParseForm}}
---@return err
function httpRequest:ParseForm()
end

---{{.httpRequestClone}}
---@param ctx contextContext
---@return httpRequest
function httpRequest:Clone(ctx)
end

---{{.httpRequestWrite}}
---@param w ioWriter
---@return err
function httpRequest:Write(w)
end

---{{.httpRequestFormValue}}
---@param key string
---@return string
function httpRequest:FormValue(key)
end

---{{.httpRequestPostFormValue}}
---@param key string
---@return string
function httpRequest:PostFormValue(key)
end

---{{.httpRequestWithContext}}
---@param ctx contextContext
---@return httpRequest
function httpRequest:WithContext(ctx)
end

---{{.httpRequestProtoAtLeast}}
---@param major number
---@param minor number
---@return boolean
function httpRequest:ProtoAtLeast(major, minor)
end

---{{.httpRequestSetBasicAuth}}
---@param username string
---@param password string
function httpRequest:SetBasicAuth(username, password)
end

---{{.httpRequestCookie}}
---@param name string
---@return httpCookie, err
function httpRequest:Cookie(name)
end

---{{.httpRequestReferer}}
---@return string
function httpRequest:Referer()
end

---{{.httpRequestBasicAuth}}
---@return string, boolean
function httpRequest:BasicAuth()
end

---@class httpServeMux
local httpServeMux = {}

---{{.httpServeMuxServeHTTP}}
---@param w httpResponseWriter
---@param r httpRequest
function httpServeMux:ServeHTTP(w, r)
end

---{{.httpServeMuxHandle}}
---@param pattern string
---@param handler httpHandler
function httpServeMux:Handle(pattern, handler)
end

---{{.httpServeMuxHandleFunc}}
---@param pattern string
---@param handler any
function httpServeMux:HandleFunc(pattern, handler)
end

---{{.httpServeMuxHandler}}
---@param r httpRequest
---@return httpHandler, string
function httpServeMux:Handler(r)
end

---@class httpCloseNotifier
local httpCloseNotifier = {}

---@class httpPusher
local httpPusher = {}

---@class httpResponseWriter
local httpResponseWriter = {}

---@class any
---@field Addr string
---@field Handler httpHandler
---@field DisableGeneralOptionsHandler boolean
---@field TLSConfig any
---@field ReadTimeout any
---@field ReadHeaderTimeout any
---@field WriteTimeout any
---@field IdleTimeout any
---@field MaxHeaderBytes number
---@field TLSNextProto any
---@field ConnState any
---@field ErrorLog any
---@field BaseContext any
---@field ConnContext any
local any = {}

---{{.anyServeTLS}}
---@param l netListener
---@param certFile string
---@param keyFile string
---@return err
function any:ServeTLS(l, certFile, keyFile)
end

---{{.anyListenAndServeTLS}}
---@param certFile string
---@param keyFile string
---@return err
function any:ListenAndServeTLS(certFile, keyFile)
end

---{{.anyClose}}
---@return err
function any:Close()
end

---{{.anyServe}}
---@param l netListener
---@return err
function any:Serve(l)
end

---{{.anySetKeepAlivesEnabled}}
---@param v boolean
function any:SetKeepAlivesEnabled(v)
end

---{{.anyRegisterOnShutdown}}
---@param f any
function any:RegisterOnShutdown(f)
end

---{{.anyListenAndServe}}
---@return err
function any:ListenAndServe()
end

---{{.anyShutdown}}
---@param ctx contextContext
---@return err
function any:Shutdown(ctx)
end

---@class httpHijacker
local httpHijacker = {}

---@class httpHeader
local httpHeader = {}

---{{.httpHeaderAdd}}
---@param key string
---@param value string
function httpHeader:Add(key, value)
end

---{{.httpHeaderGet}}
---@param key string
---@return string
function httpHeader:Get(key)
end

---{{.httpHeaderClone}}
---@return httpHeader
function httpHeader:Clone()
end

---{{.httpHeaderDel}}
---@param key string
function httpHeader:Del(key)
end

---{{.httpHeaderWrite}}
---@param w ioWriter
---@return err
function httpHeader:Write(w)
end

---{{.httpHeaderWriteSubset}}
---@param w ioWriter
---@param exclude any
---@return err
function httpHeader:WriteSubset(w, exclude)
end

---{{.httpHeaderSet}}
---@param key string
---@param value string
function httpHeader:Set(key, value)
end

---{{.httpHeaderValues}}
---@param key string
---@return string[]
function httpHeader:Values(key)
end

---@class httpFileSystem
local httpFileSystem = {}

---@class httpCookieJar
local httpCookieJar = {}

---@class httpHandlerFunc
local httpHandlerFunc = {}

---{{.httpHandlerFuncServeHTTP}}
---@param w httpResponseWriter
---@param r httpRequest
function httpHandlerFunc:ServeHTTP(w, r)
end

---@class httpRoundTripper
local httpRoundTripper = {}
