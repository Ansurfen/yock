// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"golang.org/x/sys/windows/registry"
)

const (
	BatTemplate = `@echo off
ark exec %%* -n %s`
)

// WinEnv manage windows sys and user environment variable throught regedit
type WinEnv struct {
	sysVar  map[string]RegistryValue
	userVar map[string]RegistryValue
	sysReg  registry.Key
	userReg registry.Key
}

func NewWinEnv() *WinEnv {
	return &WinEnv{
		sysVar:  make(map[string]RegistryValue),
		userVar: make(map[string]RegistryValue),
	}
}

// GetModuleFileName return absolute path according to exe name
func (env *WinEnv) GetModuleFileName(exe string) string {
	path, err := Exec("where", exe)
	if err != nil {
		panic(err)
	}
	return string(path)
}

// GetExeEnvVar return enviroment variable which matchs with exe name
func (env *WinEnv) GetExeEnvVar(exe string) (string, bool) {
	if len(env.sysVar) == 0 {
		env.ReadSysEnv()
	}
	if len(env.userVar) == 0 {
		env.ReadUserVar()
	}
	path := env.GetModuleFileName(exe)
	ret := []string{}
	ret = append(ret, env.SearchUserEnv(EnvVarSearchOpt{
		Rule:       path,
		Reverse:    true,
		MatchValue: true,
	})...)
	if len(ret) > 0 {
		return ret[0], envVarUser
	}
	ret = append(ret, env.SearchSysEnv(EnvVarSearchOpt{
		Rule:       path,
		Reverse:    true,
		MatchValue: true,
	})...)
	if len(ret) > 0 {
		return ret[0], envVarSys
	}
	return "", false
}

// ReadUserVar read user variable from regedit
func (env *WinEnv) ReadUserVar() *WinEnv {
	userVarKeys, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err != nil {
		panic(err)
	}
	env.userReg = userVarKeys
	userVarNames, err := userVarKeys.ReadValueNames(-1)
	if err != nil {
		panic(err)
	}
	for _, name := range userVarNames {
		env.userVar[name] = GetValue(userVarKeys, name)
	}
	return env
}

