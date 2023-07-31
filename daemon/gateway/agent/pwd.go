// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package agent

import (
	"github.com/ansurfen/yock/daemon/gateway/rule"
	"github.com/ansurfen/yock/util"
)

type PwdAgent struct {
	rules map[string]*rule.PwdRule
}

func NewPwdAgent(files ...string) *PwdAgent {
	agent := &PwdAgent{
		rules: make(map[string]*rule.PwdRule),
	}
	for _, file := range files {
		err := agent.loadFromFile(file)
		if err != nil {
			panic(err)
		}
	}
	return agent
}

func (agent *PwdAgent) loadFromFile(file string) error {
	policies, err := util.OpenConf(file)
	if err != nil {
		return err
	}
	err = policies.ReadInConfig()
	if err != nil {
		return err
	}
	for k, v := range policies.AllSettings() {
		if vv, ok := v.(string); ok {
			agent.rules[k] = rule.NewPwdRule(k, vv)
		}
	}
	return nil
}

func (agent *PwdAgent) Release(name string, v map[string]any) rule.Rule {
	if tt, ok := agent.rules[name]; ok {
		return tt
	}
	t := rule.NewPwdRule(name, v[""].(string))
	agent.rules[name] = t
	return t
}

func (agent *PwdAgent) Del(name string) {
	delete(agent.rules, name)
}

func (agent *PwdAgent) Get(name string) rule.Rule {
	if r := agent.rules[name]; r != nil {
		return r
	}
	return nil
}
