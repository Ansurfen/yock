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

--- Handle registers the handler for the given pattern
--- in the DefaultServeMux.
--- The documentation for ServeMux explains how patterns are matched.
---@param pattern string
---@param handler httpHandler
function http.Handle(pattern, handler) end

--- NotFoundHandler returns a simple request handler
--- that replies to each request with a “404 page not found” reply.
---@return httpHandler
function http.NotFoundHandler() end

--- NotFound replies to the request with an HTTP 404 not found error.
---@param w httpResponseWriter
---@param r httpRequest
function http.NotFound(w, r) end

--- NewServeMux allocates and returns a new ServeMux.
---@return httpServeMux
function http.NewServeMux() end

--- MaxBytesReader is similar to io.LimitReader but is intended for
--- limiting the size of incoming request bodies. In contrast to
--- io.LimitReader, MaxBytesReader's result is a ReadCloser, returns a
--- non-nil error of type *MaxBytesError for a Read beyond the limit,
--- and closes the underlying reader when its Close method is called.
---
--- MaxBytesReader prevents clients from accidentally or maliciously
--- sending a large request and wasting server resources. If possible,
--- it tells the ResponseWriter to close the connection after the limit
--- has been reached.
---@param w httpResponseWriter
---@param r ioReadCloser
---@param n number
---@return ioReadCloser
function http.MaxBytesReader(w, r, n) end

--- HandleFunc registers the handler function for the given pattern
--- in the DefaultServeMux.
--- The documentation for ServeMux explains how patterns are matched.
---@param pattern string
---@param handler function
function http.HandleFunc(pattern, handler) end

--- ServeTLS accepts incoming HTTPS connections on the listener l,
--- creating a new service goroutine for each. The service goroutines
--- read requests and then call handler to reply to them.
---
--- The handler is typically nil, in which case the DefaultServeMux is used.
---
--- Additionally, files containing a certificate and matching private key
--- for the server must be provided. If the certificate is signed by a
--- certificate authority, the certFile should be the concatenation
--- of the server's certificate, any intermediates, and the CA's certificate.
---
--- ServeTLS always returns a non-nil error.
---@param l netListener
---@param handler httpHandler
---@param certFile string
---@param keyFile string
---@return err
function http.ServeTLS(l, handler, certFile, keyFile) end

--- TimeoutHandler returns a Handler that runs h with the given time limit.
---
--- The new Handler calls h.ServeHTTP to handle each request, but if a
--- call runs for longer than its time limit, the handler responds with
--- a 503 Service Unavailable error and the given message in its body.
--- (If msg is empty, a suitable default message will be sent.)
--- After such a timeout, writes by h to its ResponseWriter will return
--- ErrHandlerTimeout.
---
--- TimeoutHandler supports the Pusher interface but does not support
--- the Hijacker or Flusher interfaces.
---@param h httpHandler
---@param dt timeDuration
---@param msg string
---@return httpHandler
function http.TimeoutHandler(h, dt, msg) end

--- ProxyURL returns a proxy function (for use in a Transport)
--- that always returns the same URL.
---@param fixedURL urlURL
---@return any
function http.ProxyURL(fixedURL) end


---@param w httpResponseWriter
---@param req httpRequest
function http.FlagServer(w, req) end

--- FS converts fsys to a FileSystem implementation,
--- for use with FileServer and NewFileTransport.
--- The files provided by fsys must implement io.Seeker.
---@param fsys fsFS
---@return httpFileSystem
function http.FS(fsys) end

--- NewRequestWithContext returns a new Request given a method, URL, and
--- optional body.
---
--- If the provided body is also an io.Closer, the returned
--- Request.Body is set to body and will be closed by the Client
--- methods Do, Post, and PostForm, and Transport.RoundTrip.
---
--- NewRequestWithContext returns a Request suitable for use with
--- Client.Do or Transport.RoundTrip. To create a request for use with
--- testing a Server Handler, either use the NewRequest function in the
--- net/http/httptest package, use ReadRequest, or manually update the
--- Request fields. For an outgoing client request, the context
--- controls the entire lifetime of a request and its response:
--- obtaining a connection, sending the request, and reading the
--- response headers and body. See the Request type's documentation for
--- the difference between inbound and outbound request fields.
---
--- If body is of type *bytes.Buffer, *bytes.Reader, or
--- *strings.Reader, the returned request's ContentLength is set to its
--- exact value (instead of -1), GetBody is populated (so 307 and 308
--- redirects can replay the body), and Body is set to NoBody if the
--- ContentLength is 0.
---@param ctx contextContext
---@param method string
---@param url string
---@param body ioReader
---@return httpRequest, err
function http.NewRequestWithContext(ctx, method, url, body) end

--- simple argument server
---@param w httpResponseWriter
---@param req httpRequest
function http.ArgServer(w, req) end

--- Head issues a HEAD to the specified URL. If the response is one of
--- the following redirect codes, Head follows the redirect, up to a
--- maximum of 10 redirects:
---
---	301 (Moved Permanently)
---	302 (Found)
---	303 (See Other)
---	307 (Temporary Redirect)
---	308 (Permanent Redirect)
---
--- Head is a wrapper around DefaultClient.Head.
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and DefaultClient.Do.
---@param url string
---@return httpResponse, err
function http.Head(url) end


---@param w httpResponseWriter
---@param req httpRequest
function http.HelloServer(w, req) end

--- ServeFile replies to the request with the contents of the named
--- file or directory.
---
--- If the provided file or directory name is a relative path, it is
--- interpreted relative to the current directory and may ascend to
--- parent directories. If the provided name is constructed from user
--- input, it should be sanitized before calling ServeFile.
---
--- As a precaution, ServeFile will reject requests where r.URL.Path
--- contains a ".." path element; this protects against callers who
--- might unsafely use filepath.Join on r.URL.Path without sanitizing
--- it and then use that filepath.Join result as the name argument.
---
--- As another special case, ServeFile redirects any request where r.URL.Path
--- ends in "/index.html" to the same path, without the final
--- "index.html". To avoid such redirects either modify the path or
--- use ServeContent.
---
--- Outside of those two special cases, ServeFile does not use
--- r.URL.Path for selecting the file or directory to serve; only the
--- file or directory provided in the name argument is used.
---@param w httpResponseWriter
---@param r httpRequest
---@param name string
function http.ServeFile(w, r, name) end

--- ParseHTTPVersion parses an HTTP version string according to RFC 7230, section 2.6.
--- "HTTP/1.0" returns (1, 0, true). Note that strings without
--- a minor version, such as "HTTP/2", are not valid.
---@param vers string
---@return number, boolean
function http.ParseHTTPVersion(vers) end

--- Redirect replies to the request with a redirect to url,
--- which may be a path relative to the request path.
---
--- The provided code should be in the 3xx range and is usually
--- StatusMovedPermanently, StatusFound or StatusSeeOther.
---
--- If the Content-Type header has not been set, Redirect sets it
--- to "text/html; charset=utf-8" and writes a small HTML body.
--- Setting the Content-Type header to any value, including nil,
--- disables that behavior.
---@param w httpResponseWriter
---@param r httpRequest
---@param url string
---@param code number
function http.Redirect(w, r, url, code) end

