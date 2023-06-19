// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

type yockdConfGateway struct {
	Policy string                          `yaml:"policy"`
	Path   string                          `yaml:"path"`
	Rule   map[string]yockdConfGatewayRule `yaml:"rule"`
}

type yockdConfGatewayRule struct {
	Enable bool     `yaml:"enable"`
	Path   []string `yaml:"path"`
}

type yockdConfTLS struct {
	Cert string   `yaml:"cert"`
	Key  string   `yaml:"key"`
	Ca   []string `yaml:"ca"`
}
