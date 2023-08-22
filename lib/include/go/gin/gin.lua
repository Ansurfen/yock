-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class gin
---@field AuthUserKey any
---@field MIMEJSON any
---@field MIMEHTML any
---@field MIMEXML any
---@field MIMEXML2 any
---@field MIMEPlain any
---@field MIMEPOSTForm any
---@field MIMEMultipartPOSTForm any
---@field MIMEYAML any
---@field MIMETOML any
---@field BodyBytesKey any
---@field ContextKey any
---@field ErrorTypeBind any
---@field ErrorTypeRender any
---@field ErrorTypePrivate any
---@field ErrorTypePublic any
---@field ErrorTypeAny any
---@field ErrorTypeNu any
---@field PlatformGoogleAppEngine any
---@field PlatformCloudflare any
---@field EnvGinMode any
---@field DebugMode any
---@field ReleaseMode any
---@field TestMode any
---@field BindKey any
---@field Version any
---@field DebugPrintRouteFunc any
---@field DefaultWriter any
---@field DefaultErrorWriter any
gin = {}

---@alias ginContext any

--- WrapF is a helper function for wrapping http.HandlerFunc and returns a Gin middleware.
---@param f httpHandlerFunc
---@return ginHandlerFunc
function gin.WrapF(f) end

--- Dir returns a http.FileSystem that can be used by http.FileServer(). It is used internally
--- in router.Static().
--- if listDirectory == true, then it works the same as http.Dir() otherwise it returns
--- a filesystem that prevents http.FileServer() to list the directory files.
---@param root string
---@param listDirectory boolean
---@return any
function gin.Dir(root, listDirectory) end

--- CreateTestContextOnly returns a fresh context base on the engine for testing purposes
---@param w httpResponseWriter
---@param r ginEngine
---@return ginContext
function gin.CreateTestContextOnly(w, r) end

--- SetMode sets gin mode according to input string.
---@param value string
function gin.SetMode(value) end

--- ErrorLoggerT returns a HandlerFunc for a given error type.
---@param typ ginErrorType
---@return ginHandlerFunc
function gin.ErrorLoggerT(typ) end

--- LoggerWithConfig instance a Logger middleware with config.
---@param conf ginLoggerConfig
---@return ginHandlerFunc
function gin.LoggerWithConfig(conf) end

--- ForceConsoleColor force color output in the console.
function gin.ForceConsoleColor() end

--- CustomRecoveryWithWriter returns a middleware for a given writer that recovers from any panics and calls the provided handle func to handle it.
---@param out_ ioWriter
---@param handle ginRecoveryFunc
---@return ginHandlerFunc
function gin.CustomRecoveryWithWriter(out_, handle) end

--- WrapH is a helper function for wrapping http.Handler and returns a Gin middleware.
---@param h httpHandler
---@return ginHandlerFunc
function gin.WrapH(h) end

--- LoggerWithFormatter instance a Logger middleware with the specified log format function.
---@param f ginLogFormatter
---@return ginHandlerFunc
function gin.LoggerWithFormatter(f) end

--- LoggerWithWriter instance a Logger middleware with the specified writer buffer.
--- Example: os.Stdout, a file opened in write mode, a socket...
---@param out_ ioWriter
---@vararg string
---@return ginHandlerFunc
function gin.LoggerWithWriter(out_, ...) end

--- Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
--- By default, gin.DefaultWriter = os.Stdout.
---@return ginHandlerFunc
function gin.Logger() end

--- EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
--- call the DisallowUnknownFields method on the JSON Decoder instance.
function gin.EnableJsonDecoderDisallowUnknownFields() end

--- Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
---@return ginHandlerFunc
function gin.Recovery() end

--- CustomRecovery returns a middleware that recovers from any panics and calls the provided handle func to handle it.
---@param handle ginRecoveryFunc
---@return ginHandlerFunc
function gin.CustomRecovery(handle) end

