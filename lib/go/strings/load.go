// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package libstrings

import (
	"strings"

	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	lua "github.com/yuin/gopher-lua"
)

func LoadStrings(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("strings")
	lib.SetYFunction(map[string]yockr.YGFunction{
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
func stringsIndex(l *yockr.YockState) int {
	idx := strings.Index(l.CheckString(1), l.CheckString(2))
	l.PushInt(idx)
	return 1
}

/*
* @param s string
* @param prefix string
* @return bool
 */
func stringsHasPrefix(l *yockr.YockState) int {
	ok := strings.HasPrefix(l.CheckString(1), l.CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param suffix string
* @return bool
 */
func stringsHasSuffix(l *yockr.YockState) int {
	ok := strings.HasSuffix(l.CheckString(1), l.CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param substr string
* @return bool
 */
func stringsContains(l *yockr.YockState) int {
	ok := strings.Contains(l.CheckString(1), l.CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param elems string[]
* @param sep string
* @return string
 */
func stringsJoin(l *yockr.YockState) int {
	elems := []string{}
	l.CheckTable(1).ForEach(func(_, s lua.LValue) {
		elems = append(elems, s.String())
	})
	l.Push(lua.LString(strings.Join(elems, l.CheckString(2))))
	return 1
}

/*
* @param s string
* @param sep string
* @return string, string, bool
 */
func stringsCut(l *yockr.YockState) int {
	before, after, ok := strings.Cut(l.CheckString(1), l.CheckString(2))
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
func stringsCutSuffix(l *yockr.YockState) int {
	before, found := strings.CutSuffix(l.CheckString(1), l.CheckString(2))
	l.Push(lua.LString(before))
	l.PushBool(found)
	return 2
}

/*
* @param s string
* @param sep string
* @return string, bool
 */
func stringsCutPrefix(l *yockr.YockState) int {
	after, found := strings.CutPrefix(l.CheckString(1), l.CheckString(2))
	l.Push(lua.LString(after))
	l.PushBool(found)
	return 2
}

// @param s string
//
// @return string
func stringsClone(l *yockr.YockState) int {
	l.Push(lua.LString(strings.Clone(l.CheckString(1))))
	return 1
}

/*
* @param a string
* @param b string
* @return number
 */
func stringsCompare(l *yockr.YockState) int {
	l.Push(lua.LNumber(strings.Compare(l.CheckString(1), l.CheckString(2))))
	return 1
}

/*
* @param s string
* @param chars string
* @return bool
 */
func stringsContainsAny(l *yockr.YockState) int {
	ok := strings.ContainsAny(l.CheckString(1), l.CheckString(2))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param r string
* @return bool
 */
func stringsContainsRune(l *yockr.YockState) int {
	ok := strings.ContainsRune(l.CheckString(1), rune(l.CheckString(2)[0]))
	l.PushBool(ok)
	return 1
}

/*
* @param s string
* @param substr string
* @return number
 */
func stringsCount(l *yockr.YockState) int {
	l.Push(lua.LNumber(strings.Count(l.CheckString(1), l.CheckString(2))))
	return 1
}

/*
* @param s string
* @param old string
* @param new string
* @param n number
* @return string
 */
func stringsReplace(l *yockr.YockState) int {
	l.Push(lua.LString(strings.Replace(l.CheckString(1), l.CheckString(2), l.CheckString(3), l.CheckInt(4))))
	return 1
}

/*
* @param s string
* @param old string
* @param new string
* @return string
 */
func stringsReplaceAll(l *yockr.YockState) int {
	l.Push(lua.LString(strings.ReplaceAll(l.CheckString(1), l.CheckString(2), l.CheckString(3))))
	return 1
}

/*
* @param s string
* @param sep string
* @return table
 */
func stringsSplit(l *yockr.YockState) int {
	res := &lua.LTable{}
	for i, s := range strings.Split(l.CheckString(1), l.CheckString(2)) {
		res.Insert(i+1, lua.LString(s))
	}
	l.Push(res)
	return 1
}

// @param s string
//
// @return userdata
func stringsNewReader(l *yockr.YockState) int {
	l.Pusha(strings.NewReader(l.CheckString(1)))
	return 1
}

// @param s string
//
// @return string
func stringsTrimSpace(l *yockr.YockState) int {
	l.PushString(strings.TrimSpace(l.CheckString(1)))
	return 1
}

// @param s string
//
// @param substr string
//
// @return number
func stringsLastIndex(l *yockr.YockState) int {
	l.PushInt(strings.LastIndex(l.CheckString(1), l.CheckString(2)))
	return 1
}

// @param s string
//
// @param c integer
//
// @return number
func stringsIndexByte(l *yockr.YockState) int {
	l.PushInt(strings.IndexByte(l.CheckString(1), byte(l.CheckInt(2))))
	return 1
}

// @param s string
//
// @param r integer
//
// @return number
func stringsIndexRune(l *yockr.YockState) int {
	l.PushInt(strings.IndexRune(l.CheckString(1), rune(l.CheckInt(2))))
	return 1
}

// @param s string
//
// @param chars string
//
// @return number
func stringsIndexAny(l *yockr.YockState) int {
	l.PushInt(strings.IndexAny(l.CheckString(1), l.CheckString(2)))
	return 1
}

// @param s string
//
// @param chars string
//
// @return number
func stringsLastIndexAny(l *yockr.YockState) int {
	l.PushInt(strings.LastIndexAny(l.CheckString(1), l.CheckString(2)))
	return 1
}

// @param s string
//
// @param c integer
//
// @return number
func stringsLastIndexByte(l *yockr.YockState) int {
	l.PushInt(strings.LastIndexByte(l.CheckString(1), byte(l.CheckInt(2))))
	return 1
}

// @param s string
//
// @param sep string
//
// @param n integer
//
// @return string[]
func stringsSplitN(l *yockr.YockState) int {
	res := strings.SplitN(l.CheckString(1), l.CheckString(2), l.CheckInt(3))
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
func stringsSplitAfterN(l *yockr.YockState) int {
	res := strings.SplitAfterN(l.CheckString(1), l.CheckString(2), l.CheckInt(3))
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
func stringsSplitAfter(l *yockr.YockState) int {
	res := strings.SplitAfter(l.CheckString(1), l.CheckString(2))
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
func stringsFields(l *yockr.YockState) int {
	res := strings.Fields(l.CheckString(1))
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
func stringsRepeat(l *yockr.YockState) int {
	l.PushString(strings.Repeat(l.CheckString(1), l.CheckInt(2)))
	return 1
}

func stringsToUpper(l *yockr.YockState) int {
	l.PushString(strings.ToUpper(l.CheckString(1)))
	return 1
}

func stringsToLower(l *yockr.YockState) int {
	l.PushString(strings.ToLower(l.CheckString(1)))
	return 1
}

func stringsToTitle(l *yockr.YockState) int {
	l.PushString(strings.ToTitle(l.CheckString(1)))
	return 1
}
