// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocke

// MetaTable is an interface to abstract plist(darwin) and regedit(windows).
// In other posix os, uniform use of plist as MetaTable interface implement
type MetaTable interface {
	// SetValue set MetaTable's value and have two different rules:
	// regedit(windows): MetaValue ✔ MetaMap ✔ MetaArr x (MetaArr not work);
	// plist(darwin, posix): MetaValue ✔ MetaMap ✔ MetaArr ✔
	SetValue(MetaValue)
	// SafeSetValue set MetaTable's value when key isn't exist and have two different rules:
	// regedit(windows): MetaValue ✔ MetaMap ✔ MetaArr x (MetaArr not work);
	// plist(darwin, posix): MetaValue ✔ MetaMap ✔ MetaArr ✔
	SafeSetValue(MetaValue)
	// GetValue return MetaValue according to key
	GetValue(string) MetaValue
	// CreateSubTable have two different effect:
	// regedit(windows): create sub key and written file depond on its feture.
	// plist(darwin, posix): create sub element of map or array, but not be saved automatically
	// comparing regedit. It's required to call Write method save manually.
	CreateSubTable(string) MetaTable
	// Write to persist MetaValue in disk.
	// note: The regedit (windows) is written when it is created,
	// and this method is only valid for plist (darwin, posix).
	// The regedit is just an empty method.
	Write() error
	// Backup save a copy which could restore MetaValue
	Backup() error
	// Close to free MetaTable memory
	Close()
}

type (
	MetaValue any
	MetaMap   map[string]any
	MetaArr   []any
)
