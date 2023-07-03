// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package libstrings

import (
	"strings"

	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
)

func LoadStrings(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("strings")
	lib.SetYFunction(map[string]yocki.YGFunction{
		"HasPrefix":     stringsHasPrefix,
		"HasSuffix":     stringsHasSuffix,
		"Contains":      stringsContains,
		"Join":          stringsJoin,
		"Cut":           stringsCut,
		"CutSuffix":     stringsCutSuffix,
		"CutPrefix":     stringsCutPrefix,
		"Clone":         stringsClone,
		"Compare":       stringsCompare,
		"ContainsAny":   stringsContainsAny,
		"ContainsRune":  stringsContainsRune,
		"Count":         stringsCount,
		"Replace":       stringsReplace,
		"ReplaceAll":    stringsReplaceAll,
		"Split":         stringsSplit,
		"SplitN":        stringsSplitN,
		"SplitAfterN":   stringsSplitAfterN,
		"Index":         stringsIndex,
		"NewReader":     stringsNewReader,
		"TrimSpace":     stringsTrimSpace,
		"LastIndex":     stringsLastIndex,
		"IndexByte":     stringsIndexByte,
		"IndexRune":     stringsIndexRune,
		"IndexAny":      stringsIndexAny,
		"LastIndexAny":  stringsLastIndexAny,
		"LastIndexByte": stringsLastIndexByte,
		"SplitAfter":    stringsSplitAfter,
		"Fields":        stringsFields,
		"Repeat":        stringsRepeat,
		"ToUpper":       stringsToUpper,
		"ToLower":       stringsToLower,
		"ToTitle":       stringsToTitle,
		"TrimPrefix":    stringsTrimPrefix,
		"TrimSuffix":    stringsTrimSuffix,
	})
	lib.SetField(map[string]any{
		"FieldsFunc":    strings.FieldsFunc,
		"Map":           strings.Map,
		"TrimLeftFunc":  strings.TrimLeftFunc,
		"TrimRightFunc": strings.TrimRightFunc,
		"TrimFunc":      strings.TrimFunc,
		"IndexFunc":     strings.IndexFunc,
		"LastIndexFunc": strings.LastIndexFunc,
	})
}

/*
* @param s string
* @param sub string
* @return number
 */
func stringsIndex(l yocki.YockState) int {
	idx := strings.Index(l.LState().CheckString(1), l.LState().CheckString(2))
	l.PushInt(idx)
	return 1
}

/*
* @param s string
* @param prefix string
* @return bool
 */