--- Serve accepts incoming HTTP connections on the listener l,
--- creating a new service goroutine for each. The service goroutines
--- read requests and then call handler to reply to them.
---
--- The handler is typically nil, in which case the DefaultServeMux is used.
---
--- HTTP/2 support is only enabled if the Listener returns *tls.Conn
--- connections and they were configured with "h2" in the TLS
--- Config.NextProtos.
---
--- Serve always returns a non-nil error.
---@param l netListener
---@param handler httpHandler
---@return err
function http.Serve(l, handler) end

--- ListenAndServeTLS acts identically to ListenAndServe, except that it
--- expects HTTPS connections. Additionally, files containing a certificate and
--- matching private key for the server must be provided. If the certificate
--- is signed by a certificate authority, the certFile should be the concatenation
--- of the server's certificate, any intermediates, and the CA's certificate.
---@param addr string
---@param certFile string
---@param keyFile string
---@param handler httpHandler
---@return err
function http.ListenAndServeTLS(addr, certFile, keyFile, handler) end

--- Error replies to the request with the specified error message and HTTP code.
--- It does not otherwise end the request; the caller should ensure no further
--- writes are done to w.
--- The error message should be plain text.
---@param w httpResponseWriter
---@param error string
---@param code number
function http.Error(w, error, code) end

--- Get issues a GET to the specified URL. If the response is one of
--- the following redirect codes, Get follows the redirect, up to a
--- maximum of 10 redirects:
---
---	301 (Moved Permanently)
---	302 (Found)
---	303 (See Other)
---	307 (Temporary Redirect)
---	308 (Permanent Redirect)
---
--- An error is returned if there were too many redirects or if there
--- was an HTTP protocol error. A non-2xx response doesn't cause an
--- error. Any returned error will be of type *url.Error. The url.Error
--- value's Timeout method will report true if the request timed out.
---
--- When err is nil, resp always contains a non-nil resp.Body.
--- Caller should close resp.Body when done reading from it.
---
--- Get is a wrapper around DefaultClient.Get.
---
--- To make a request with custom headers, use NewRequest and
--- DefaultClient.Do.
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and DefaultClient.Do.
---@param url string
---@return httpResponse, err
function http.Get(url) end

--- FileServer returns a handler that serves HTTP requests
--- with the contents of the file system rooted at root.
---
--- As a special case, the returned file server redirects any request
--- ending in "/index.html" to the same path, without the final
--- "index.html".
---
--- To use the operating system's file system implementation,
--- use http.Dir:
---
---	http.Handle("/", http.FileServer(http.Dir("/tmp")))
---
--- To use an fs.FS implementation, use http.FS to convert it:
---
---	http.Handle("/", http.FileServer(http.FS(fsys)))
---@param root httpFileSystem
---@return httpHandler
function http.FileServer(root) end


---@return httpChan
function http.ChanCreate() end


---@param w httpResponseWriter
---@param req httpRequest
function http.Logger(w, req) end

--- MaxBytesHandler returns a Handler that runs h with its ResponseWriter and Request.Body wrapped by a MaxBytesReader.
---@param h httpHandler
---@param n number
---@return httpHandler
function http.MaxBytesHandler(h, n) end

--- StatusText returns a text for the HTTP status code. It returns the empty
--- string if the code is unknown.
---@param code number
---@return string
function http.StatusText(code) end

--- StripPrefix returns a handler that serves HTTP requests by removing the
--- given prefix from the request URL's Path (and RawPath if set) and invoking
--- the handler h. StripPrefix handles a request for a path that doesn't begin
--- with prefix by replying with an HTTP 404 not found error. The prefix must
--- match exactly: if the prefix in the request contains escaped characters
--- the reply is also an HTTP 404 not found error.
---@param prefix string
---@param h httpHandler
---@return httpHandler
function http.StripPrefix(prefix, h) end

--- ProxyFromEnvironment returns the URL of the proxy to use for a
--- given request, as indicated by the environment variables
--- HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
--- thereof). Requests use the proxy from the environment variable
--- matching their scheme, unless excluded by NO_PROXY.
---
--- The environment values may be either a complete URL or a
--- "host[:port]", in which case the "http" scheme is assumed.
--- The schemes "http", "https", and "socks5" are supported.
--- An error is returned if the value is a different form.
---
--- A nil URL and nil error are returned if no proxy is defined in the
--- environment, or a proxy should not be used for the given request,
--- as defined by NO_PROXY.
---
--- As a special case, if req.URL.Host is "localhost" (with or without
--- a port number), then a nil URL and nil error will be returned.
---@param req httpRequest
---@return urlURL, err
function http.ProxyFromEnvironment(req) end

--- exec a program, redirecting output.
---@param rw httpResponseWriter
---@param req httpRequest
function http.DateServer(rw, req) end

--- ReadRequest reads and parses an incoming request from b.
---
--- ReadRequest is a low-level function and should only be used for
--- specialized applications; most code should use the Server to read
--- requests and handle them via the Handler interface. ReadRequest
--- only supports HTTP/1.x requests. For HTTP/2, use golang.org/x/net/http2.
---@param b bufioReader
---@return httpRequest, err
function http.ReadRequest(b) end

--- ReadResponse reads and returns an HTTP response from r.
--- The req parameter optionally specifies the Request that corresponds
--- to this Response. If nil, a GET request is assumed.
--- Clients must call resp.Body.Close when finished reading resp.Body.
--- After that call, clients can inspect resp.Trailer to find key/value
--- pairs included in the response trailer.
---@param r bufioReader
---@param req httpRequest
---@return httpResponse, err
function http.ReadResponse(r, req) end

--- PostForm issues a POST to the specified URL, with data's keys and
--- values URL-encoded as the request body.
---
--- The Content-Type header is set to application/x-www-form-urlencoded.
--- To set other headers, use NewRequest and DefaultClient.Do.
---
--- When err is nil, resp always contains a non-nil resp.Body.
--- Caller should close resp.Body when done reading from it.
---
--- PostForm is a wrapper around DefaultClient.PostForm.
---
--- See the Client.Do method documentation for details on how redirects
--- are handled.
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and DefaultClient.Do.
---@param url string
---@param data urlValues
---@return httpResponse, err
function http.PostForm(url, data) end

--- ParseTime parses a time header (such as the Date: header),
--- trying each of the three formats allowed by HTTP/1.1:
--- TimeFormat, time.RFC850, and time.ANSIC.
---@param text string
---@return timeTime, err
function http.ParseTime(text) end

--- AllowQuerySemicolons returns a handler that serves requests by converting any
--- unescaped semicolons in the URL query to ampersands, and invoking the handler h.
---
--- This restores the pre-Go 1.17 behavior of splitting query parameters on both
--- semicolons and ampersands. (See golang.org/issue/25192). Note that this
--- behavior doesn't match that of many proxies, and the mismatch can lead to
--- security issues.
---
--- AllowQuerySemicolons should be invoked before Request.ParseForm is called.
---@param h httpHandler
---@return httpHandler
function http.AllowQuerySemicolons(h) end

