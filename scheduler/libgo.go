// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"github.com/ansurfen/yock/lib/go/fmt"
	"github.com/ansurfen/yock/lib/go/net"
	"github.com/ansurfen/yock/lib/go/os"
	"github.com/ansurfen/yock/lib/go/path"
	"github.com/ansurfen/yock/lib/go/reflect"
	"github.com/ansurfen/yock/lib/go/regexp"
	"github.com/ansurfen/yock/lib/go/strings"
	"github.com/ansurfen/yock/lib/go/sync"
	"github.com/ansurfen/yock/lib/go/time"
)

var libgo = []loader{
	reflect.LoadReflect,
	fmt.LoadFmt,
	net.LoadNet,
	path.LoadPath,
	regexp.LoadRegexp,
	libstrings.LoadStrings,
	libtime.LoadTime,
	libsync.LoadSync,
	os.LoadOS,
}
