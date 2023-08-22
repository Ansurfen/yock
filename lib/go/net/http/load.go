// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package httplib

import (
	"net/http"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadHttp(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("http")
	lib.SetField(map[string]any{
		// functions
		"ListenAndServe": http.ListenAndServe,
		"TimeoutHandler": http.TimeoutHandler,
		// "ChanCreate":            http.ChanCreate,
		"NotFound":        http.NotFound,
		"StripPrefix":     http.StripPrefix,
		"ParseTime":       http.ParseTime,
		"MaxBytesHandler": http.MaxBytesHandler,
		"Error":           http.Error,
		// "HelloServer":           http.HelloServer,
		"NewFileTransport":  http.NewFileTransport,
		"FileServer":        http.FileServer,
		"ListenAndServeTLS": http.ListenAndServeTLS,
		"HandleFunc":        http.HandleFunc,
		// "FlagServer":            http.FlagServer,
		"Get":             http.Get,
		"SetCookie":       http.SetCookie,
		"NotFoundHandler": http.NotFoundHandler,
		"StatusText":      http.StatusText,
		"MaxBytesReader":  http.MaxBytesReader,
		"NewRequest":      http.NewRequest,
		"RedirectHandler": http.RedirectHandler,
		"Redirect":        http.Redirect,
		// "ArgServer":             http.ArgServer,
		// "DateServer":            http.DateServer,
		"PostForm":              http.PostForm,
		"ReadResponse":          http.ReadResponse,
		"NewRequestWithContext": http.NewRequestWithContext,
		"NewResponseController": http.NewResponseController,
		"ServeTLS":              http.ServeTLS,
		"NewServeMux":           http.NewServeMux,
		"ProxyFromEnvironment":  http.ProxyFromEnvironment,
		"Post":                  http.Post,
		"ServeContent":          http.ServeContent,
		"ReadRequest":           http.ReadRequest,
		"ParseHTTPVersion":      http.ParseHTTPVersion,
		"AllowQuerySemicolons":  http.AllowQuerySemicolons,
		"Handle":                http.Handle,
		"Serve":                 http.Serve,
		"DetectContentType":     http.DetectContentType,
		"Head":                  http.Head,
		"ServeFile":             http.ServeFile,
		"ProxyURL":              http.ProxyURL,
		// "Logger":                http.Logger,
		"FS":                 http.FS,
		"CanonicalHeaderKey": http.CanonicalHeaderKey,
		// constants
		"DefaultMaxIdleConnsPerHost": http.DefaultMaxIdleConnsPerHost,
		// variable
		"DefaultClient":           http.DefaultClient,
		"ErrUseLastResponse":      http.ErrUseLastResponse,
		"NoBody":                  http.NoBody,
		"ErrMissingFile":          http.ErrMissingFile,
		"ErrNotSupported":         http.ErrNotSupported,
		"ErrUnexpectedTrailer":    http.ErrUnexpectedTrailer,
		"ErrMissingBoundary":      http.ErrMissingBoundary,
		"ErrNotMultipart":         http.ErrNotMultipart,
		"ErrHeaderTooLong":        http.ErrHeaderTooLong,
		"ErrShortBody":            http.ErrShortBody,
		"ErrMissingContentLength": http.ErrMissingContentLength,
		"ErrNoCookie":             http.ErrNoCookie,
		"ErrNoLocation":           http.ErrNoLocation,
		"ErrBodyNotAllowed":       http.ErrBodyNotAllowed,
		"ErrHijacked":             http.ErrHijacked,
		"ErrContentLength":        http.ErrContentLength,
		"ErrWriteAfterFlush":      http.ErrWriteAfterFlush,
		"ServerContextKey":        http.ServerContextKey,
		"LocalAddrContextKey":     http.LocalAddrContextKey,
		"ErrAbortHandler":         http.ErrAbortHandler,
		"DefaultServeMux":         http.DefaultServeMux,
		"ErrServerClosed":         http.ErrServerClosed,
		"ErrHandlerTimeout":       http.ErrHandlerTimeout,
		"ErrLineTooLong":          http.ErrLineTooLong,
		"ErrBodyReadAfterClose":   http.ErrBodyReadAfterClose,
		"DefaultTransport":        http.DefaultTransport,
		"ErrSkipAltProtocol":      http.ErrSkipAltProtocol,
		"Client":                  netHttpClient,
	})
}

func netHttpClient() *http.Client {
	return &http.Client{}
}