--- CanonicalHeaderKey returns the canonical format of the
--- header key s. The canonicalization converts the first
--- letter and any letter following a hyphen to upper case;
--- the rest are converted to lowercase. For example, the
--- canonical key for "accept-encoding" is "Accept-Encoding".
--- If s contains a space or invalid header field bytes, it is
--- returned without modifications.
---@param s string
---@return string
function http.CanonicalHeaderKey(s) end

--- NewResponseController creates a ResponseController for a request.
---
--- The ResponseWriter should be the original value passed to the Handler.ServeHTTP method,
--- or have an Unwrap method returning the original ResponseWriter.
---
--- If the ResponseWriter implements any of the following methods, the ResponseController
--- will call them as appropriate:
---
---	Flush()
---	FlushError() error // alternative Flush returning an error
---	Hijack() (net.Conn, *bufio.ReadWriter, error)
---	SetReadDeadline(deadline time.Time) error
---	SetWriteDeadline(deadline time.Time) error
---
--- If the ResponseWriter does not support a method, ResponseController returns
--- an error matching ErrNotSupported.
---@param rw httpResponseWriter
---@return httpResponseController
function http.NewResponseController(rw) end

--- NewFileTransport returns a new RoundTripper, serving the provided
--- FileSystem. The returned RoundTripper ignores the URL host in its
--- incoming requests, as well as most other properties of the
--- request.
---
--- The typical use case for NewFileTransport is to register the "file"
--- protocol with a Transport, as in:
---
---	t := &http.Transport{}
---	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
---	c := &http.Client{Transport: t}
---	res, err := c.Get("file:///etc/passwd")
---	...
---@param fs httpFileSystem
---@return httpRoundTripper
function http.NewFileTransport(fs) end

--- ServeContent replies to the request using the content in the
--- provided ReadSeeker. The main benefit of ServeContent over io.Copy
--- is that it handles Range requests properly, sets the MIME type, and
--- handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since,
--- and If-Range requests.
---
--- If the response's Content-Type header is not set, ServeContent
--- first tries to deduce the type from name's file extension and,
--- if that fails, falls back to reading the first block of the content
--- and passing it to DetectContentType.
--- The name is otherwise unused; in particular it can be empty and is
--- never sent in the response.
---
--- If modtime is not the zero time or Unix epoch, ServeContent
--- includes it in a Last-Modified header in the response. If the
--- request includes an If-Modified-Since header, ServeContent uses
--- modtime to decide whether the content needs to be sent at all.
---
--- The content's Seek method must work: ServeContent uses
--- a seek to the end of the content to determine its size.
---
--- If the caller has set w's ETag header formatted per RFC 7232, section 2.3,
--- ServeContent uses it to handle requests using If-Match, If-None-Match, or If-Range.
---
--- Note that *os.File implements the io.ReadSeeker interface.
---@param w httpResponseWriter
---@param req httpRequest
---@param name string
---@param modtime timeTime
---@param content ioReadSeeker
function http.ServeContent(w, req, name, modtime, content) end

--- NewRequest wraps NewRequestWithContext using context.Background.
---@param method string
---@param url string
---@param body ioReader
---@return httpRequest, err
function http.NewRequest(method, url, body) end

--- RedirectHandler returns a request handler that redirects
--- each request it receives to the given url using the given
--- status code.
---
--- The provided code should be in the 3xx range and is usually
--- StatusMovedPermanently, StatusFound or StatusSeeOther.
---@param url string
---@param code number
---@return httpHandler
function http.RedirectHandler(url, code) end

--- ListenAndServe listens on the TCP network address addr and then calls
--- Serve with handler to handle requests on incoming connections.
--- Accepted connections are configured to enable TCP keep-alives.
---
--- The handler is typically nil, in which case the DefaultServeMux is used.
---
--- ListenAndServe always returns a non-nil error.
---@param addr string
---@param handler httpHandler
---@return err
function http.ListenAndServe(addr, handler) end

--- DetectContentType implements the algorithm described
--- at https://mimesniff.spec.whatwg.org/ to determine the
--- Content-Type of the given data. It considers at most the
--- first 512 bytes of data. DetectContentType always returns
--- a valid MIME type: if it cannot determine a more specific one, it
--- returns "application/octet-stream".
---@param data byte[]
---@return string
function http.DetectContentType(data) end

--- Post issues a POST to the specified URL.
---
--- Caller should close resp.Body when done reading from it.
---
--- If the provided body is an io.Closer, it is closed after the
--- request.
---
--- Post is a wrapper around DefaultClient.Post.
---
--- To set custom headers, use NewRequest and DefaultClient.Do.
---
--- See the Client.Do method documentation for details on how redirects
--- are handled.
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and DefaultClient.Do.
---@param url string
---@param contentType string
---@param body ioReader
---@return httpResponse, err
function http.Post(url, contentType, body) end

--- SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
--- The provided cookie must have a valid Name. Invalid cookies may be
--- silently dropped.
---@param w httpResponseWriter
---@param cookie httpCookie
function http.SetCookie(w, cookie) end

--- A Client is an HTTP client. Its zero value (DefaultClient) is a
--- usable client that uses DefaultTransport.
---
--- The Client's Transport typically has internal state (cached TCP
--- connections), so Clients should be reused instead of created as
--- needed. Clients are safe for concurrent use by multiple goroutines.
---
--- A Client is higher-level than a RoundTripper (such as Transport)
--- and additionally handles HTTP details such as cookies and
--- redirects.
---
--- When following redirects, the Client will forward all headers set on the
--- initial Request except:
---
--- • when forwarding sensitive headers like "Authorization",
--- "WWW-Authenticate", and "Cookie" to untrusted targets.
--- These headers will be ignored when following a redirect to a domain
--- that is not a subdomain match or exact match of the initial domain.
--- For example, a redirect from "foo.com" to either "foo.com" or "sub.foo.com"
--- will forward the sensitive headers, but a redirect to "bar.com" will not.
---
--- • when forwarding the "Cookie" header with a non-nil cookie Jar.
--- Since each redirect may mutate the state of the cookie jar,
--- a redirect may possibly alter a cookie set in the initial request.
--- When forwarding the "Cookie" header, any mutated cookies will be omitted,
--- with the expectation that the Jar will insert those mutated cookies
--- with the updated values (assuming the origin matches).
--- If Jar is nil, the initial cookies are forwarded without change.
---@class httpClient
---@field Transport httpRoundTripper
---@field CheckRedirect any
---@field Jar httpCookieJar
---@field Timeout any
local httpClient = {}

--- Head issues a HEAD to the specified URL. If the response is one of the
--- following redirect codes, Head follows the redirect after calling the
--- Client's CheckRedirect function:
---
---	301 (Moved Permanently)
---	302 (Found)
---	303 (See Other)
---	307 (Temporary Redirect)
---	308 (Permanent Redirect)
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and Client.Do.
---@param url string
---@return httpResponse, err
function httpClient:Head(url) end

