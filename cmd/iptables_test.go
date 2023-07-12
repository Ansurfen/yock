// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"testing"

	"github.com/ansurfen/yock/util/test"
)

func TestIPTablesList4Linux(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/iptables/linux",
	}).Range(func(data string) error {
		fmt.Println(IPTablesListLinux(string(data)))
		return nil
	})
}

func TestIPTablesList4Darwin(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/iptables/darwin",
	}).Range(func(data string) error {
		fmt.Println(IPTablesListDarwin(data))
		return nil
	})
}

func TestIPTablesList4WindowsOfUnicode(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path:    "./testdata/iptables/windows",
		Filters: []string{"b"},
	}).Range(func(data string) error {
		fmt.Println(IPTablesListWindows(data))
		return nil
	})
}

func TestIPTablesList4WindowsOfUTF8(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path:    "./testdata/iptables/windows",
		Filters: []string{"a"},
	}).Range(func(data string) error {
		fmt.Println(IPTablesListWindows(data))
		return nil
	})
}

func TestIPTablesList(t *testing.T) {
	fmt.Println(IPTablesList(IPTablesListOpt{
		Legacy: true,
		Name:   "MyRule",
	}))
}

func TestIPTablesAddOptCvt(t *testing.T) {
	addOpt := IPTablesOpOpt{
		Chain:       "INPUT",
		Name:        "MyRule",
		Protocol:    "tcp",
		Destination: "8080",
		Action:      "DROP",
		Legacy:      true,
		Op:          IPTablesAdd,
	}
	fmt.Println(addOpt.ToLinux())
	fmt.Println(addOpt.ToWindows())
	delOpt := addOpt
	delOpt.Op = IPTablesDel
	fmt.Println(delOpt.ToLinux())
	fmt.Println(delOpt.ToWindows())
}

func TestIPTablesAdd(t *testing.T) {
	fmt.Println(IPTablesOp(IPTablesOpOpt{
		Op:          IPTablesDel,
		Chain:       "INPUT",
		Name:        "MyRule",
		Protocol:    "tcp",
		Destination: "8080",
		Action:      "DROP",
	}))
}
