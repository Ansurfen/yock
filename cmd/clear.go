package cmd

import "github.com/ansurfen/cushion/utils"

func Clear() error {
	var term *Terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = WindowsTerm("cls")
	default:
		term = PosixTerm("clear")
	}
	if _, err := term.Exec(&ExecOpt{}); err != nil {
		return err
	}
	return nil
}