--- Get issues a GET to the specified URL. If the response is one of the
--- following redirect codes, Get follows the redirect after calling the
--- Client's CheckRedirect function:
---
---	301 (Moved Permanently)
---	302 (Found)
---	303 (See Other)
---	307 (Temporary Redirect)
---	308 (Permanent Redirect)
---
--- An error is returned if the Client's CheckRedirect function fails
--- or if there was an HTTP protocol error. A non-2xx response doesn't
--- cause an error. Any returned error will be of type *url.Error. The
--- url.Error value's Timeout method will report true if the request
--- timed out.
---
--- When err is nil, resp always contains a non-nil resp.Body.
--- Caller should close resp.Body when done reading from it.
---
--- To make a request with custom headers, use NewRequest and Client.Do.
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and Client.Do.
---@param url string
---@return httpResponse, err
function httpClient:Get(url) end

--- PostForm issues a POST to the specified URL,
--- with data's keys and values URL-encoded as the request body.
---
--- The Content-Type header is set to application/x-www-form-urlencoded.
--- To set other headers, use NewRequest and Client.Do.
---
--- When err is nil, resp always contains a non-nil resp.Body.
--- Caller should close resp.Body when done reading from it.
---
--- See the Client.Do method documentation for details on how redirects
--- are handled.
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and Client.Do.
---@param url string
---@param data urlValues
---@return httpResponse, err
function httpClient:PostForm(url, data) end

--- CloseIdleConnections closes any connections on its Transport which
--- were previously connected from previous requests but are now
--- sitting idle in a "keep-alive" state. It does not interrupt any
--- connections currently in use.
---
--- If the Client's Transport does not have a CloseIdleConnections method
--- then this method does nothing.
function httpClient:CloseIdleConnections() end

--- Do sends an HTTP request and returns an HTTP response, following
--- policy (such as redirects, cookies, auth) as configured on the
--- client.
---
--- An error is returned if caused by client policy (such as
--- CheckRedirect), or failure to speak HTTP (such as a network
--- connectivity problem). A non-2xx status code doesn't cause an
--- error.
---
--- If the returned error is nil, the Response will contain a non-nil
--- Body which the user is expected to close. If the Body is not both
--- read to EOF and closed, the Client's underlying RoundTripper
--- (typically Transport) may not be able to re-use a persistent TCP
--- connection to the server for a subsequent "keep-alive" request.
---
--- The request Body, if non-nil, will be closed by the underlying
--- Transport, even on errors.
---
--- On error, any Response can be ignored. A non-nil Response with a
--- non-nil error only occurs when CheckRedirect fails, and even then
--- the returned Response.Body is already closed.
---
--- Generally Get, Post, or PostForm will be used instead of Do.
---
--- If the server replies with a redirect, the Client first uses the
--- CheckRedirect function to determine whether the redirect should be
--- followed. If permitted, a 301, 302, or 303 redirect causes
--- subsequent requests to use HTTP method GET
--- (or HEAD if the original request was HEAD), with no body.
--- A 307 or 308 redirect preserves the original HTTP method and body,
--- provided that the Request.GetBody function is defined.
--- The NewRequest function automatically sets GetBody for common
--- standard library body types.
---
--- Any returned error will be of type *url.Error. The url.Error
--- value's Timeout method will report true if the request timed out.
---@param req httpRequest
---@return httpResponse, err
function httpClient:Do(req) end

--- Post issues a POST to the specified URL.
---
--- Caller should close resp.Body when done reading from it.
---
--- If the provided body is an io.Closer, it is closed after the
--- request.
---
--- To set custom headers, use NewRequest and Client.Do.
---
--- To make a request with a specified context.Context, use NewRequestWithContext
--- and Client.Do.
---
--- See the Client.Do method documentation for details on how redirects
--- are handled.
---@param url string
---@param contentType string
---@param body ioReader
---@return httpResponse, err
function httpClient:Post(url, contentType, body) end

--- A Request represents an HTTP request received by a server
--- or to be sent by a client.
---
--- The field semantics differ slightly between client and server
--- usage. In addition to the notes on the fields below, see the
--- documentation for Request.Write and RoundTripper.
---@class httpRequest
---@field Method string
---@field URL any
---@field Proto string
---@field ProtoMajor number
---@field ProtoMinor number
---@field Header gzipHeader
---@field Body any
---@field GetBody any
---@field ContentLength number
---@field TransferEncoding any
---@field Close boolean
---@field Host string
---@field Form any
---@field PostForm any
---@field MultipartForm any
---@field Trailer gzipHeader
---@field RemoteAddr string
---@field RequestURI string
---@field TLS any
---@field Cancel any
---@field Response httpResponse
local httpRequest = {}

--- ParseForm populates r.Form and r.PostForm.
---
--- For all requests, ParseForm parses the raw query from the URL and updates
--- r.Form.
---
--- For POST, PUT, and PATCH requests, it also reads the request body, parses it
--- as a form and puts the results into both r.PostForm and r.Form. Request body
--- parameters take precedence over URL query string values in r.Form.
---
--- If the request Body's size has not already been limited by MaxBytesReader,
--- the size is capped at 10MB.
---
--- For other HTTP methods, or when the Content-Type is not
--- application/x-www-form-urlencoded, the request Body is not read, and
--- r.PostForm is initialized to a non-nil, empty value.
---
--- ParseMultipartForm calls ParseForm automatically.
--- ParseForm is idempotent.
---@return err
function httpRequest:ParseForm() end

--- SetBasicAuth sets the request's Authorization header to use HTTP
--- Basic Authentication with the provided username and password.
---
--- With HTTP Basic Authentication the provided username and password
--- are not encrypted. It should generally only be used in an HTTPS
--- request.
---
--- The username may not contain a colon. Some protocols may impose
--- additional requirements on pre-escaping the username and
--- password. For instance, when used with OAuth2, both arguments must
--- be URL encoded first with url.QueryEscape.
---@param username string
---@param password string
function httpRequest:SetBasicAuth(username, password) end

--- Write writes an HTTP/1.1 request, which is the header and body, in wire format.
--- This method consults the following fields of the request:
---
---	Host
---	URL
---	Method (defaults to "GET")
---	Header
---	ContentLength
---	TransferEncoding
---	Body
---
--- If Body is present, Content-Length is <= 0 and TransferEncoding
--- hasn't been set to "identity", Write adds "Transfer-Encoding:
--- chunked" to the header. Body is closed after it is sent.
---@param w ioWriter
---@return err
function httpRequest:Write(w) end

--- Context returns the request's context. To change the context, use
--- Clone or WithContext.
---
--- The returned context is always non-nil; it defaults to the
--- background context.
---
--- For outgoing client requests, the context controls cancellation.
---
--- For incoming server requests, the context is canceled when the
--- client's connection closes, the request is canceled (with HTTP/2),
--- or when the ServeHTTP method returns.
---@return contextContext
function httpRequest:Context() end

