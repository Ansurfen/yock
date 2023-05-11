package yock

import "os"

type PwdCmd struct{}

func NewPwdCmd() Cmd {
	return &PwdCmd{}
}

func (pwd *PwdCmd) Exec(args string) ([]byte, error) {
	wd, err := os.Getwd()
	return []byte(wd), err
}
