// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var yockCmd = &cobra.Command{
	Use:   "yock",
	Short: "Yock is a solution of cross platform to compose distributed build stream.",
	Long: `Yock is a solution of cross platform to compose distributed build stream. 
It's able to act as software package tool, like Homebrew, rpm, winget and so on. It also is
used for dependency manager (pip, npm, maven, etc.) of programming languages. On top of this,
yock also implements distributed build tasks based on grpc and goroutines (and can even build 
cluster for this). You can think of it as the lua version of the nodejs framework, except that
it focuses on composition and is more lightweight.`,
}

func Execute() {
	err := yockCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