--- MultipartReader returns a MIME multipart reader if this is a
--- multipart/form-data or a multipart/mixed POST request, else returns nil and an error.
--- Use this function instead of ParseMultipartForm to
--- process the request body as a stream.
---@return any, err
function httpRequest:MultipartReader() end

--- Clone returns a deep copy of r with its context changed to ctx.
--- The provided ctx must be non-nil.
---
--- For an outgoing client request, the context controls the entire
--- lifetime of a request and its response: obtaining a connection,
--- sending the request, and reading the response headers and body.
---@param ctx contextContext
---@return httpRequest
function httpRequest:Clone(ctx) end

--- ProtoAtLeast reports whether the HTTP protocol used
--- in the request is at least major.minor.
---@param major number
---@param minor number
---@return boolean
function httpRequest:ProtoAtLeast(major, minor) end

--- Cookies parses and returns the HTTP cookies sent with the request.
---@return any
function httpRequest:Cookies() end

--- FormValue returns the first value for the named component of the query.
--- POST and PUT body parameters take precedence over URL query string values.
--- FormValue calls ParseMultipartForm and ParseForm if necessary and ignores
--- any errors returned by these functions.
--- If key is not present, FormValue returns the empty string.
--- To access multiple values of the same key, call ParseForm and
--- then inspect Request.Form directly.
---@param key string
---@return string
function httpRequest:FormValue(key) end

--- WithContext returns a shallow copy of r with its context changed
--- to ctx. The provided ctx must be non-nil.
---
--- For outgoing client request, the context controls the entire
--- lifetime of a request and its response: obtaining a connection,
--- sending the request, and reading the response headers and body.
---
--- To create a new request with a context, use NewRequestWithContext.
--- To make a deep copy of a request with a new context, use Request.Clone.
---@param ctx contextContext
---@return httpRequest
function httpRequest:WithContext(ctx) end

--- ParseMultipartForm parses a request body as multipart/form-data.
--- The whole request body is parsed and up to a total of maxMemory bytes of
--- its file parts are stored in memory, with the remainder stored on
--- disk in temporary files.
--- ParseMultipartForm calls ParseForm if necessary.
--- If ParseForm returns an error, ParseMultipartForm returns it but also
--- continues parsing the request body.
--- After one call to ParseMultipartForm, subsequent calls have no effect.
---@param maxMemory number
---@return err
function httpRequest:ParseMultipartForm(maxMemory) end

--- PostFormValue returns the first value for the named component of the POST,
--- PATCH, or PUT request body. URL query parameters are ignored.
--- PostFormValue calls ParseMultipartForm and ParseForm if necessary and ignores
--- any errors returned by these functions.
--- If key is not present, PostFormValue returns the empty string.
---@param key string
---@return string
function httpRequest:PostFormValue(key) end

--- UserAgent returns the client's User-Agent, if sent in the request.
---@return string
function httpRequest:UserAgent() end

--- BasicAuth returns the username and password provided in the request's
--- Authorization header, if the request uses HTTP Basic Authentication.
--- See RFC 2617, Section 2.
---@return string, boolean
function httpRequest:BasicAuth() end

--- AddCookie adds a cookie to the request. Per RFC 6265 section 5.4,
--- AddCookie does not attach more than one Cookie header field. That
--- means all cookies, if any, are written into the same line,
--- separated by semicolon.
--- AddCookie only sanitizes c's name and value, and does not sanitize
--- a Cookie header already present in the request.
---@param c httpCookie
function httpRequest:AddCookie(c) end

--- Referer returns the referring URL, if sent in the request.
---
--- Referer is misspelled as in the request itself, a mistake from the
--- earliest days of HTTP.  This value can also be fetched from the
--- Header map as Header["Referer"]; the benefit of making it available
--- as a method is that the compiler can diagnose programs that use the
--- alternate (correct English) spelling req.Referrer() but cannot
--- diagnose programs that use Header["Referrer"].
---@return string
function httpRequest:Referer() end

--- WriteProxy is like Write but writes the request in the form
--- expected by an HTTP proxy. In particular, WriteProxy writes the
--- initial Request-URI line of the request with an absolute URI, per
--- section 5.3 of RFC 7230, including the scheme and host.
--- In either case, WriteProxy also writes a Host header, using
--- either r.Host or r.URL.Host.
---@param w ioWriter
---@return err
function httpRequest:WriteProxy(w) end

--- Cookie returns the named cookie provided in the request or
--- ErrNoCookie if not found.
--- If multiple cookies match the given name, only one cookie will
--- be returned.
---@param name string
---@return httpCookie, err
function httpRequest:Cookie(name) end

--- FormFile returns the first file for the provided form key.
--- FormFile calls ParseMultipartForm and ParseForm if necessary.
---@param key string
---@return any, any, err
function httpRequest:FormFile(key) end

--- A ResponseController is used by an HTTP handler to control the response.
---
--- A ResponseController may not be used after the Handler.ServeHTTP method has returned.
---@class httpResponseController
local httpResponseController = {}

--- Hijack lets the caller take over the connection.
--- See the Hijacker interface for details.
---@return any, any, err
function httpResponseController:Hijack() end

--- SetReadDeadline sets the deadline for reading the entire request, including the body.
--- Reads from the request body after the deadline has been exceeded will return an error.
--- A zero value means no deadline.
---
--- Setting the read deadline after it has been exceeded will not extend it.
---@param deadline timeTime
---@return err
function httpResponseController:SetReadDeadline(deadline) end

--- SetWriteDeadline sets the deadline for writing the response.
--- Writes to the response body after the deadline has been exceeded will not block,
--- but may succeed if the data has been buffered.
--- A zero value means no deadline.
---
--- Setting the write deadline after it has been exceeded will not extend it.
---@param deadline timeTime
---@return err
function httpResponseController:SetWriteDeadline(deadline) end

--- Flush flushes buffered data to the client.
---@return err
function httpResponseController:Flush() end

--- ServeMux is an HTTP request multiplexer.
--- It matches the URL of each incoming request against a list of registered
--- patterns and calls the handler for the pattern that
--- most closely matches the URL.
---
--- Patterns name fixed, rooted paths, like "/favicon.ico",
--- or rooted subtrees, like "/images/" (note the trailing slash).
--- Longer patterns take precedence over shorter ones, so that
--- if there are handlers registered for both "/images/"
--- and "/images/thumbnails/", the latter handler will be
--- called for paths beginning "/images/thumbnails/" and the
--- former will receive requests for any other paths in the
--- "/images/" subtree.
---
--- Note that since a pattern ending in a slash names a rooted subtree,
--- the pattern "/" matches all paths not matched by other registered
--- patterns, not just the URL with Path == "/".
---
--- If a subtree has been registered and a request is received naming the
--- subtree root without its trailing slash, ServeMux redirects that
--- request to the subtree root (adding the trailing slash). This behavior can
--- be overridden with a separate registration for the path without
--- the trailing slash. For example, registering "/images/" causes ServeMux
--- to redirect a request for "/images" to "/images/", unless "/images" has
--- been registered separately.
---
--- Patterns may optionally begin with a host name, restricting matches to
--- URLs on that host only. Host-specific patterns take precedence over
--- general patterns, so that a handler might register for the two patterns
--- "/codesearch" and "codesearch.google.com/" without also taking over
--- requests for "http://www.google.com/".
---
--- ServeMux also takes care of sanitizing the URL request path and the Host
--- header, stripping the port number and redirecting any request containing . or
--- .. elements or repeated slashes to an equivalent, cleaner URL.
---@class httpServeMux
local httpServeMux = {}

