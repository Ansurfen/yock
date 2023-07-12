package yockp

import (
	"io/fs"
	"path/filepath"

	yockr "github.com/ansurfen/yock/runtime"
	yocks "github.com/ansurfen/yock/scheduler"
	"github.com/ansurfen/yock/util"
)

const script = ""

// ! try to build virtual file system
func (yockp *YockPack[T]) CC(entry string, main string) {
	main, err := filepath.Abs(main)
	if err != nil {
		panic(err)
	}
	filepath.Walk(entry, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			path, err = filepath.Abs(path)
			if err != nil {
				panic(err)
			}
			if path == main {
				_, err := util.ReadStraemFromFile(main)
				if err != nil {
					panic(err)
				}

				// ! pull yocks, and hook yocks
				yocks := yocks.New()
				// ! change the mode for load of yock path
				go yocks.EventLoop()
				fn := yockp.Compile(CompileOpt{
					VM: yocks.YockRuntime,
				}, main)
				// ! hook import function
				yocks.Eval("print(6)")
				err = yockr.LuaDoFunc(yocks.State().LState(), fn)
				if err != nil {
					panic(err)
				}
			}
		}
		return nil
	})

}
