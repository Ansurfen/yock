// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package agent

import (
	"github.com/ansurfen/yock/daemon/gateway/rule"
	"github.com/ansurfen/yock/util"
)

type JWTAgent struct {
	rules map[string]*rule.JWTRule
}

func NewJWTAgent(files ...string) *JWTAgent {
	agent := &JWTAgent{
		rules: make(map[string]*rule.JWTRule),
	}
	for _, file := range files {
		err := agent.loadFromFile(file)
		if err != nil {
			panic(err)
		}
	}
	return agent
}

func (agent *JWTAgent) loadFromFile(file string) error {
	policies, err := util.OpenConf(file)
	if err != nil {
		return err
	}
	err = policies.ReadInConfig()
	if err != nil {
		return err
	}
	for k, v := range policies.AllSettings() {
		if vv, ok := v.(map[string]any); ok {
			agent.rules[k] = rule.NewJWTRule(k, vv)
		}
	}
	return nil
}

func (agent *JWTAgent) Release(name string, v map[string]any) rule.Rule {
	if tt, ok := agent.rules[name]; ok {
		return tt
	}
	t := rule.NewJWTRule(name, v)
	agent.rules[name] = t
	t.Release()
	return t
}

func (agent *JWTAgent) Del(name string) {
	delete(agent.rules, name)
}

func (agent *JWTAgent) Get(name string) rule.Rule {
	if r := agent.rules[name]; r != nil {
		return r
	}
	return nil
}
