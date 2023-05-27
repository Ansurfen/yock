// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"fmt"
	"testing"
)

func TestDNS(t *testing.T) {
	dns := CreateDNS("global.json")
	fmt.Println(dns.GetDriver("yock"))
	fmt.Println(dns.PutDriver("yock", "https://", "https://"))
	fmt.Println(dns.GetDriver("yock"))
	fmt.Println(dns.PutPlugin("yock", "https://", "https://"))
}
