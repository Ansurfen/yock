package cmd

import (
	"flag"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
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
	return NilByte, cmd.Mv(cmd.MvOpt{}, mv.Src, mv.Dst)
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
	return NilByte, cmd.Cp(cmd.CpOpt{
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
	return NilByte, cmd.Rm(cmd.RmOpt{
		Safe: !rm.r,
	}, []string{rm.Path})
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
	var term *cmd.Terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = cmd.WindowsTerm("rmdir")
	default:
		term = cmd.PosixTerm("rmdir")
	}
	return term.Exec(&cmd.ExecOpt{})
}