func stringsHasPrefix(l yocki.YockState) int {
	ok := strings.HasPrefix(l.LState().CheckString(1), l.LState().CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param suffix string
* @return bool
 */
func stringsHasSuffix(l yocki.YockState) int {
	ok := strings.HasSuffix(l.LState().CheckString(1), l.LState().CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param substr string
* @return bool
 */
func stringsContains(l yocki.YockState) int {
	ok := strings.Contains(l.LState().CheckString(1), l.LState().CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param elems string[]
* @param sep string
* @return string
 */
func stringsJoin(l yocki.YockState) int {
	elems := []string{}
	l.CheckTable(1).Value().ForEach(func(_, s lua.LValue) {
		elems = append(elems, s.String())
	})
	l.Push(lua.LString(strings.Join(elems, l.LState().CheckString(2))))
	return 1
}

/*
* @param s string
* @param sep string
* @return string, string, bool
 */
func stringsCut(l yocki.YockState) int {
	before, after, ok := strings.Cut(l.LState().CheckString(1), l.LState().CheckString(2))
	l.Push(lua.LString(before))
	l.Push(lua.LString(after))
	l.PushBool(ok)
	return 3
}

/*
* @param s string
* @param sep string
* @return string, bool
 */
func stringsCutSuffix(l yocki.YockState) int {
	before, found := strings.CutSuffix(l.LState().CheckString(1), l.LState().CheckString(2))
	l.Push(lua.LString(before))
	l.PushBool(found)
	return 2
}

/*
* @param s string
* @param sep string
* @return string, bool
 */
func stringsCutPrefix(l yocki.YockState) int {
	after, found := strings.CutPrefix(l.LState().CheckString(1), l.LState().CheckString(2))
	l.Push(lua.LString(after))
	l.PushBool(found)
	return 2
}

// @param s string
//
// @return string
func stringsClone(l yocki.YockState) int {
	l.Push(lua.LString(strings.Clone(l.LState().CheckString(1))))
	return 1
}

/*
* @param a string
* @param b string
* @return number
 */
func stringsCompare(l yocki.YockState) int {
	l.Push(lua.LNumber(strings.Compare(l.LState().CheckString(1), l.LState().CheckString(2))))
	return 1
}

/*
* @param s string
* @param chars string
* @return bool
 */
func stringsContainsAny(l yocki.YockState) int {
	ok := strings.ContainsAny(l.LState().CheckString(1), l.LState().CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param r string
* @return bool
 */
func stringsContainsRune(l yocki.YockState) int {
	ok := strings.ContainsRune(l.LState().CheckString(1), rune(l.LState().CheckString(2)[0]))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param substr string
* @return number
 */
func stringsCount(l yocki.YockState) int {
	l.Push(lua.LNumber(strings.Count(l.LState().CheckString(1), l.LState().CheckString(2))))
	return 1
}

/*
* @param s string
* @param old string
* @param new string
* @param n number
* @return string
 */
func stringsReplace(l yocki.YockState) int {
	l.Push(lua.LString(strings.Replace(l.LState().CheckString(1), l.LState().CheckString(2), l.LState().CheckString(3), l.LState().CheckInt(4))))
	return 1
}

/*
* @param s string
* @param old string
* @param new string
* @return string
 */
func stringsReplaceAll(l yocki.YockState) int {
	l.Push(lua.LString(strings.ReplaceAll(l.LState().CheckString(1), l.LState().CheckString(2), l.LState().CheckString(3))))
	return 1
}

/*
* @param s string
* @param sep string
* @return table
 */
func stringsSplit(l yocki.YockState) int {
	res := &lua.LTable{}
	for i, s := range strings.Split(l.LState().CheckString(1), l.LState().CheckString(2)) {
		res.Insert(i+1, lua.LString(s))
	}
	l.Push(res)
	return 1
}

// @param s string
//
// @return userdata
func stringsNewReader(l yocki.YockState) int {
	l.Pusha(strings.NewReader(l.LState().CheckString(1)))
	return 1
}

// @param s string
//
// @return string
func stringsTrimSpace(l yocki.YockState) int {
	l.PushString(strings.TrimSpace(l.LState().CheckString(1)))
	return 1
}

// @param s string
//
// @param substr string
//
// @return number
func stringsLastIndex(l yocki.YockState) int {
	l.PushInt(strings.LastIndex(l.LState().CheckString(1), l.LState().CheckString(2)))
	return 1
}

// @param s string
//
// @param c integer
//
// @return number
func stringsIndexByte(l yocki.YockState) int {
	l.PushInt(strings.IndexByte(l.LState().CheckString(1), byte(l.LState().CheckInt(2))))
	return 1
}

// @param s string
//
// @param r integer
//
// @return number
func stringsIndexRune(l yocki.YockState) int {
	l.PushInt(strings.IndexRune(l.LState().CheckString(1), rune(l.LState().CheckInt(2))))
	return 1
}

// @param s string
//
// @param chars string
//
// @return number
func stringsIndexAny(l yocki.YockState) int {
	l.PushInt(strings.IndexAny(l.LState().CheckString(1), l.LState().CheckString(2)))
	return 1
}

// @param s string
//
// @param chars string
//
// @return number
func stringsLastIndexAny(l yocki.YockState) int {
	l.PushInt(strings.LastIndexAny(l.LState().CheckString(1), l.LState().CheckString(2)))
	return 1
}

// @param s string
//
// @param c integer
//
// @return number
func stringsLastIndexByte(l yocki.YockState) int {
	l.PushInt(strings.LastIndexByte(l.LState().CheckString(1), byte(l.LState().CheckInt(2))))
	return 1
}

// @param s string
//
// @param sep string
//
// @param n integer
//
// @return string[]
func stringsSplitN(l yocki.YockState) int {
	res := strings.SplitN(l.LState().CheckString(1), l.LState().CheckString(2), l.LState().CheckInt(3))
	t := &lua.LTable{}
	for _, r := range res {
		t.Append(lua.LString(r))
	}
	l.Push(t)
	return 1
}

// @param s string
//
// @param sep string
//
// @param n integer
//
// @return string[]
func stringsSplitAfterN(l yocki.YockState) int {
	res := strings.SplitAfterN(l.LState().CheckString(1), l.LState().CheckString(2), l.LState().CheckInt(3))
	t := &lua.LTable{}
	for _, r := range res {
		t.Append(lua.LString(r))
	}
	l.Push(t)
	return 1
}

// @param s string
//
// @param sep string
//
// @return string[]
func stringsSplitAfter(l yocki.YockState) int {
	res := strings.SplitAfter(l.LState().CheckString(1), l.LState().CheckString(2))
	t := &lua.LTable{}
	for _, r := range res {
		t.Append(lua.LString(r))
	}
	l.Push(t)
	return 1
}

// @param s string
//
// @return string[]
func stringsFields(l yocki.YockState) int {
	res := strings.Fields(l.LState().CheckString(1))
	t := &lua.LTable{}
	for _, r := range res {
		t.Append(lua.LString(r))
	}
	l.Push(t)
	return 1
}

// @param s string
//
// @param count integer
//
// @return string
func stringsRepeat(l yocki.YockState) int {
	l.PushString(strings.Repeat(l.LState().CheckString(1), l.LState().CheckInt(2)))
	return 1
}

func stringsToUpper(l yocki.YockState) int {
	l.PushString(strings.ToUpper(l.LState().CheckString(1)))
	return 1
}

func stringsToLower(l yocki.YockState) int {
	l.PushString(strings.ToLower(l.LState().CheckString(1)))
	return 1
}

func stringsToTitle(l yocki.YockState) int {
	l.PushString(strings.ToTitle(l.LState().CheckString(1)))
	return 1
}

func stringsTrimPrefix(s yocki.YockState) int {
	s.PushString(strings.TrimPrefix(s.CheckString(1), s.CheckString(2)))
	return 1
}

func stringsTrimSuffix(s yocki.YockState) int {
	s.PushString(strings.TrimSuffix(s.CheckString(1), s.CheckString(2)))
	return 1
}
