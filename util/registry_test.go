//go:build windows
// +build windows

// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"

	"golang.org/x/sys/windows/registry"
)

func TestRegistry(t *testing.T) {
	NewWinEnv().ReadUserVar().DumpUserVar().SearchUserEnv(EnvVarSearchOpt{
		Case: false,
		Rule: "t",
		Reg:  true,
	})
}

func TestRegistryValue(t *testing.T) {
	esz := ExpandSZValue{
		val: []string{"a", "b", "c"},
	}
	fmt.Println(esz.ToString())
}

func TestEnvVarExport(t *testing.T) {
	NewWinEnv().ReadUserVar().
		SafeSetUserEnv("this_is_a_test", ExpandSZValue{val: []string{"a", "b"}}).
		DeleteUserVar(EnvVarDeleteOpt{
			Rules: []string{"this_is_a_test"},
			Safe:  false,
		}).ExportUserVar(EnvVarExportOpt{
		File: "tmp.ini",
	})
	NewWinEnv().LoadEnvVar(WinEnvVarLoadOpt{
		File: "tmp.ini",
		Spec: envVarUser,
	}).DumpUserVar()
}

func TestRegistryPage(t *testing.T) {
	page := CreateRegistryPage(registry.LOCAL_MACHINE, "SOFTWARE\\a_this_is_a_demo")
	page.SetValue("demo", BinaryValue{
		val: []byte("demo"),
	})
	page.DumpValue()
	defer page.Free()
	page.CreateSubKeys("\\a\\cccc\\a").CreateSubKeys("\\a\\b\\c\\d\\e")
	page.GetSubKeys("\\a\\b\\c\\d\\e").CreateSubKey("f")
	page.Walk(func(cur *RegistryPage, path string, level int, end bool) bool {
		fmt.Println(path)
		return true
	})
	page.GetSubKey("a").SetValue("a_value", SZValue{
		val: "a_value",
	})
	page.GetSubKey("a").Walk(func(cur *RegistryPage, path string, level int, end bool) bool {
		cur.Backup()
		cur.Delete()
		return true
	})
	fmt.Println()
	page.Walk(func(cur *RegistryPage, path string, level int, end bool) bool {
		fmt.Println(path)
		return true
	})
}

func TestRegistryRollback(t *testing.T) {
	RollbackRegistryPage(registry.LOCAL_MACHINE, "SOFTWARE")
	page := CreateRegistryPage(registry.LOCAL_MACHINE, "SOFTWARE\\a_this_is_a_demo")
	page.DumpValue()
	defer page.Free()
	page.Walk(func(cur *RegistryPage, path string, level int, end bool) bool {
		fmt.Println(path)
		return true
	})
}