--- Handler returns the handler to use for the given request,
--- consulting r.Method, r.Host, and r.URL.Path. It always returns
--- a non-nil handler. If the path is not in its canonical form, the
--- handler will be an internally-generated handler that redirects
--- to the canonical path. If the host contains a port, it is ignored
--- when matching handlers.
---
--- The path and host are used unchanged for CONNECT requests.
---
--- Handler also returns the registered pattern that matches the
--- request or, in the case of internally-generated redirects,
--- the pattern that will match after following the redirect.
---
--- If there is no registered handler that applies to the request,
--- Handler returns a “page not found” handler and an empty pattern.
---@param r httpRequest
---@return httpHandler, string
function httpServeMux:Handler(r) end

--- ServeHTTP dispatches the request to the handler whose
--- pattern most closely matches the request URL.
---@param w httpResponseWriter
---@param r httpRequest
function httpServeMux:ServeHTTP(w, r) end

--- Handle registers the handler for the given pattern.
--- If a handler already exists for pattern, Handle panics.
---@param pattern string
---@param handler httpHandler
function httpServeMux:Handle(pattern, handler) end

--- HandleFunc registers the handler function for the given pattern.
---@param pattern string
---@param handler any
function httpServeMux:HandleFunc(pattern, handler) end

--- RoundTripper is an interface representing the ability to execute a
--- single HTTP transaction, obtaining the Response for a given Request.
---
--- A RoundTripper must be safe for concurrent use by multiple
--- goroutines.
---@class httpRoundTripper
local httpRoundTripper = {}

--- A Dir implements FileSystem using the native file system restricted to a
--- specific directory tree.
---
--- While the FileSystem.Open method takes '/'-separated paths, a Dir's string
--- value is a filename on the native file system, not a URL, so it is separated
--- by filepath.Separator, which isn't necessarily '/'.
---
--- Note that Dir could expose sensitive files and directories. Dir will follow
--- symlinks pointing out of the directory tree, which can be especially dangerous
--- if serving from a directory in which users are able to create arbitrary symlinks.
--- Dir will also allow access to files and directories starting with a period,
--- which could expose sensitive directories like .git or sensitive files like
--- .htpasswd. To exclude files with a leading period, remove the files/directories
--- from the server or create a custom FileSystem implementation.
---
--- An empty Dir is treated as ".".
---@class httpDir
local httpDir = {}

--- Open implements FileSystem using os.Open, opening files for reading rooted
--- and relative to the directory d.
---@param name string
---@return osFile, err
function httpDir:Open(name) end

--- A File is returned by a FileSystem's Open method and can be
--- served by the FileServer implementation.
---
--- The methods should behave the same as those on an *os.File.
---@class osFile
local osFile = {}

--- The CloseNotifier interface is implemented by ResponseWriters which
--- allow detecting when the underlying connection has gone away.
---
--- This mechanism can be used to cancel long operations on the server
--- if the client has disconnected before the response is ready.
---
--- Deprecated: the CloseNotifier interface predates Go's context package.
--- New code should use Request.Context instead.
---@class httpCloseNotifier
local httpCloseNotifier = {}

--- A Handler responds to an HTTP request.
---
--- ServeHTTP should write reply headers and data to the ResponseWriter
--- and then return. Returning signals that the request is finished; it
--- is not valid to use the ResponseWriter or read from the
--- Request.Body after or concurrently with the completion of the
--- ServeHTTP call.
---
--- Depending on the HTTP client software, HTTP protocol version, and
--- any intermediaries between the client and the Go server, it may not
--- be possible to read from the Request.Body after writing to the
--- ResponseWriter. Cautious handlers should read the Request.Body
--- first, and then reply.
---
--- Except for reading the body, handlers should not modify the
--- provided Request.
---
--- If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
--- that the effect of the panic was isolated to the active request.
--- It recovers the panic, logs a stack trace to the server error log,
--- and either closes the network connection or sends an HTTP/2
--- RST_STREAM, depending on the HTTP protocol. To abort a handler so
--- the client sees an interrupted response but the server doesn't log
--- an error, panic with the value ErrAbortHandler.
---@class httpHandler
local httpHandler = {}

--- A Header represents the key-value pairs in an HTTP header.
---
--- The keys should be in canonical form, as returned by
--- CanonicalHeaderKey.
---@class gzipHeader
local gzipHeader = {}

--- Clone returns a copy of h or nil if h is nil.
---@return gzipHeader
function gzipHeader:Clone() end

--- Set sets the header entries associated with key to the
--- single element value. It replaces any existing values
--- associated with key. The key is case insensitive; it is
--- canonicalized by textproto.CanonicalMIMEHeaderKey.
--- To use non-canonical keys, assign to the map directly.
---@param key string
---@param value string
function gzipHeader:Set(key, value) end

--- Get gets the first value associated with the given key. If
--- there are no values associated with the key, Get returns "".
--- It is case insensitive; textproto.CanonicalMIMEHeaderKey is
--- used to canonicalize the provided key. Get assumes that all
--- keys are stored in canonical form. To use non-canonical keys,
--- access the map directly.
---@param key string
---@return string
function gzipHeader:Get(key) end

--- Del deletes the values associated with key.
--- The key is case insensitive; it is canonicalized by
--- CanonicalHeaderKey.
---@param key string
function gzipHeader:Del(key) end

--- WriteSubset writes a header in wire format.
--- If exclude is not nil, keys where exclude[key] == true are not written.
--- Keys are not canonicalized before checking the exclude map.
---@param w ioWriter
---@param exclude any
---@return err
function gzipHeader:WriteSubset(w, exclude) end

--- Add adds the key, value pair to the header.
--- It appends to any existing values associated with key.
--- The key is case insensitive; it is canonicalized by
--- CanonicalHeaderKey.
---@param key string
---@param value string
function gzipHeader:Add(key, value) end

--- Values returns all values associated with the given key.
--- It is case insensitive; textproto.CanonicalMIMEHeaderKey is
--- used to canonicalize the provided key. To use non-canonical
--- keys, access the map directly.
--- The returned slice is not a copy.
---@param key string
---@return string[]
function gzipHeader:Values(key) end

--- Write writes a header in wire format.
---@param w ioWriter
---@return err
function gzipHeader:Write(w) end

