// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocke

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"howett.net/plist"
)

// PlistFile manage plist, which serves for PosixMetaTable
type PlistFile struct {
	file string
	fp   *os.File
	v    any
}

// CreatePlistFile create specify plist when plist isn't exist
func CreatePlistFile(file string) (*PlistFile, error) {
	fp, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	pf := &PlistFile{
		fp:   fp,
		file: file,
	}
	if err := pf.read(); err != nil {
		return pf, err
	}
	return pf, err
}

// OpenPlistFile open specify plist
func OpenPlistFile(file string) (*PlistFile, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	pf := &PlistFile{
		fp:   fp,
		file: file,
	}
	if err := pf.read(); err != nil {
		return pf, err
	}
	return pf, err
}

func (pf *PlistFile) read() error {
	return plist.NewDecoder(pf.fp).Decode(&pf.v)
}

func (pf *PlistFile) Write() error {
	return plist.NewEncoder(pf.fp).Encode(pf.v)
}

// GetValue returns CFValue according to key
func (pf *PlistFile) GetValue(key string) CFValue {
	switch vv := pf.v.(type) {
	case cfDictionary:
		return any2CFValue(vv[key])
	default:
		if len(key) > 0 {
			return CFNone{}
		}
	}
	return any2CFValue(pf.v)
}

// GetBaseValue returns CFValue. If not exist in plist, it'll return CFNone.
func (pf *PlistFile) GetBaseValue() CFValue {
	switch vv := pf.v.(type) {
	case cfString, cfNumber, cfBool, cfData, cfUID, cfReal:
		return any2CFValue(vv)
	}
	return CFNone{}
}

// GetDict returns CFDictionary. If not exist in plist, it'll return a empty CFDictionary.
func (pf *PlistFile) GetDict() CFDictionary {
	switch v := pf.v.(type) {
	case cfDictionary:
		return CFDictionary{val: v}
	}
	return CFDictionary{}
}

// GetDict returns CFArray. If not exist in plist, it'll return a empty CFArray.
func (pf *PlistFile) GetArr() CFArray {
	switch v := pf.GetValue("").(type) {
	case CFArray:
		return v
	}
	return CFArray{}
}

// GetArrByField return CFArray according to field
func (pf *PlistFile) GetArrByField(field string) CFArray {
	fields := strings.Split(field, ".")
	if len(fields) == 0 {
		return CFArray{}
	}
	dict := pf.GetDict()
	ret := CFArray{}
	for i, f := range fields {
		switch v := dict.GetCFValue(f).(type) {
		case CFDictionary:
			dict = v
		case CFArray:
			if i == len(fields)-1 {
				ret = v
			}
		default:
			return CFArray{}
		}
	}
	return ret
}

func any2CFValue(v any) CFValue {
	switch vv := v.(type) {
	case string:
		return CFString{val: vv}
	case uint64:
		return CFNumber{val: vv}
	case int:
		return CFNumber{val: uint64(vv)}
	case bool:
		return CFBool{val: vv}
	case float64:
		return CFReal{val: vv}
	case []byte:
		return CFData{val: vv}
	case []any:
		return CFArray{val: vv}
	case cfDictionary:
		return CFDictionary{val: vv}
	case plist.UID:
		return CFUID{val: vv}
	default:
	}
	return CFNone{}
}

// Set set root element of plist
func (pf *PlistFile) Set(value any) {
	pf.v = value
}

// SetByField set element in sprcify field, which support recurse and meant that you can set a.b.c field.
func (pf *PlistFile) SetByField(field string, value CFValue) error {
	fields := strings.Split(field, ".")
	if len(fields) == 0 {
		return nil
	}
	t := pf.v
	sp := len(fields) - 1
	stack := make([]cfDictionary, sp)
	for i, f := range fields {
		switch v := t.(type) {
		case cfDictionary:
			if i == sp {
				v[f] = value
				break
			}
			stack[i] = v
			t = v[f]
		default:
			return errors.New("invalid type")
		}
	}
	return nil
}

// SetByIdx set element in specify position if position is available.
// For map, index will be convert string as key to set.
func (pf *PlistFile) SetByIdx(idx int, value any) error {
	if pf.Len() <= idx {
		return errors.New("out of range")
	}
	switch v := pf.v.(type) {
	case cfString, cfNumber, cfBool, cfData, cfUID, cfReal:
		pf.v = value
	case cfArray:
		v[idx] = value
		pf.v = v
	case cfDictionary:
		v[strconv.Itoa(idx)] = value
		pf.v = v
	}
	return nil
}

