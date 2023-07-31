-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

url = {}

---@class urlValues
---@field Encode fun(): string

--  Values maps a string key to a list of values.
--  It is typically used for query parameters and form values.
--  Unlike in the http.Header map, the keys in a Values map
--  are case-sensitive.
---@param v table<string, string[]>
---@return urlValues
function url.Values(v) end

--- ParseRequestURI parses a raw url into a URL structure. It assumes that
--- url was received in an HTTP request, so the url is interpreted
--- only as an absolute URI or an absolute path.
--- The string url is assumed not to have a #fragment suffix.
--- (Web browsers strip #fragment before sending the URL to a web server.)
---@param rawURL string
---@return urlURL, err
function url.ParseRequestURI(rawURL) end

--- User returns a Userinfo containing the provided username
--- and no password set.
---@param username string
---@return urlUserinfo
function url.User(username) end

--- PathUnescape does the inverse transformation of PathEscape,
--- converting each 3-byte encoded substring of the form "%AB" into the
--- hex-decoded byte 0xAB. It returns an error if any % is not followed
--- by two hexadecimal digits.
---
--- PathUnescape is identical to QueryUnescape except that it does not
--- unescape '+' to ' ' (space).
---@param s string
---@return string, err
function url.PathUnescape(s) end

--- Parse parses a raw url into a URL structure.
---
--- The url may be relative (a path, without a host) or absolute
--- (starting with a scheme). Trying to parse a hostname and path
--- without a scheme is invalid but may not necessarily return an
--- error, due to parsing ambiguities.
---@param rawURL string
---@return urlURL, err
function url.Parse(rawURL) end

--- JoinPath returns a URL string with the provided path elements joined to
--- the existing path of base and the resulting path cleaned of any ./ or ../ elements.
---@param base string
---@vararg string
---@return string, err
function url.JoinPath(base, ...) end

--- QueryEscape escapes the string so it can be safely placed
--- inside a URL query.
---@param s string
---@return string
function url.QueryEscape(s) end

--- UserPassword returns a Userinfo containing the provided username
--- and password.
---
--- This functionality should only be used with legacy web sites.
--- RFC 2396 warns that interpreting Userinfo this way
--- “is NOT RECOMMENDED, because the passing of authentication
--- information in clear text (such as URI) has proven to be a
--- security risk in almost every case where it has been used.”
---@param username string
---@param password string
---@return urlUserinfo
function url.UserPassword(username, password) end

--- QueryUnescape does the inverse transformation of QueryEscape,
--- converting each 3-byte encoded substring of the form "%AB" into the
--- hex-decoded byte 0xAB.
--- It returns an error if any % is not followed by two hexadecimal
--- digits.
---@param s string
---@return string, err
function url.QueryUnescape(s) end

--- PathEscape escapes the string so it can be safely placed inside a URL path segment,
--- replacing special characters (including /) with %XX sequences as needed.
---@param s string
---@return string
function url.PathEscape(s) end

--- ParseQuery parses the URL-encoded query string and returns
--- a map listing the values specified for each key.
--- ParseQuery always returns a non-nil map containing all the
--- valid query parameters found; err describes the first decoding error
--- encountered, if any.
---
--- Query is expected to be a list of key=value settings separated by ampersands.
--- A setting without an equals sign is interpreted as a key set to an empty
--- value.
--- Settings containing a non-URL-encoded semicolon are considered invalid.
---@param query string
---@return urlValues, err
function url.ParseQuery(query) end

--- Error reports an error and the operation and URL that caused it.
---@class execError
---@field Op string
---@field URL string
---@field Err err
local execError = {}


---@return string
function execError:Error() end

---@return boolean
function execError:Timeout() end

---@return boolean
function execError:Temporary() end

---@return err
function execError:Unwrap() end

---@class urlEscapeError
local urlEscapeError = {}


---@return string
function urlEscapeError:Error() end

---@class urlInvalidHostError
local urlInvalidHostError = {}


---@return string
function urlInvalidHostError:Error() end

--- A URL represents a parsed URL (technically, a URI reference).
---
--- The general form represented is:
---
---	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
---
--- URLs that do not start with a slash after the scheme are interpreted as:
---
---	scheme:opaque[?query][#fragment]
---
--- Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
--- A consequence is that it is impossible to tell which slashes in the Path were
--- slashes in the raw URL and which were %2f. This distinction is rarely important,
--- but when it is, the code should use the EscapedPath method, which preserves
--- the original encoding of Path.
---
--- The RawPath field is an optional field which is only set when the default
--- encoding of Path is different from the escaped path. See the EscapedPath method
--- for more details.
---
--- URL's String method uses the EscapedPath method to obtain the path.
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

--- Redacted is like String but replaces any password with "xxxxx".
--- Only the password in u.URL is redacted.
---@return string
function urlURL:Redacted() end

--- RequestURI returns the encoded path?query or opaque?query
--- string that would be used in an HTTP request for u.
---@return string
function urlURL:RequestURI() end

--- Hostname returns u.Host, stripping any valid port number if present.
---
--- If the result is enclosed in square brackets, as literal IPv6 addresses are,
--- the square brackets are removed from the result.
---@return string
function urlURL:Hostname() end

---@param text byte[]
---@return err
function urlURL:UnmarshalBinary(text) end

--- EscapedPath returns the escaped form of u.Path.
--- In general there are multiple possible escaped forms of any path.
--- EscapedPath returns u.RawPath when it is a valid escaping of u.Path.
--- Otherwise EscapedPath ignores u.RawPath and computes an escaped
--- form on its own.
--- The String and RequestURI methods use EscapedPath to construct
--- their results.
--- In general, code should call EscapedPath instead of
--- reading u.RawPath directly.
---@return string
function urlURL:EscapedPath() end

--- Parse parses a URL in the context of the receiver. The provided URL
--- may be relative or absolute. Parse returns nil, err on parse
--- failure, otherwise its return value is the same as ResolveReference.
---@param ref string
---@return urlURL, err
function urlURL:Parse(ref) end

--- ResolveReference resolves a URI reference to an absolute URI from
--- an absolute base URI u, per RFC 3986 Section 5.2. The URI reference
--- may be relative or absolute. ResolveReference always returns a new
--- URL instance, even if the returned URL is identical to either the
--- base or reference. If ref is an absolute URL, then ResolveReference
--- ignores base and returns a copy of ref.
---@param ref urlURL
---@return urlURL
function urlURL:ResolveReference(ref) end

--- String reassembles the URL into a valid URL string.
--- The general form of the result is one of:
---
---	scheme:opaque?query#fragment
---	scheme://userinfo@host/path?query#fragment
---
--- If u.Opaque is non-empty, String uses the first form;
--- otherwise it uses the second form.
--- Any non-ASCII characters in host are escaped.
--- To obtain the path, String uses u.EscapedPath().
---
--- In the second form, the following rules apply:
---   - if u.Scheme is empty, scheme: is omitted.
---   - if u.User is nil, userinfo@ is omitted.
---   - if u.Host is empty, host/ is omitted.
---   - if u.Scheme and u.Host are empty and u.User is nil,
---     the entire scheme://userinfo@host/ is omitted.
---   - if u.Host is non-empty and u.Path begins with a /,
---     the form host/path does not add its own /.
---   - if u.RawQuery is empty, ?query is omitted.
---   - if u.Fragment is empty, #fragment is omitted.
---@return string
function urlURL:String() end

--- IsAbs reports whether the URL is absolute.
--- Absolute means that it has a non-empty scheme.
---@return boolean
function urlURL:IsAbs() end

--- Query parses RawQuery and returns the corresponding values.
--- It silently discards malformed value pairs.
--- To check errors use ParseQuery.
---@return urlValues
function urlURL:Query() end

--- Port returns the port part of u.Host, without the leading colon.
---
--- If u.Host doesn't contain a valid numeric port, Port returns an empty string.
---@return string
function urlURL:Port() end

--- JoinPath returns a new URL with the provided path elements joined to
--- any existing path and the resulting path cleaned of any ./ or ../ elements.
--- Any sequences of multiple / characters will be reduced to a single /.
---@vararg string
---@return urlURL
function urlURL:JoinPath(...) end

--- EscapedFragment returns the escaped form of u.Fragment.
--- In general there are multiple possible escaped forms of any fragment.
--- EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment.
--- Otherwise EscapedFragment ignores u.RawFragment and computes an escaped
--- form on its own.
--- The String method uses EscapedFragment to construct its result.
--- In general, code should call EscapedFragment instead of
--- reading u.RawFragment directly.
---@return string
function urlURL:EscapedFragment() end

---@return byte[], err
function urlURL:MarshalBinary() end

--- The Userinfo type is an immutable encapsulation of username and
--- password details for a URL. An existing Userinfo value is guaranteed
--- to have a username set (potentially empty, as allowed by RFC 2396),
--- and optionally a password.
---@class urlUserinfo
local urlUserinfo = {}

--- String returns the encoded userinfo information in the standard form
--- of "username[:password]".
---@return string
function urlUserinfo:String() end

--- Username returns the username.
---@return string
function urlUserinfo:Username() end

--- Password returns the password in case it is set, and whether it is set.
---@return string, boolean
function urlUserinfo:Password() end

--- Values maps a string key to a list of values.
--- It is typically used for query parameters and form values.
--- Unlike in the http.Header map, the keys in a Values map
--- are case-sensitive.
---@class urlValues
local urlValues = {}

--- Set sets the key to value. It replaces any existing
--- values.
---@param key string
---@param value string
function urlValues:Set(key, value) end

--- Add adds the value to key. It appends to any existing
--- values associated with key.
---@param key string
---@param value string
function urlValues:Add(key, value) end

--- Del deletes the values associated with key.
---@param key string
function urlValues:Del(key) end

--- Has checks whether a given key is set.
---@param key string
---@return boolean
function urlValues:Has(key) end

--- Encode encodes the values into “URL encoded” form
--- ("bar=baz&foo=quux") sorted by key.
---@return string
function urlValues:Encode() end

--- Get gets the first value associated with the given key.
--- If there are no values associated with the key, Get returns
--- the empty string. To access multiple values, use the map
--- directly.
---@param key string
---@return string
function urlValues:Get(key) end