--- BasicAuth returns a Basic HTTP Authorization middleware. It takes as argument a map[string]string where
--- the key is the user name and the value is the password.
---@param accounts ginAccounts
---@return ginHandlerFunc
function gin.BasicAuth(accounts) end

--- ErrorLogger returns a HandlerFunc for any error type.
---@return ginHandlerFunc
function gin.ErrorLogger() end

--- EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumber to
--- call the UseNumber method on the JSON Decoder instance.
function gin.EnableJsonDecoderUseNumber() end

--- Bind is a helper function for given interface object and returns a Gin middleware.
---@param val any
---@return ginHandlerFunc
function gin.Bind(val) end

--- Default returns an Engine instance with the Logger and Recovery middleware already attached.
---@return ginEngine
function gin.Default() end

--- DisableConsoleColor disables color output in the console.
function gin.DisableConsoleColor() end

--- DisableBindValidation closes the default validator.
function gin.DisableBindValidation() end

--- RecoveryWithWriter returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
---@param out_ ioWriter
---@vararg ginRecoveryFunc
---@return ginHandlerFunc
function gin.RecoveryWithWriter(out_, ...) end

--- CreateTestContext returns a fresh engine and context for testing purposes
---@param w httpResponseWriter
---@return ginContext, ginEngine
function gin.CreateTestContext(w) end

--- IsDebugging returns true if the framework is running in debug mode.
--- Use SetMode(gin.ReleaseMode) to disable debug mode.
---@return boolean
function gin.IsDebugging() end

--- New returns a new blank Engine instance without any middleware attached.
--- By default, the configuration is:
--- - RedirectTrailingSlash:  true
--- - RedirectFixedPath:      false
--- - HandleMethodNotAllowed: false
--- - ForwardedByClientIP:    true
--- - UseRawPath:             false
--- - UnescapePathValues:     true
---@return ginEngine
function gin.New() end

