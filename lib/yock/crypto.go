// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
)

func LoadCrypto(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("crypto")
	lib.SetField(map[string]any{
		"md5":        util.MD5,
		"sha256":     util.SHA256,
		"encode_aes": util.EncodeAESWithKey,
		"decode_aes": util.DecodeAESWithKey,
	})
}