--- A ResponseWriter interface is used by an HTTP handler to
--- construct an HTTP response.
---
--- A ResponseWriter may not be used after the Handler.ServeHTTP method
--- has returned.
---@class httpResponseWriter
local httpResponseWriter = {}

--- The Flusher interface is implemented by ResponseWriters that allow
--- an HTTP handler to flush buffered data to the client.
---
--- The default HTTP/1.x and HTTP/2 ResponseWriter implementations
--- support Flusher, but ResponseWriter wrappers may not. Handlers
--- should always test for this ability at runtime.
---
--- Note that even for ResponseWriters that support Flush,
--- if the client is connected through an HTTP proxy,
--- the buffered data may not reach the client until the response
--- completes.
---@class httpFlusher
local httpFlusher = {}

--- SameSite allows a server to define a cookie attribute making it impossible for
--- the browser to send this cookie along with cross-site requests. The main
--- goal is to mitigate the risk of cross-origin information leakage, and provide
--- some protection against cross-site request forgery attacks.
---
--- See https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 for details.
---@class httpSameSite
local httpSameSite = {}

--- Transport is an implementation of RoundTripper that supports HTTP,
--- HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
---
--- By default, Transport caches connections for future re-use.
--- This may leave many open connections when accessing many hosts.
--- This behavior can be managed using Transport's CloseIdleConnections method
--- and the MaxIdleConnsPerHost and DisableKeepAlives fields.
---
--- Transports should be reused instead of created as needed.
--- Transports are safe for concurrent use by multiple goroutines.
---
--- A Transport is a low-level primitive for making HTTP and HTTPS requests.
--- For high-level functionality, such as cookies and redirects, see Client.
---
--- Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2
--- for HTTPS URLs, depending on whether the server supports HTTP/2,
--- and how the Transport is configured. The DefaultTransport supports HTTP/2.
--- To explicitly enable HTTP/2 on a transport, use golang.org/x/net/http2
--- and call ConfigureTransport. See the package docs for more about HTTP/2.
---
--- Responses with status codes in the 1xx range are either handled
--- automatically (100 expect-continue) or ignored. The one
--- exception is HTTP status code 101 (Switching Protocols), which is
--- considered a terminal status and returned by RoundTrip. To see the
--- ignored 1xx responses, use the httptrace trace package's
--- ClientTrace.Got1xxResponse.
---
--- Transport only retries a request upon encountering a network error
--- if the request is idempotent and either has no body or has its
--- Request.GetBody defined. HTTP requests are considered idempotent if
--- they have HTTP methods GET, HEAD, OPTIONS, or TRACE; or if their
--- Header map contains an "Idempotency-Key" or "X-Idempotency-Key"
--- entry. If the idempotency key value is a zero-length slice, the
--- request is treated as idempotent but the header is not sent on the
--- wire.
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
---@field ProxyConnectHeader gzipHeader
---@field GetProxyConnectHeader any
---@field MaxResponseHeaderBytes number
---@field WriteBufferSize number
---@field ReadBufferSize number
---@field ForceAttemptHTTP2 boolean
local httpTransport = {}

--- CancelRequest cancels an in-flight request by closing its connection.
--- CancelRequest should only be called after RoundTrip has returned.
---
--- Deprecated: Use Request.WithContext to create a request with a
--- cancelable context instead. CancelRequest cannot cancel HTTP/2
--- requests.
---@param req httpRequest
function httpTransport:CancelRequest(req) end

--- RegisterProtocol registers a new protocol with scheme.
--- The Transport will pass requests using the given scheme to rt.
--- It is rt's responsibility to simulate HTTP request semantics.
---
--- RegisterProtocol can be used by other packages to provide
--- implementations of protocol schemes like "ftp" or "file".
---
--- If rt.RoundTrip returns ErrSkipAltProtocol, the Transport will
--- handle the RoundTrip itself for that one request, as if the
--- protocol were not registered.
---@param scheme string
---@param rt httpRoundTripper
function httpTransport:RegisterProtocol(scheme, rt) end

--- Clone returns a deep copy of t's exported fields.
---@return httpTransport
function httpTransport:Clone() end

--- CloseIdleConnections closes any connections which were previously
--- connected from previous requests but are now sitting idle in
--- a "keep-alive" state. It does not interrupt any connections currently
--- in use.
function httpTransport:CloseIdleConnections() end

--- a channel (just for the fun of it)
---@class httpChan
local httpChan = {}


---@param w httpResponseWriter
---@param req httpRequest
function httpChan:ServeHTTP(w, req) end

--- ProtocolError represents an HTTP protocol error.
---
--- Deprecated: Not all errors in the http package related to protocol errors
--- are of type ProtocolError.
---@class httpProtocolError
---@field ErrorString string
local httpProtocolError = {}


---@return string
function httpProtocolError:Error() end

--- Simple counter server. POSTing to it will set the value.
---@class httpCounter
local httpCounter = {}

--- This makes Counter satisfy the expvar.Var interface, so we can export
--- it directly.
---@return string
function httpCounter:String() end


---@param w httpResponseWriter
---@param req httpRequest
function httpCounter:ServeHTTP(w, req) end

--- A CookieJar manages storage and use of cookies in HTTP requests.
---
--- Implementations of CookieJar must be safe for concurrent use by multiple
--- goroutines.
---
--- The net/http/cookiejar package provides a CookieJar implementation.
---@class httpCookieJar
local httpCookieJar = {}

--- A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an
--- HTTP response or the Cookie header of an HTTP request.
---
--- See https://tools.ietf.org/html/rfc6265 for details.
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

--- String returns the serialization of the cookie for use in a Cookie
--- header (if only Name and Value are set) or a Set-Cookie response
--- header (if other fields are set).
--- If c is nil or c.Name is invalid, the empty string is returned.
---@return string
function httpCookie:String() end

--- Valid reports whether the cookie is valid.
---@return err
function httpCookie:Valid() end

--- PushOptions describes options for Pusher.Push.
---@class httpPushOptions
---@field Method string
---@field Header gzipHeader
local httpPushOptions = {}

--- MaxBytesError is returned by MaxBytesReader when its read limit is exceeded.
---@class httpMaxBytesError
---@field Limit number
local httpMaxBytesError = {}


---@return string
function httpMaxBytesError:Error() end

--- Response represents the response from an HTTP request.
---
--- The Client and Transport return Responses from servers once
--- the response headers have been received. The response body
--- is streamed on demand as the Body field is read.
---@class httpResponse
---@field Status string
---@field StatusCode number
---@field Proto string
---@field ProtoMajor number
---@field ProtoMinor number
---@field Header gzipHeader
---@field Body any
---@field ContentLength number
---@field TransferEncoding any
---@field Close boolean
---@field Uncompressed boolean
---@field Trailer gzipHeader
---@field Request httpRequest
---@field TLS any
local httpResponse = {}

--- Cookies parses and returns the cookies set in the Set-Cookie headers.
---@return any
function httpResponse:Cookies() end

--- Location returns the URL of the response's "Location" header,
--- if present. Relative redirects are resolved relative to
--- the Response's Request. ErrNoLocation is returned if no
--- Location header is present.
---@return urlURL, err
function httpResponse:Location() end

