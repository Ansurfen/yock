// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filepathlib

import (
	"path/filepath"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadFilepath(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("filepath")
	lib.SetField(map[string]any{
		// functions
		"Abs":          filepath.Abs,
		"Walk":         filepath.Walk,
		"ToSlash":      filepath.ToSlash,
		"FromSlash":    filepath.FromSlash,
		"Clean":        filepath.Clean,
		"SplitList":    filepath.SplitList,
		"Rel":          filepath.Rel,
		"EvalSymlinks": filepath.EvalSymlinks,
		"IsLocal":      filepath.IsLocal,
		"Glob":         filepath.Glob,
		"Base":         filepath.Base,
		"Split":        filepath.Split,
		"VolumeName":   filepath.VolumeName,
		"WalkDir":      filepath.WalkDir,
		"IsAbs":        filepath.IsAbs,
		"Match":        filepath.Match,
		"Join":         filepath.Join,
		"Dir":          filepath.Dir,
		"Ext":          filepath.Ext,
		"HasPrefix":    filepath.HasPrefix,
		// constants
		"Separator":     filepath.Separator,
		"ListSeparator": filepath.ListSeparator,
		// variable
		"ErrBadPattern": filepath.ErrBadPattern,
		"SkipDir":       filepath.SkipDir,
		"SkipAll":       filepath.SkipAll,
	})
}
