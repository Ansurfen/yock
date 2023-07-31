// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

type YockdConfNet struct {
	Proxy map[string]yockdConfNetProxy `yaml:"proxy"`
	Stun  YockConfNetStun              `yaml:"stun"`
}

type yockdConfNetProxy struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type YockConfNetStun struct {
	RetryCnt int `yaml:"retryCnt"`
}