--- ProtoAtLeast reports whether the HTTP protocol used
--- in the response is at least major.minor.
---@param major number
---@param minor number
---@return boolean
function httpResponse:ProtoAtLeast(major, minor) end

--- Write writes r to w in the HTTP/1.x server response format,
--- including the status line, headers, body, and optional trailer.
---
--- This method consults the following fields of the response r:
---
---	StatusCode
---	ProtoMajor
---	ProtoMinor
---	Request.Method
---	TransferEncoding
---	Trailer
---	Body
---	ContentLength
---	Header, values for non-canonical keys will have unpredictable behavior
---
--- The Response Body is closed after it is sent.
---@param w ioWriter
---@return err
function httpResponse:Write(w) end

--- The Hijacker interface is implemented by ResponseWriters that allow
--- an HTTP handler to take over the connection.
---
--- The default ResponseWriter for HTTP/1.x connections supports
--- Hijacker, but HTTP/2 connections intentionally do not.
--- ResponseWriter wrappers may also not support Hijacker. Handlers
--- should always test for this ability at runtime.
---@class httpHijacker
local httpHijacker = {}

--- A ConnState represents the state of a client connection to a server.
--- It's used by the optional Server.ConnState hook.
---@class httpConnState
local httpConnState = {}


---@return string
function httpConnState:String() end

--- A FileSystem implements access to a collection of named files.
--- The elements in a file path are separated by slash ('/', U+002F)
--- characters, regardless of host operating system convention.
--- See the FileServer function to convert a FileSystem to a Handler.
---
--- This interface predates the fs.FS interface, which can be used instead:
--- the FS adapter function converts an fs.FS to a FileSystem.
---@class httpFileSystem
local httpFileSystem = {}

--- Pusher is the interface implemented by ResponseWriters that support
--- HTTP/2 server push. For more background, see
--- https://tools.ietf.org/html/rfc7540#section-8.2.
---@class httpPusher
local httpPusher = {}


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

--- Close immediately closes all active net.Listeners and any
--- connections in state StateNew, StateActive, or StateIdle. For a
--- graceful shutdown, use Shutdown.
---
--- Close does not attempt to close (and does not even know about)
--- any hijacked connections, such as WebSockets.
---
--- Close returns any error returned from closing the Server's
--- underlying Listener(s).
---@return err
function any:Close() end

--- ListenAndServeTLS listens on the TCP network address srv.Addr and
--- then calls ServeTLS to handle requests on incoming TLS connections.
--- Accepted connections are configured to enable TCP keep-alives.
---
--- Filenames containing a certificate and matching private key for the
--- server must be provided if neither the Server's TLSConfig.Certificates
--- nor TLSConfig.GetCertificate are populated. If the certificate is
--- signed by a certificate authority, the certFile should be the
--- concatenation of the server's certificate, any intermediates, and
--- the CA's certificate.
---
--- If srv.Addr is blank, ":https" is used.
---
--- ListenAndServeTLS always returns a non-nil error. After Shutdown or
--- Close, the returned error is ErrServerClosed.
---@param certFile string
---@param keyFile string
---@return err
function any:ListenAndServeTLS(certFile, keyFile) end

--- SetKeepAlivesEnabled controls whether HTTP keep-alives are enabled.
--- By default, keep-alives are always enabled. Only very
--- resource-constrained environments or servers in the process of
--- shutting down should disable them.
---@param v boolean
function any:SetKeepAlivesEnabled(v) end

--- Shutdown gracefully shuts down the server without interrupting any
--- active connections. Shutdown works by first closing all open
--- listeners, then closing all idle connections, and then waiting
--- indefinitely for connections to return to idle and then shut down.
--- If the provided context expires before the shutdown is complete,
--- Shutdown returns the context's error, otherwise it returns any
--- error returned from closing the Server's underlying Listener(s).
---
--- When Shutdown is called, Serve, ListenAndServe, and
--- ListenAndServeTLS immediately return ErrServerClosed. Make sure the
--- program doesn't exit and waits instead for Shutdown to return.
---
--- Shutdown does not attempt to close nor wait for hijacked
--- connections such as WebSockets. The caller of Shutdown should
--- separately notify such long-lived connections of shutdown and wait
--- for them to close, if desired. See RegisterOnShutdown for a way to
--- register shutdown notification functions.
---
--- Once Shutdown has been called on a server, it may not be reused;
--- future calls to methods such as Serve will return ErrServerClosed.
---@param ctx contextContext
---@return err
function any:Shutdown(ctx) end

--- RegisterOnShutdown registers a function to call on Shutdown.
--- This can be used to gracefully shutdown connections that have
--- undergone ALPN protocol upgrade or that have been hijacked.
--- This function should start protocol-specific graceful shutdown,
--- but should not wait for shutdown to complete.
---@param f any
function any:RegisterOnShutdown(f) end

--- ListenAndServe listens on the TCP network address srv.Addr and then
--- calls Serve to handle requests on incoming connections.
--- Accepted connections are configured to enable TCP keep-alives.
---
--- If srv.Addr is blank, ":http" is used.
---
--- ListenAndServe always returns a non-nil error. After Shutdown or Close,
--- the returned error is ErrServerClosed.
---@return err
function any:ListenAndServe() end

--- Serve accepts incoming connections on the Listener l, creating a
--- new service goroutine for each. The service goroutines read requests and
--- then call srv.Handler to reply to them.
---
--- HTTP/2 support is only enabled if the Listener returns *tls.Conn
--- connections and they were configured with "h2" in the TLS
--- Config.NextProtos.
---
--- Serve always returns a non-nil error and closes l.
--- After Shutdown or Close, the returned error is ErrServerClosed.
---@param l netListener
---@return err
function any:Serve(l) end

--- ServeTLS accepts incoming connections on the Listener l, creating a
--- new service goroutine for each. The service goroutines perform TLS
--- setup and then read requests, calling srv.Handler to reply to them.
---
--- Files containing a certificate and matching private key for the
--- server must be provided if neither the Server's
--- TLSConfig.Certificates nor TLSConfig.GetCertificate are populated.
--- If the certificate is signed by a certificate authority, the
--- certFile should be the concatenation of the server's certificate,
--- any intermediates, and the CA's certificate.
---
--- ServeTLS always returns a non-nil error. After Shutdown or Close, the
--- returned error is ErrServerClosed.
---@param l netListener
---@param certFile string
---@param keyFile string
---@return err
function any:ServeTLS(l, certFile, keyFile) end

--- The HandlerFunc type is an adapter to allow the use of
--- ordinary functions as HTTP handlers. If f is a function
--- with the appropriate signature, HandlerFunc(f) is a
--- Handler that calls f.
---@class httpHandlerFunc
local httpHandlerFunc = {}

--- ServeHTTP calls f(w, r).
---@param w httpResponseWriter
---@param r httpRequest
function httpHandlerFunc:ServeHTTP(w, r) end
