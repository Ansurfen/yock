// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"strconv"
	"time"
)

// FmtTimestamp return a time string whom format is 2006-01-02
func FmtTimestamp(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// NowTimestamp returm current timestamp
func NowTimestamp() int64 {
	return time.Now().Unix()
}

// NowTimestamp returm current timestamp string
func NowTimestampByString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
