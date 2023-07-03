// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestCrypto(t *testing.T) {
	Key := `Yock Key`
	Raw := "Hello World!"
	enc := EncodeAESWithKey(Key, Raw)
	dec := DecodeAESWithKey(Key, enc)
	fmt.Printf("Raw: %s\nEnc: %s\nDec: %s\n", Raw, enc, dec)
}