--- BasicAuthForRealm returns a Basic HTTP Authorization middleware. It takes as arguments a map[string]string where
--- the key is the user name and the value is the password, as well as the name of the Realm.
--- If the realm is empty, "Authorization Required" will be used by default.
--- (see http://tools.ietf.org/html/rfc2617#section-1.2)
---@param accounts ginAccounts
---@param realm string
---@return ginHandlerFunc
function gin.BasicAuthForRealm(accounts, realm) end

--- Mode returns current gin mode.
---@return string
function gin.Mode() end

--- HandlersChain defines a HandlerFunc slice.
---@class ginHandlersChain
local ginHandlersChain = {}

--- Last returns the last handler in the chain. i.e. the last handler is the main one.
---@return ginHandlerFunc
function ginHandlersChain:Last() end

--- Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
--- Create an instance of Engine, by using New() or Default()
---@class ginEngine
---@field RedirectTrailingSlash boolean
---@field RedirectFixedPath boolean
---@field HandleMethodNotAllowed boolean
---@field ForwardedByClientIP boolean
---@field AppEngine boolean
---@field UseRawPath boolean
---@field UnescapePathValues boolean
---@field RemoveExtraSlash boolean
---@field RemoteIPHeaders any
---@field TrustedPlatform string
---@field MaxMultipartMemory number
---@field UseH2C boolean
---@field ContextWithFallback boolean
---@field HTMLRender any
---@field FuncMap any
local ginEngine = {}

--- LoadHTMLFiles loads a slice of HTML files
--- and associates the result with HTML renderer.
---@vararg string
function ginEngine:LoadHTMLFiles(...) end

--- NoRoute adds handlers for NoRoute. It returns a 404 code by default.
---@vararg ginHandlerFunc
function ginEngine:NoRoute(...) end

--- RunListener attaches the router to a http.Server and starts listening and serving HTTP requests
--- through the specified net.Listener
---@param listener netListener
---@return err
function ginEngine:RunListener(listener) end

--- ServeHTTP conforms to the http.Handler interface.
---@param w httpResponseWriter
---@param req httpRequest
function ginEngine:ServeHTTP(w, req) end

--- LoadHTMLGlob loads HTML files identified by glob pattern
--- and associates the result with HTML renderer.
---@param pattern string
function ginEngine:LoadHTMLGlob(pattern) end

--- RunUnix attaches the router to a http.Server and starts listening and serving HTTP requests
--- through the specified unix socket (i.e. a file).
--- Note: this method will block the calling goroutine indefinitely unless an error happens.
---@param file string
---@return err
function ginEngine:RunUnix(file) end

--- RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
--- through the specified file descriptor.
--- Note: this method will block the calling goroutine indefinitely unless an error happens.
---@param fd number
---@return err
function ginEngine:RunFd(fd) end

--- SecureJsonPrefix sets the secureJSONPrefix used in Context.SecureJSON.
---@param prefix string
---@return ginEngine
function ginEngine:SecureJsonPrefix(prefix) end

--- RunTLS attaches the router to a http.Server and starts listening and serving HTTPS (secure) requests.
--- It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
--- Note: this method will block the calling goroutine indefinitely unless an error happens.
---@param addr string
---@param certFile string
---@param keyFile string
---@return err
function ginEngine:RunTLS(addr, certFile, keyFile) end

--- Routes returns a slice of registered routes, including some useful information, such as:
--- the http method, path and the handler name.
---@return ginRoutesInfo
function ginEngine:Routes() end

---@return httpHandler
function ginEngine:Handler() end

--- SetTrustedProxies set a list of network origins (IPv4 addresses,
--- IPv4 CIDRs, IPv6 addresses or IPv6 CIDRs) from which to trust
--- request's headers that contain alternative client IP when
--- `(*gin.Engine).ForwardedByClientIP` is `true`. `TrustedProxies`
--- feature is enabled by default, and it also trusts all proxies
--- by default. If you want to disable this feature, use
--- Engine.SetTrustedProxies(nil), then Context.ClientIP() will
--- return the remote address directly.
---@param trustedProxies string[]
---@return err
function ginEngine:SetTrustedProxies(trustedProxies) end

--- SetHTMLTemplate associate a template with HTML renderer.
---@param templ templateTemplate
function ginEngine:SetHTMLTemplate(templ) end

--- SetFuncMap sets the FuncMap used for template.FuncMap.
---@param funcMap templateFuncMap
function ginEngine:SetFuncMap(funcMap) end

--- Use attaches a global middleware to the router. i.e. the middleware attached through Use() will be
--- included in the handlers chain for every single request. Even 404, 405, static files...
--- For example, this is the right place for a logger or error management middleware.
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginEngine:Use(...) end

--- HandleContext re-enters a context that has been rewritten.
--- This can be done by setting c.Request.URL.Path to your new target.
--- Disclaimer: You can loop yourself to deal with this, use wisely.
---@param c ginContext
function ginEngine:HandleContext(c) end

--- NoMethod sets the handlers called when Engine.HandleMethodNotAllowed = true.
---@vararg ginHandlerFunc
function ginEngine:NoMethod(...) end

--- Run attaches the router to a http.Server and starts listening and serving HTTP requests.
--- It is a shortcut for http.ListenAndServe(addr, router)
--- Note: this method will block the calling goroutine indefinitely unless an error happens.
---@vararg string
---@return err
function ginEngine:Run(...) end

--- Delims sets template left and right delims and returns an Engine instance.
---@param left string
---@param right string
---@return ginEngine
function ginEngine:Delims(left, right) end

--- IRouter defines all router handle interface includes single and group router.
---@class ginIRouter
local ginIRouter = {}

--- IRoutes defines all router handle interface.
---@class ginIRoutes
local ginIRoutes = {}

--- Param is a single URL parameter, consisting of a key and a value.
---@class ginParam
---@field Key string
---@field Value string
local ginParam = {}

--- H is a shortcut for map[string]any
---@class ginH
local ginH = {}

--- MarshalXML allows type H to be used with xml.Marshal.
---@param e xmlEncoder
---@param start xmlStartElement
---@return err
function ginH:MarshalXML(e, start) end

--- Accounts defines a key/value for user/pass list of authorized logins.
---@class ginAccounts
local ginAccounts = {}

--- HandlerFunc defines the handler used by gin middleware as return value.
---@class ginHandlerFunc
local ginHandlerFunc = {}

--- LogFormatter gives the signature of the formatter function passed to LoggerWithFormatter
---@class ginLogFormatter
local ginLogFormatter = {}

--- RecoveryFunc defines the function passable to CustomRecovery.
---@class ginRecoveryFunc
local ginRecoveryFunc = {}

--- ResponseWriter ...
---@class ginResponseWriter
local ginResponseWriter = {}

--- Negotiate contains all negotiations data.
---@class ginNegotiate
---@field Offered any
---@field HTMLName string
---@field HTMLData any
---@field JSONData any
---@field XMLData any
---@field YAMLData any
---@field Data any
---@field TOMLData any
local ginNegotiate = {}

--- Error represents a error's specification.
---@class ginError
---@field Err err
---@field Type ginErrorType
---@field Meta any
local ginError = {}

--- JSON creates a properly formatted JSON
---@return any
function ginError:JSON() end

--- MarshalJSON implements the json.Marshaller interface.
---@return byte[], err
function ginError:MarshalJSON() end

--- Error implements the error interface.
---@return string
function ginError:Error() end

--- IsType judges one error.
---@param flags ginErrorType
---@return boolean
function ginError:IsType(flags) end

--- Unwrap returns the wrapped error, to allow interoperability with errors.Is(), errors.As() and errors.Unwrap()
---@return err
function ginError:Unwrap() end

--- SetType sets the error's type.
---@param flags ginErrorType
---@return ginError
function ginError:SetType(flags) end

--- SetMeta sets the error's meta data.
---@param data any
---@return ginError
function ginError:SetMeta(data) end

--- LoggerConfig defines the config for Logger middleware.
---@class ginLoggerConfig
---@field Formatter ginLogFormatter
---@field Output any
---@field SkipPaths any
local ginLoggerConfig = {}

--- LogFormatterParams is the structure any formatter will be handed when time to log comes
---@class ginLogFormatterParams
---@field Request any
---@field TimeStamp any
---@field StatusCode number
---@field Latency any
---@field ClientIP string
---@field Method string
---@field Path string
---@field ErrorMessage string
---@field BodySize number
---@field Keys any
local ginLogFormatterParams = {}

--- ResetColor resets all escape attributes.
---@return string
function ginLogFormatterParams:ResetColor() end

--- IsOutputColor indicates whether can colors be outputted to the log.
---@return boolean
function ginLogFormatterParams:IsOutputColor() end

--- StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
---@return string
function ginLogFormatterParams:StatusCodeColor() end

--- MethodColor is the ANSI color for appropriately logging http method to a terminal.
---@return string
function ginLogFormatterParams:MethodColor() end

--- Params is a Param-slice, as returned by the router.
--- The slice is ordered, the first URL parameter is also the first slice value.
--- It is therefore safe to read values by the index.
---@class ginParams
local ginParams = {}

--- Get returns the value of the first Param which key matches the given name and a boolean true.
--- If no matching Param is found, an empty string is returned and a boolean false .
---@param name string
---@return string, boolean
function ginParams:Get(name) end

--- ByName returns the value of the first Param which key matches the given name.
--- If no matching Param is found, an empty string is returned.
---@param name string
---@return string
function ginParams:ByName(name) end

--- ErrorType is an unsigned 64-bit error code as defined in the gin spec.
---@class ginErrorType
local ginErrorType = {}

--- RouteInfo represents a request route's specification which contains method and path and its handler.
---@class ginRouteInfo
---@field Method string
---@field Path string
---@field Handler string
---@field HandlerFunc ginHandlerFunc
local ginRouteInfo = {}

--- RoutesInfo defines a RouteInfo slice.
---@class ginRoutesInfo
local ginRoutesInfo = {}

--- RouterGroup is used internally to configure router, a RouterGroup is associated with
--- a prefix and an array of handlers (middleware).
---@class ginRouterGroup
---@field Handlers ginHandlersChain
local ginRouterGroup = {}

--- PATCH is a shortcut for router.Handle("PATCH", path, handlers).
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:PATCH(relativePath, ...) end

--- GET is a shortcut for router.Handle("GET", path, handlers).
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:GET(relativePath, ...) end

--- PUT is a shortcut for router.Handle("PUT", path, handlers).
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:PUT(relativePath, ...) end

--- HEAD is a shortcut for router.Handle("HEAD", path, handlers).
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:HEAD(relativePath, ...) end

--- Any registers a route that matches all the HTTP methods.
--- GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:Any(relativePath, ...) end

--- Match registers a route that matches the specified methods that you declared.
---@param methods string[]
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:Match(methods, relativePath, ...) end

--- Static serves files from the given file system root.
--- Internally a http.FileServer is used, therefore http.NotFound is used instead
--- of the Router's NotFound handler.
--- To use the operating system's file system implementation,
--- use :
---
---	router.Static("/static", "/var/www")
---@param relativePath string
---@param root string
---@return ginIRoutes
function ginRouterGroup:Static(relativePath, root) end

--- Use adds middleware to the group, see example code in GitHub.
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:Use(...) end

--- Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
--- For example, all the routes that use a common middleware for authorization could be grouped.
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginRouterGroup
function ginRouterGroup:Group(relativePath, ...) end

--- BasePath returns the base path of router group.
--- For example, if v := router.Group("/rest/n/v1/api"), v.BasePath() is "/rest/n/v1/api".
---@return string
function ginRouterGroup:BasePath() end

--- DELETE is a shortcut for router.Handle("DELETE", path, handlers).
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:DELETE(relativePath, ...) end

--- OPTIONS is a shortcut for router.Handle("OPTIONS", path, handlers).
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:OPTIONS(relativePath, ...) end

--- StaticFile registers a single route in order to serve a single file of the local filesystem.
--- router.StaticFile("favicon.ico", "./resources/favicon.ico")
---@param relativePath string
---@param filepath string
---@return ginIRoutes
function ginRouterGroup:StaticFile(relativePath, filepath) end

--- StaticFileFS works just like `StaticFile` but a custom `http.FileSystem` can be used instead..
--- router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
--- Gin by default uses: gin.Dir()
---@param relativePath string
---@param filepath string
---@param fs httpFileSystem
---@return ginIRoutes
function ginRouterGroup:StaticFileFS(relativePath, filepath, fs) end

--- StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
--- Gin by default uses: gin.Dir()
---@param relativePath string
---@param fs httpFileSystem
---@return ginIRoutes
function ginRouterGroup:StaticFS(relativePath, fs) end

--- Handle registers a new request handle and middleware with the given path and method.
--- The last handler should be the real handler, the other ones should be middleware that can and should be shared among different routes.
--- See the example code in GitHub.
---
--- For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
--- functions can be used.
---
--- This function is intended for bulk loading and to allow the usage of less
--- frequently used, non-standardized or custom methods (e.g. for internal
--- communication with a proxy).
---@param httpMethod string
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:Handle(httpMethod, relativePath, ...) end

--- POST is a shortcut for router.Handle("POST", path, handlers).
---@param relativePath string
---@vararg ginHandlerFunc
---@return ginIRoutes
function ginRouterGroup:POST(relativePath, ...) end
