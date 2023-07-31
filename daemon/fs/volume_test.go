package fs

import (
	"fmt"
	du "github.com/ansurfen/yock/daemon/util"
	"testing"
)

func TestVolumeCRUD(t *testing.T) {
	// v := NewVolume("D")
	// v.Put(DirectoryEntry{})
}

func TestVolumeSnapshot(t *testing.T) {
	// for _, dir := range ParseDir("./disk", "/") {
	// 	fmt.Println(dir)
	// }
	// fmt.Println("--------")
	// fmt.Println(ParseDir("./dir.go", "./E"))
}

func TestVolume(t *testing.T) {
	v := NewVolume("D")
	for _, entry := range ParseDir("./testdata", "./") {
		v.Put(entry)
	}
	for _, entry := range ParseDir("./testdata", "./b") {
		v.Put(entry)
	}
	for _, entry := range ParseDir("./testdata", "./a/exf") {
		v.Put(entry)
	}
	for _, path := range v.List("%") {
		fmt.Println(ResolvePath("%") + ResolvePath(path))
	}
	fmt.Println("test same name")
	for _, path := range v.List("%a%exf") {
		fmt.Println(ResolvePath("%a%exf") + ResolvePath(path))
	}
	fmt.Println("--------------")
	fmt.Println(v.Get("%b%a%1.txt").Info(du.ID))
	fmt.Println(v.Get("%unknown"), v.List("%unknown"))
}
