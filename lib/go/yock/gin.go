// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	"github.com/gin-gonic/gin"
)

func LoadGin(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("gin")
	lib.SetField(map[string]any{
		// functions
		"Recovery":                               gin.Recovery,
		"CustomRecovery":                         gin.CustomRecovery,
		"BasicAuth":                              gin.BasicAuth,
		"ErrorLogger":                            gin.ErrorLogger,
		"Logger":                                 gin.Logger,
		"EnableJsonDecoderDisallowUnknownFields": gin.EnableJsonDecoderDisallowUnknownFields,
		"Default":                                gin.Default,
		"DisableConsoleColor":                    gin.DisableConsoleColor,
		"EnableJsonDecoderUseNumber":             gin.EnableJsonDecoderUseNumber,
		"Bind":                                   gin.Bind,
		"CreateTestContext":                      gin.CreateTestContext,
		"IsDebugging":                            gin.IsDebugging,
		"New":                                    gin.New,
		"DisableBindValidation":                  gin.DisableBindValidation,
		"RecoveryWithWriter":                     gin.RecoveryWithWriter,
		"BasicAuthForRealm":                      gin.BasicAuthForRealm,
		"Mode":                                   gin.Mode,
		"Dir":                                    gin.Dir,
		"CreateTestContextOnly":                  gin.CreateTestContextOnly,
		"WrapF":                                  gin.WrapF,
		"ErrorLoggerT":                           gin.ErrorLoggerT,
		"LoggerWithConfig":                       gin.LoggerWithConfig,
		"SetMode":                                gin.SetMode,
		"ForceConsoleColor":                      gin.ForceConsoleColor,
		"CustomRecoveryWithWriter":               gin.CustomRecoveryWithWriter,
		"LoggerWithFormatter":                    gin.LoggerWithFormatter,
		"LoggerWithWriter":                       gin.LoggerWithWriter,
		"WrapH":                                  gin.WrapH,
		// constants
		"AuthUserKey":             gin.AuthUserKey,
		"MIMEJSON":                gin.MIMEJSON,
		"MIMEHTML":                gin.MIMEHTML,
		"MIMEXML":                 gin.MIMEXML,
		"MIMEXML2":                gin.MIMEXML2,
		"MIMEPlain":               gin.MIMEPlain,
		"MIMEPOSTForm":            gin.MIMEPOSTForm,
		"MIMEMultipartPOSTForm":   gin.MIMEMultipartPOSTForm,
		"MIMEYAML":                gin.MIMEYAML,
		"MIMETOML":                gin.MIMETOML,
		"BodyBytesKey":            gin.BodyBytesKey,
		"ContextKey":              gin.ContextKey,
		"ErrorTypeBind":           gin.ErrorTypeBind,
		"ErrorTypeRender":         gin.ErrorTypeRender,
		"ErrorTypePrivate":        gin.ErrorTypePrivate,
		"ErrorTypePublic":         gin.ErrorTypePublic,
		"ErrorTypeAny":            gin.ErrorTypeAny,
		"ErrorTypeNu":             gin.ErrorTypeNu,
		"PlatformGoogleAppEngine": gin.PlatformGoogleAppEngine,
		"PlatformCloudflare":      gin.PlatformCloudflare,
		"EnvGinMode":              gin.EnvGinMode,
		"DebugMode":               gin.DebugMode,
		"ReleaseMode":             gin.ReleaseMode,
		"TestMode":                gin.TestMode,
		"BindKey":                 gin.BindKey,
		"Version":                 gin.Version,
		// variable
		"DebugPrintRouteFunc": gin.DebugPrintRouteFunc,
		"DefaultWriter":       gin.DefaultWriter,
		"DefaultErrorWriter":  gin.DefaultErrorWriter,
	})
}
