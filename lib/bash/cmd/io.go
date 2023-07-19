package cmd

import (
	"flag"

	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
)

type MoveCmd struct {
	Src string
	Dst string
}

func NewMoveCmd() Cmd {
	return &MoveCmd{}
}

func (mv *MoveCmd) Exec(args string) ([]byte, error) {
	initCmd(mv, args, func(cli *flag.FlagSet, cc *MoveCmd) {
	}, map[string]uint8{}, func(cc *MoveCmd, s string) error {
		if len(cc.Src) == 0 {
			cc.Src = s
		} else {
			cc.Dst = s
		}
		return nil
	})
	return NilByte, yockc.Mv(yockc.MvOpt{}, mv.Src, mv.Dst)
}

type CpCmd struct {
	Src string
	Dst string
	r   bool
}

func NewCpCmd() Cmd {
	return &CpCmd{}
}

func (cp *CpCmd) Exec(args string) ([]byte, error) {
	initCmd(cp, args, func(cli *flag.FlagSet, cc *CpCmd) {
		cli.BoolVar(&cp.r, "r", false, "")
	}, map[string]uint8{"-r": FlagBool}, func(cc *CpCmd, s string) error {
		if len(cc.Src) == 0 {
			cc.Src = s
		} else {
			cc.Dst = s
		}
		return nil
	})
	return NilByte, yockc.Cp(yockc.CpOpt{
		Recurse: cp.r,
	}, cp.Src, cp.Dst)
}

type RmCmd struct {
	Path string
	r    bool
}

func NewRmCmd() Cmd {
	return &RmCmd{}
}

func (rm *RmCmd) Exec(args string) ([]byte, error) {
	initCmd(rm, args, func(cli *flag.FlagSet, cc *RmCmd) {
		cli.BoolVar(&cc.r, "r", false, "")
	}, map[string]uint8{
		"-r": FlagBool,
	}, func(cc *RmCmd, s string) error {
		rm.Path = s
		return nil
	})
	return NilByte, yockc.Rm(yockc.RmOpt{
		Safe: !rm.r,
	}, rm.Path)
}

type RmdirCmd struct {
	Path string
}

func NewRmdirCmd() Cmd {
	return &RmdirCmd{}
}

func (rm *RmdirCmd) Exec(args string) ([]byte, error) {
	initCmd(rm, args, func(cli *flag.FlagSet, cc *RmdirCmd) {
	}, map[string]uint8{}, func(cc *RmdirCmd, s string) error {
		rm.Path = s
		return nil
	})
	var term *yockc.Terminal
	switch util.CurPlatform.OS {
	case "windows":
		term = yockc.WindowsTerm("rmdir")
	default:
		term = yockc.PosixTerm("rmdir")
	}
	return term.Exec(&yockc.ExecOpt{})
}
