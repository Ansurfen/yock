package cmd

import "github.com/ansurfen/cushion/utils"

// Clear clears the output on the screen
func Clear() error {
	var term *Terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = WindowsTerm("cls")
	default:
		term = PosixTerm("clear")
	}
	if _, err := term.Exec(&ExecOpt{Quiet: true}); err != nil {
		return err
	}
	return nil
}
