// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ypm

import (
	"encoding/json"
	"fmt"

	"github.com/ansurfen/cushion/utils"
)

type Module struct {
	Dependency []yockDependency `json:"dependency"`
	Version    string           `json:"version"`
	Module     string           `json:"module"`
}

const BlankModule = `{"version": "v1", "module": "%s", "dependency": {}}`

type yockDependency struct{}

func OpenModule(file string) (*Module, error) {
	mod := &Module{}
	out, err := utils.ReadStraemFromFile(file)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(out, mod); err != nil {
		return nil, err
	}
	return mod, nil
}

func CreateModule(file, name string) (*Module, error) {
	err := utils.SafeWriteFile(file, []byte(fmt.Sprintf(BlankModule, name)))
	if err != nil {
		return nil, err
	}
	return OpenModule(file)
}

func (mod *Module) Write() {

}
