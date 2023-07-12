// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// OpenConfFromPath unmarshal file which located in disk to memory according to path
func OpenConf(path string, opts ...viper.Option) (*viper.Viper, error) {
	conf := viper.NewWithOptions(opts...)
	conf.SetConfigFile(path)
	return conf, conf.ReadInConfig()
}

// ReadStraemFromFile return total data from specify file
func ReadStraemFromFile(file string) ([]byte, error) {
	fp, err := os.Open(file)
	if err != nil {
		return []byte(""), err
	}
	defer fp.Close()
	raw, err := ioutil.ReadAll(fp)
	if err != nil {
		return []byte(""), err
	}
	return raw, nil
}

// ReadStraemFromFile return data to be filter from specify file
func ReadLineFromFile(file string, filter func(string) string) ([]byte, error) {
	fp, err := os.Open(file)
	if err != nil {
		return []byte(""), err
	}
	defer fp.Close()
	fileScanner := bufio.NewScanner(fp)
	var ret []byte
	for fileScanner.Scan() {
		ret = append(ret, []byte(filter(fileScanner.Text()))...)
	}
	if err := fileScanner.Err(); err != nil {
		return []byte(""), err
	}
	return ret, nil
}

// ReadStraemFromFile return data to be filter from string
func ReadLineFromString(str string, filter func(string) string) ([]byte, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanLines)
	var ret []byte
	for scanner.Scan() {
		ret = append(ret, []byte(filter(scanner.Text()))...)
	}
	if err := scanner.Err(); err != nil {
		return []byte(""), err
	}
	return ret, nil
}

// WriteFile write data or create file to write data according to file
func WriteFile(file string, data []byte) error {
	fp, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fp.Close()
	fp.Write(data)
	return nil
}

// WriteFile write data or create file to write data according to file when file isn't exist
func SafeWriteFile(file string, data []byte) error {
	if ok, err := PathIsExist(file); err != nil {
		return err
	} else if !ok {
		if err := WriteFile(file, data); err != nil {
			return err
		}
	}
	return nil
}

// PathIsExist judge whether path exist. If exist, return true.
func PathIsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Mkdirs recurse to create path
func Mkdirs(path string) error {
	return os.MkdirAll(path, 0777)
}

// SafeMkdirs recurse to create path when path isn't exist
func SafeMkdirs(path string) error {
	if ok, err := PathIsExist(path); err != nil {
		return err
	} else if !ok {
		if err := Mkdirs(path); err != nil {
			return err
		}
	}
	return nil
}

// SafeBatchMkdirs recurse to create dirs when path isn't exist
func SafeBatchMkdirs(dirs []string) error {
	for _, dir := range dirs {
		if err := SafeMkdirs(dir); err != nil {
			return err
		}
	}
	return nil
}

// Exec automatically fit in os enviroment to execute command.
// windows 10+ -> powershell, others -> cmd;
// linux, darwin -> /bin/bash
func Exec(arg ...string) ([]byte, error) {
	switch CurPlatform.OS {
	case "windows":
		switch CurPlatform.Ver {
		case "10", "11":
			out, err := exec.Command("powershell", arg...).CombinedOutput()
			if err != nil {
				return out, err
			}
			return out, nil
		default:
			out, err := exec.Command("cmd", append([]string{"/C"}, arg...)...).CombinedOutput()
			if err != nil {
				return out, err
			}
			return out, nil
		}
	case "linux", "darwin":
		out, err := exec.Command("/bin/bash", append([]string{"/C"}, arg...)...).CombinedOutput()
		if err != nil {
			return out, err
		}
		return out, nil
	default:
	}
	return []byte(""), nil
}

// ExecStr automatically split string to string arrary, then call Exec to execute
func ExecStr(args string) ([]byte, error) {
	return Exec(strings.Fields(args)...)
}

// Filename returns the last element name of fullpath.
func Filename(fullpath string) string {
	filename := filepath.Base(fullpath)
	ext := path.Ext(filename)
	return filename[:len(filename)-len(ext)]
}

type PrintfOpt struct {
	MaxLen int
}

// Printf represent title and rows with tidy
func Prinf(opt PrintfOpt, title []string, rows [][]string) {
	if len(rows) <= 0 {
		for _, t := range title {
			fmt.Printf("%s ", t)
		}
		return
	}
	rowMaxLen := make([]int, len(title))
	for ri, row := range rows {
		for fi, field := range row {
			if fieldLen := len(field); opt.MaxLen <= fieldLen {
				rowMaxLen[fi] = opt.MaxLen
				rows[ri][fi] = fmt.Sprintf("%s...", field[:opt.MaxLen-3])
			} else if rowMaxLen[fi] < fieldLen {
				rowMaxLen[fi] = fieldLen
			}
		}
	}
	for ti, t := range title {
		if tLen := len(t); rowMaxLen[ti] <= tLen {
			fmt.Printf("%s ", t)
			rowMaxLen[ti] = tLen
		} else {
			fmt.Printf("%s%s ", t, strings.Repeat(" ", rowMaxLen[ti]-tLen))
		}
		if ti == len(title)-1 {
			fmt.Println()
		}
	}
	for _, row := range rows {
		for fi, field := range row {
			if fLen := len(field); rowMaxLen[fi] <= fLen {
				fmt.Printf("%s ", field)
				rowMaxLen[fi] = fLen
			} else {
				fmt.Printf("%s%s ", field, strings.Repeat(" ", rowMaxLen[fi]-fLen))
			}
			if fi == len(row)-1 {
				fmt.Println()
			}
		}
	}
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