// ReadUserVar read sys variable from regedit
func (env *WinEnv) ReadSysEnv() *WinEnv {
	sysVarKeys, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\ControlSet001\Control\Session Manager\Environment`, registry.ALL_ACCESS)
	if err != nil {
		panic(err)
	}
	env.sysReg = sysVarKeys
	sysVarNames, err := sysVarKeys.ReadValueNames(-1)
	if err != nil {
		panic(err)
	}
	for _, name := range sysVarNames {
		if name == "Path" {
			val, _, err := sysVarKeys.GetStringValue("Path")
			if err != nil {
				panic(err)
			}
			env.sysVar[name] = NewExpandSZValue(val)
		} else {
			env.sysVar[name] = GetValue(sysVarKeys, name)
		}
	}
	return env
}

// SafeSetUserEnv set user variable when key isn't exist
func (env *WinEnv) SafeSetUserEnv(k string, v RegistryValue) *WinEnv {
	if vv, ok := env.userVar[k]; ok {
		if v.Type() != registry.EXPAND_SZ {
			fmt.Println("key exist already")
		} else {
			val := fmt.Sprintf("%s;%s", vv.ToString(), v.(ExpandSZValue).ToString())
			if err := env.userReg.SetExpandStringValue(k, val); err != nil {
				panic(err)
			}
			env.userVar[k] = NewExpandSZValue(val)
		}
		return env
	}
	return env.SetUserEnv(k, v)
}

// SafeSetUserEnv set user variable
func (env *WinEnv) SetUserEnv(k string, v RegistryValue) *WinEnv {
	var err error
	switch vv := v.(type) {
	case SZValue:
		err = env.userReg.SetStringValue(k, vv.val)
	case ExpandSZValue:
		err = env.userReg.SetExpandStringValue(k, v.ToString())
	case DWordValue:
		err = env.userReg.SetDWordValue(k, uint32(vv.val))
	}
	if err != nil {
		panic(err)
	}
	env.userVar[k] = v
	return env
}

// SafeSetUserEnv set sys variable when key isn't exist
func (env *WinEnv) SafeSetSysEnv(k string, v RegistryValue) *WinEnv {
	if vv, ok := env.sysVar[k]; ok {
		if v.Type() != registry.EXPAND_SZ {
			fmt.Println("key exist already")
		} else {
			val := fmt.Sprintf("%s;%s", vv.ToString(), v.(ExpandSZValue).ToString())
			if err := env.sysReg.SetExpandStringValue(k, val); err != nil {
				panic(err)
			}
			env.sysVar[k] = NewExpandSZValue(val)
		}
		return env
	}
	return env.SetSysEnv(k, v)
}

// SafeSetUserEnv set sys variable
func (env *WinEnv) SetSysEnv(k string, v RegistryValue) *WinEnv {
	var err error
	switch vv := v.(type) {
	case SZValue:
		err = env.sysReg.SetStringValue(k, vv.val)
	case ExpandSZValue:
		err = env.userReg.SetExpandStringValue(k, strings.Join(vv.val, ";"))
	case DWordValue:
		err = env.sysReg.SetDWordValue(k, uint32(vv.val))
	}
	if err != nil {
		panic(err)
	}
	env.sysVar[k] = v
	return env
}

// DumpUserVar print user variable on console
func (env *WinEnv) DumpUserVar() *WinEnv {
	return env.dumpEnvVar(env.userVar)
}

// DumpSysVar print sys variable on console
func (env *WinEnv) DumpSysVar() *WinEnv {
	return env.dumpEnvVar(env.sysVar)
}

func backgroudColor(color, str string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(str)
}

func (env *WinEnv) dumpEnvVar(envVar map[string]RegistryValue) *WinEnv {
	for name, value := range envVar {
		fmt.Print(backgroudColor("#8866e9", name))
		switch v := value.(type) {
		case SZValue:
			fmt.Printf("(SZ): %s\n", v.val)
		case ExpandSZValue:
			fmt.Println("(ExpandSZ):")
			for _, sz := range v.val {
				fmt.Printf("  %s\n", sz)
			}
		default:
			fmt.Println()
		}
	}
	return env
}

// ExportUserVar export user varible to specify ini file.
// If no file is specified, it randomly generates a file name based on time
func (env *WinEnv) ExportUserVar(opt EnvVarExportOpt) *WinEnv {
	file := opt.File
	if len(file) == 0 {
		file = fmt.Sprintf("user_%s.ini", NowTimestampByString())
	}
	conf := NewRegistryValueFile(file)
	for name, value := range env.userVar {
		conf.SetType(name, value.Type())
		conf.SetValue(name, value.ToString())
	}
	conf.Write()
	return env
}

// ExportSysVar export sys varible to specify ini file.
// If no file is specified, it randomly generates a file name based on time
func (env *WinEnv) ExportSysVar(opt EnvVarExportOpt) *WinEnv {
	file := opt.File
	if len(file) == 0 {
		file = fmt.Sprintf("sys_%s.ini", NowTimestampByString())
	}
	conf := NewRegistryValueFile(file)
	for name, value := range env.sysVar {
		conf.SetType(name, value.Type())
		conf.SetValue(name, value.ToString())
	}
	conf.Write()
	return env
}

func (env *WinEnv) SearchSysEnv(opt EnvVarSearchOpt) []string {
	return env.searchEnvVar(opt, env.sysVar)
}

func (env *WinEnv) SearchUserEnv(opt EnvVarSearchOpt) []string {
	return env.searchEnvVar(opt, env.userVar)
}

func (env *WinEnv) searchEnvVar(opt EnvVarSearchOpt, envVar map[string]RegistryValue) []string {
	if opt.Reg && opt.MatchValue {
		return env.searchEnvVarWithRegexpMatchValue(opt, envVar)
	} else if opt.Reg {
		return env.searchEnvVarWithRegexp(opt, envVar)
	} else if opt.Reverse && opt.MatchValue {
		return env.searchEnvVarWithReverseMatchValue(opt, envVar)
	} else if opt.Reverse {
		return env.searchEnvVarWithReverse(opt, envVar)
	}
	return env.searchEnvVarDefault(opt, envVar)
}

// default match keys
func (env *WinEnv) searchEnvVarWithRegexp(opt EnvVarSearchOpt, envVar map[string]RegistryValue) []string {
	reg := regexp.MustCompile(opt.Rule)
	ret := make([]string, 0)
	for name := range envVar {
		tmp := name
		if opt.Case {
			tmp = strings.ToLower(tmp)
		}
		if reg.Match([]byte(tmp)) {
			ret = append(ret, name)
		}
	}
	return ret
}

func (env *WinEnv) searchEnvVarWithRegexpMatchValue(opt EnvVarSearchOpt, envVar map[string]RegistryValue) []string {
	reg := regexp.MustCompile(opt.Rule)
	ret := make([]string, 0)
	for name, value := range envVar {
		tmp := value.ToString()
		if opt.Case {
			tmp = strings.ToLower(tmp)
		}
		if reg.Match([]byte(tmp)) {
			ret = append(ret, name)
		}
	}
	return ret
}

func (env *WinEnv) searchEnvVarWithReverse(opt EnvVarSearchOpt, envVar map[string]RegistryValue) []string {
	ret := make([]string, 0)
	for name := range envVar {
		tmp := name
		if opt.Case {
			tmp = strings.ToLower(tmp)
		}
		if strings.HasPrefix(opt.Rule, tmp) {
			ret = append(ret, name)
		}
	}
	return ret
}

func (env *WinEnv) searchEnvVarWithReverseMatchValue(opt EnvVarSearchOpt, envVar map[string]RegistryValue) []string {
	ret := make([]string, 0)
	for name, value := range envVar {
		tmp := value.ToString()
		if opt.Case {
			tmp = strings.ToLower(tmp)
		}
		if strings.HasPrefix(opt.Rule, tmp) {
			ret = append(ret, name)
		}
	}
	return ret
}

func (env *WinEnv) searchEnvVarDefault(opt EnvVarSearchOpt, envVar map[string]RegistryValue) []string {
	ret := make([]string, 0)
	for name := range envVar {
		tmp := name
		if opt.Case {
			tmp = strings.ToLower(tmp)
		}
		if strings.HasPrefix(tmp, opt.Rule) {
			ret = append(ret, name)
		}
	}
	return ret
}

// DeleteUserVar delete key to be specified.
// If opt.Safe is true (safe is true in default), it'll backup kv to be deleted.
func (env *WinEnv) DeleteUserVar(opt EnvVarDeleteOpt) *WinEnv {
	fp := NewRegistryValueFile(fmt.Sprintf("%s.ini", NowTimestampByString()))
	for _, rule := range opt.Rules {
		regVal := GetValue(env.userReg, rule)
		fp.SetType(rule, regVal.Type())
		fp.SetValue(rule, regVal.ToString())
		env.userReg.DeleteValue(rule)
	}
	if opt.Safe {
		fp.Write()
	}
	return env
}

// DeleteSysVar delete key to be specified.
// If opt.Safe is true (safe is true in default), it'll backup kv to be deleted.
func (env *WinEnv) DeleteSysVar(opt EnvVarDeleteOpt) *WinEnv {
	fp := NewRegistryValueFile(fmt.Sprintf("%s.ini", NowTimestampByString()))
	for _, rule := range opt.Rules {
		regVal := GetValue(env.sysReg, rule)
		fp.SetType(rule, regVal.Type())
		fp.SetValue(rule, regVal.ToString())
		env.sysReg.DeleteValue(rule)
	}
	if opt.Safe {
		fp.Write()
	}
	return env
}

// LoadEnvVar restore enviroment variable from specify file
func (env *WinEnv) LoadEnvVar(opt WinEnvVarLoadOpt) *WinEnv {
	conf := viper.New()
	conf.SetConfigFile(opt.File)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	for name, value := range conf.AllSettings() {
		switch v := value.(type) {
		case map[string]any:
			var (
				valType uint32
				value   string
			)
			switch vv := v["type"].(type) {
			case string:
				i, err := strconv.Atoi(vv)
				if err != nil {
					panic(err)
				}
				valType = uint32(i)
			default:
			}
			switch vv := v["value"].(type) {
			case string:
				value = vv
			default:
			}
			if opt.Spec {
				if _, ok := env.sysVar[name]; !ok {
					env.sysVar[name] = NewRegistryValue(valType, value)
				}
			} else {
				if _, ok := env.userVar[name]; !ok {
					env.userVar[name] = NewRegistryValue(valType, value)
				}
			}
		default:
		}
	}
	return env
}

type EnvVarDeleteOpt struct {
	Case  bool
	Reg   bool
	Rules []string
	Safe  bool
}

type EnvVarSearchOpt struct {
	Reg        bool
	Case       bool
	Rule       string
	Reverse    bool
	MatchValue bool
}

type WinEnvVarLoadOpt struct {
	File string
	Spec bool // ? false -> user, true -> sys
}

type EnvVarExportOpt struct {
	File string
}

const (
	envVarUser = false
	envVarSys  = true
)

// RegistryWalk recurse to scan path along with root, you can add callback to effect each path
func RegistryWalk(root registry.Key, path string, level int, callback func(path string, level int, end bool) bool) {
	key, err := registry.OpenKey(root, path, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer key.Close()
	keyInfo, err := key.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	if keyInfo.SubKeyCount > 0 {
		subKeys, err := key.ReadSubKeyNames(-1)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, subKey := range subKeys {
			RegistryWalk(root, fmt.Sprintf("%s\\%s", path, subKey), level+1, callback)
		}
	}
	callback(path, level, true)
}

// RegistryPage manage key of specify path
type RegistryPage struct {
	key  registry.Key
	root registry.Key
	path string
}

// CreateRegistryPage open or create key
func CreateRegistryPage(root registry.Key, path string) *RegistryPage {
	path = strings.ReplaceAll(path, "/", "\\")
	key, err := registry.OpenKey(root, path, registry.ALL_ACCESS)
	if err != nil {
		if err.Error() != registry.ErrNotExist.Error() {
			panic(err)
		}
		key, exist, err := registry.CreateKey(root, path, registry.ALL_ACCESS)
		if exist {
			fmt.Println("fatal err: page exist ever")
		}
		if err != nil {
			panic(err)
		}
		return &RegistryPage{
			root: root,
			key:  key,
			path: path,
		}
	}
	return &RegistryPage{
		root: root,
		key:  key,
		path: path,
	}
}

// OpenRegistryPage open spcify key
func OpenRegistryPage(root registry.Key, path string) *RegistryPage {
	key, err := registry.OpenKey(root, path, registry.ALL_ACCESS)
	if err != nil {
		if err.Error() != registry.ErrNotExist.Error() {
			panic(err)
		}
		return nil
	}
	return &RegistryPage{
		root: root,
		key:  key,
		path: path,
	}
}

func (page *RegistryPage) SafeSetValue(k string, v RegistryValue) {
	rv := GetValue(page.key, k)
	switch rv.Type() {
	case registry.NONE:
		page.SetValue(k, v)
	case registry.EXPAND_SZ:
		val := fmt.Sprintf("%s;%s", v.ToString(), rv.(ExpandSZValue).ToString())
		page.SetValue(k, NewExpandSZValue(val))
	}
}

func (page *RegistryPage) SetValue(k string, v RegistryValue) {
	switch v.Type() {
	case registry.BINARY:
		page.key.SetBinaryValue(k, []byte(v.ToString()))
	case registry.SZ:
		page.key.SetStringValue(k, v.ToString())
	case registry.EXPAND_SZ:
		page.key.SetExpandStringValue(k, v.ToString())
	case registry.DWORD:
		i, err := strconv.Atoi(v.ToString())
		if err != nil {
			panic(err)
		}
		page.key.SetDWordValue(k, uint32(i))
	case registry.QWORD:
		i, err := strconv.Atoi(v.ToString())
		if err != nil {
			panic(err)
		}
		page.key.SetQWordValue(k, uint64(i))
	case registry.MULTI_SZ:
		page.key.SetStringValue(k, v.ToString())
	}
}

func (page *RegistryPage) CreateSubKey(subpath string) *RegistryPage {
	names, err := page.key.ReadSubKeyNames(-1)
	if err != nil {
		panic(err)
	}
	exist := false
	for _, name := range names {
		if name == subpath {
			exist = true
			break
		}
	}
	if !exist {
		return CreateRegistryPage(page.root, fmt.Sprintf("%s\\%s", page.path, subpath))
	}
	return OpenRegistryPage(page.root, fmt.Sprintf("%s\\%s", page.path, subpath))
}

func (page *RegistryPage) CreateSubKeys(subpaths string) *RegistryPage {
	paths := strings.Split(path.Join(filepath.ToSlash(subpaths)), "/")
	curPage := page
	for _, path := range paths {
		if len(path) > 0 {
			curPage = curPage.CreateSubKey(path)
		}
	}
	return page
}

func (page *RegistryPage) GetSubKey(subpath string) *RegistryPage {
	names, err := page.key.ReadSubKeyNames(-1)
	if err != nil {
		panic(err)
	}
	exist := false
	for _, name := range names {
		if name == subpath {
			exist = true
			break
		}
	}
	if !exist {
		return nil
	}
	return OpenRegistryPage(page.root, fmt.Sprintf("%s\\%s", page.path, subpath))
}

func (page *RegistryPage) GetSubKeys(subpaths string) *RegistryPage {
	paths := strings.Split(path.Join(filepath.ToSlash(subpaths)), "/")
	curPage := page
	for _, path := range paths {
		if len(path) > 0 {
			if cur := curPage.GetSubKey(path); cur != nil {
				curPage = cur
			} else {
				return nil
			}
		}
	}
	return curPage
}

// Walk recurse to scan path along with root, you can add callback to effect each path
func (page *RegistryPage) Walk(callback func(cur *RegistryPage, path string, level int, end bool) bool) {
	page.walkBuilder(page.root, page.path, 0, callback)
}

func (page *RegistryPage) walkBuilder(root registry.Key, path string, level int, callback func(cur *RegistryPage, path string, level int, end bool) bool) {
	key, err := registry.OpenKey(root, path, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer key.Close()
	keyInfo, err := key.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	if keyInfo.SubKeyCount > 0 {
		subKeys, err := key.ReadSubKeyNames(-1)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, subKey := range subKeys {
			page.walkBuilder(root, fmt.Sprintf("%s\\%s", path, subKey), level+1, callback)
		}
	}
	callback(OpenRegistryPage(page.root, path), path, level, false)
}

// DumpValue print values on key
func (page *RegistryPage) DumpValue() {
	for name, v := range page.dumpValue() {
		fmt.Println(name, v)
	}
}

func (page *RegistryPage) dumpValue() map[string]RegistryValue {
	values, err := page.key.ReadValueNames(-1)
	if err != nil {
		panic(err)
	}
	vals := make(map[string]RegistryValue)
	for _, value := range values {
		vals[value] = GetValue(page.key, value)
	}
	return vals
}

func (page *RegistryPage) Free() {
	page.key.Close()
}

func (page *RegistryPage) Delete() {
	registry.DeleteKey(page.root, page.path)
}

func (page *RegistryPage) RecurseDelete() {
	page.Walk(func(cur *RegistryPage, path string, level int, end bool) bool {
		cur.Delete()
		return true
	})
}

func (page *RegistryPage) SafeRecurseDelete() {
	page.Walk(func(cur *RegistryPage, path string, level int, end bool) bool {
		cur.Backup()
		cur.Delete()
		return true
	})
}

func (page *RegistryPage) Backup() error {
	if err := Mkdirs(page.path); err != nil {
		return err
	}
	fp := NewRegistryValueFile(fmt.Sprintf("%s/this.ini", page.path))
	for name, value := range page.dumpValue() {
		fp.SetType(name, value.Type())
		fp.SetValue(name, value.ToString())
	}
	fp.Write()
	return nil
}

// RollbackRegistryPage recurse to restore memory struct of registry from specify dir
func RollbackRegistryPage(root registry.Key, path string) {
	rollbackRegistryPageBuilder(root, path, 0)
}

func rollbackRegistryPageBuilder(root registry.Key, dir string, num int) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			rollbackRegistryPageBuilder(root, path.Join(dir, file.Name()), num+1)
		} else {
			conf := viper.New()
			target := path.Join(dir, file.Name())
			conf.SetConfigFile(target)
			if err := conf.ReadInConfig(); err != nil {
				panic(err)
			}
			flag := false
			for name, value := range conf.AllSettings() {
				flag = true
				switch v := value.(type) {
				case map[string]any:
					var (
						valType uint32
						value   string
					)
					switch vv := v["type"].(type) {
					case string:
						i, err := strconv.Atoi(vv)
						if err != nil {
							panic(err)
						}
						valType = uint32(i)
					default:
					}
					switch vv := v["value"].(type) {
					case string:
						value = vv
					default:
					}
					page := CreateRegistryPage(root, strings.ReplaceAll(path.Dir(target), "/", "\\"))
					page.SetValue(name, NewRegistryValue(valType, value))
					defer page.Free()
				default:
				}
			}
			if !flag {
				CreateRegistryPage(root, strings.ReplaceAll(path.Dir(target), "/", "\\"))
			}
		}
	}
}

type RegistryValueFile struct {
	conf *viper.Viper
	// regVals map[string]struct {
	// }
	file string
}

func NewRegistryValueFile(file string) *RegistryValueFile {
	fp := &RegistryValueFile{
		conf: viper.New(),
		file: file,
	}
	fp.conf.SetConfigType("ini")
	return fp
}

func (fp *RegistryValueFile) SetType(k string, v uint32) {
	fp.conf.Set(fmt.Sprintf("%s.type", k), v)
}

func (fp *RegistryValueFile) SetValue(k string, v string) {
	fp.conf.Set(fmt.Sprintf("%s.value", k), v)
}

func (fp *RegistryValueFile) Write() {
	fp.conf.WriteConfigAs(fp.file)
}

// DeleteRegistryKeys to delete specify keys
// It's safe recurse, which meant that all opearator will be stored.
func DeleteRegistryKeys(root string, keys []string) {
	for _, key := range keys {
		switch root {
		case "ROOT":
			CreateRegistryPage(registry.CLASSES_ROOT, key).SafeRecurseDelete()
		case "USER":
			CreateRegistryPage(registry.CURRENT_USER, key).SafeRecurseDelete()
		case "LOCAL_MACHINE":
			CreateRegistryPage(registry.LOCAL_MACHINE, key).SafeRecurseDelete()
		case "USERS":
			CreateRegistryPage(registry.USERS, key).SafeRecurseDelete()
		case "CURRENT_CONFIG":
			CreateRegistryPage(registry.CURRENT_CONFIG, key).SafeRecurseDelete()
		default:
		}
	}
}

// NewRegistryValue create a RegistryValue according to valType and value
func NewRegistryValue(valType uint32, value string) RegistryValue {
	switch valType {
	case registry.SZ:
		return SZValue{
			val: value,
		}
	case registry.EXPAND_SZ:
		return NewExpandSZValue(value)
	case registry.DWORD:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			panic(err)
		}
		return DWordValue{
			val: i,
		}
	case registry.QWORD:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			panic(err)
		}
		return QWordValue{
			val: i,
		}
	case registry.MULTI_SZ:
		return MultiSZValue{
			val: value,
		}
	case registry.BINARY:
		return BinaryValue{
			val: []byte(value),
		}
	}
	return NoneValue{}
}

type RegistryValue interface {
	Type() uint32
	ToString() string
}

type SZValue struct {
	val string
}

func (SZValue) Type() uint32 {
	return registry.SZ
}

func (sz SZValue) ToString() string {
	return sz.val
}

type ExpandSZValue struct {
	val []string
}

func NewExpandSZValue(str string) ExpandSZValue {
	return ExpandSZValue{
		val: strings.Split(str, ";"),
	}
}

func (ExpandSZValue) Type() uint32 {
	return registry.EXPAND_SZ
}

func (esz ExpandSZValue) ToString() string {
	return strings.Join(esz.val, ";")
}

type BinaryValue struct {
	val []byte
}

func (BinaryValue) Type() uint32 {
	return registry.BINARY
}

func (bny BinaryValue) ToString() string {
	return string(bny.val)
}

type DWordValue struct {
	val uint64
}

func (DWordValue) Type() uint32 {
	return registry.DWORD
}

func (dw DWordValue) ToString() string {
	return strconv.FormatInt(int64(dw.val), 10)
}

type QWordValue struct {
	val uint64
}

func (QWordValue) Type() uint32 {
	return registry.QWORD
}

func (qw QWordValue) ToString() string {
	return strconv.FormatInt(int64(qw.val), 10)
}

type NoneValue struct {
}

func (NoneValue) Type() uint32 {
	return registry.NONE
}

func (NoneValue) ToString() string {
	return ""
}

type MultiSZValue struct {
	val string
}

func (MultiSZValue) Type() uint32 {
	return registry.MULTI_SZ
}

func (msz MultiSZValue) ToString() string {
	return msz.val
}

// GetValue returns RegistryValue of k according to name
// If not exist, it'll return NoneValue.
func GetValue(k registry.Key, name string) RegistryValue {
	_, valtype, _ := k.GetValue(name, nil)
	switch valtype {
	case registry.SZ:
		val, _, err := k.GetStringValue(name)
		if err != nil {
			panic(err)
		}
		return SZValue{
			val: val,
		}
	case registry.EXPAND_SZ:
		val, _, err := k.GetStringValue(name)
		if err != nil {
			panic(err)
		}
		return ExpandSZValue{
			val: strings.Split(val, ";"),
		}
	case registry.BINARY:
		val, _, err := k.GetBinaryValue(name)
		if err != nil {
			panic(err)
		}
		return BinaryValue{
			val: val,
		}
	case registry.DWORD:
		val, _, err := k.GetIntegerValue(name)
		if err != nil {
			panic(err)
		}
		return DWordValue{
			val: val,
		}
	case registry.DWORD_BIG_ENDIAN:
		val, _, err := k.GetIntegerValue(name)
		if err != nil {
			panic(err)
		}
		return DWordValue{
			val: val,
		}
	case registry.MULTI_SZ:
		val, err := k.GetMUIStringValue(name)
		if err != nil {
			panic(err)
		}
		return MultiSZValue{
			val: val,
		}
	case registry.QWORD:
		val, _, err := k.GetIntegerValue(name)
		if err != nil {
			panic(err)
		}
		return QWordValue{
			val: val,
		}
	}
	return NoneValue{}
}
