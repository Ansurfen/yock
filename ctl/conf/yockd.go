// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

type yockDaemon struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	// MTL is abbreviation to max transfer length for file
	MTL      int                        `yaml:"MTL"`
	Name     string                     `yaml:"name"`
	SelfBoot bool                       `yaml:"self_boot"`
	Peer     map[string]yockdDaemonPeer `yaml:"peer"`
}

type yockdDaemonPeer struct {
	IP     string `yaml:"ip"`
	Port   int32  `yaml:"port"`
	Public bool   `yaml:"public"`
}
