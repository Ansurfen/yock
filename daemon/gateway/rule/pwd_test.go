// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"fmt"
	"testing"
)

func TestPwd(t *testing.T) {
	pwd := NewPwdRule("root", "root")
	fmt.Println(pwd)
}
