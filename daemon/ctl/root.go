// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ctl

import (
	"os"

	"github.com/ansurfen/yock/daemon/server"
	"github.com/ansurfen/yock/daemon/conf"
	yocke "github.com/ansurfen/yock/env"
	"github.com/spf13/cobra"
)

type yockdParameter struct {
	ip   string
	port int
	// mtl is abbreviation to max transfer length for file
	mtl  int
	name string
}

var (
	opt     = yockdParameter{}
	rootCmd = &cobra.Command{
		Use:   `yockd`,
		Short: ``,
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if opt.port != 0 {
				yocke.GetEnv[*conf.YockdConf]().Conf().Grpc.Addr.Port = uint16(opt.port)
			}
			if len(opt.ip) != 0 {
				yocke.GetEnv[*conf.YockdConf]().Conf().Grpc.Addr.IP = opt.ip
			}
			s := api.New()
			defer s.Close()
			s.Run()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&opt.port, "port", "p", 0, "")
	rootCmd.PersistentFlags().IntVarP(&opt.mtl, "mtl", "m", 1024, "set max transfer length for file")
	rootCmd.PersistentFlags().StringVarP(&opt.ip, "ip", "i", "localhost", "")
	rootCmd.PersistentFlags().StringVarP(&opt.name, "name", "n", "", "")
}