func (pf *PlistFile) Append(value any) error {

	return nil
}

// Len return plist length.
// Length of array or map depend on elements amount.
// For other type that include string, date, number and so on, it ever returns 1.
func (pf *PlistFile) Len() int {
	switch v := pf.v.(type) {
	case cfString, cfNumber, cfBool, cfData, cfUID, cfReal:
		return 1
	case cfArray:
		return len(v)
	case cfDictionary:
		return len(v)
	}
	return -1
}

func (pf *PlistFile) SafeSet(key string, value any) error {
	switch v := pf.v.(type) {
	case string:
		pf.v = value
	case uint64:
		pf.v = value
	case map[string]any:
		if _, ok := v[key]; !ok {
			v[key] = value
			pf.v = v
		}
	case []any:
		pf.v = append(v, value)
	default:
		fmt.Println(reflect.TypeOf(v).String())
	}
	return nil
}

func (pf *PlistFile) Backup() error {
	_, err := yockc.Exec(yockc.ExecOpt{},
		fmt.Sprintf("cp %s %s", pf.file,
			fmt.Sprintf("%s_%s.plist", util.Filename(pf.file), util.NowTimestampByString())))
	return err
}

func (pf *PlistFile) Free() {
	pf.fp.Close()
}

func GetPlistValue(v CFValue, name string) CFValue {
	switch v.Type() {
	case CF_DICT:
		dict := v.(CFDictionary)
		if d, ok := dict.val[name]; ok {
			return any2CFValue(d)
		}
	case CF_ARR:
	default:
	}
	return CFNone{}
}

const (
	CF_DICT = iota
	CF_ARR
	CF_STR
	CF_NUM
	CF_BOOL
	CF_DATA
	CF_UID
	CF_REAL
	CF_NONE
)

type CFValue interface {
	Type() uint8
}

type CFDictionary struct {
	val map[string]any
}

func (CFDictionary) Type() uint8 {
	return CF_DICT
}

func (c CFDictionary) Foreach(cb func(k string, v CFValue) bool) {
	for k, v := range c.val {
		if !cb(k, any2CFValue(v)) {
			break
		}
	}
}

func (c CFDictionary) GetCFValue(field string) CFValue {
	if v, ok := c.val[field]; ok {
		return any2CFValue(v)
	}
	return CFNone{}
}

func (c CFDictionary) Set(field string, v any) {
	c.val[field] = v
}

type CFArray struct {
	val []any
}

func (CFArray) Type() uint8 {
	return CF_ARR
}

func (c CFArray) Foreach(cb func(idx int, v CFValue) bool) {
	for i := 0; i < len(c.val); i++ {
		if !cb(i, any2CFValue(c.val[i])) {
			break
		}
	}
}

func (c CFArray) SetByIdx(idx int, v any) error {
	if len(c.val) <= idx {
		return errors.New("out of range")
	}
	c.val[idx] = v
	return nil
}

func (c CFArray) GetByIdx(idx int) CFValue {
	if len(c.val) <= idx {
		return CFNone{}
	}
	return any2CFValue(c.val[idx])
}

func (c CFArray) Copy() CFArray {
	arr := CFArray{}
	arr.val = append(arr.val, c.val...)
	return arr
}

type CFString struct {
	val string
}

func (CFString) Type() uint8 {
	return CF_STR
}

type CFNumber struct {
	val uint64
}

func (CFNumber) Type() uint8 {
	return CF_NUM
}

type CFBool struct {
	val bool
}

func (CFBool) Type() uint8 {
	return CF_BOOL
}

type CFData struct {
	val []byte
}

func (CFData) Type() uint8 {
	return CF_DATA
}

type CFUID struct {
	val plist.UID
}

func (CFUID) Type() uint8 {
	return CF_UID
}

type CFReal struct {
	val float64
}

func (CFReal) Type() uint8 {
	return CF_REAL
}

type CFNone struct{}

func (CFNone) Type() uint8 { return CF_NONE }

type (
	cfDictionary = map[string]any
	cfArray      = []any
	cfString     = string
	cfNumber     = int
	cfBool       = bool
	cfData       = []byte
	cfUID        = uint64
	cfReal       = float64
)
