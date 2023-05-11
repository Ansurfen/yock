package yock

import "os/user"

type WhoamiCmd struct{}

func NewWhoamiCmd() Cmd {
	return &WhoamiCmd{}
}

func (whoami *WhoamiCmd) Exec(string) ([]byte, error) {
	u, err := user.Current()
	return []byte(u.Username), err
}
