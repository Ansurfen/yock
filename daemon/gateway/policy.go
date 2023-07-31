// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gateway

import (
	"errors"

	"github.com/ansurfen/yock/daemon/gateway/rule"
	"github.com/ansurfen/yock/daemon/user"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/container"
	"github.com/ansurfen/yock/ycho"
	"google.golang.org/grpc/metadata"
)

type Policy int

func (p Policy) String() string {
	switch p {
	case PolicyNULL:
		return "null"
	case PolicyRouter:
		return "router"
	case PolicyUser:
		return "user"
	default:
		panic("unreachable")
	}
}

const (
	PolicyNULL Policy = iota
	PolicyRouter
	PolicyUser
)

type PermPolicy interface {
	Policy() Policy
	Auth(ctx *metadata.MD, method string) error
	SetRule(key string, v ...rule.Rule)
	AppendRule(key string, v ...rule.Rule)
	String() string
}

var (
	_ PermPolicy = (*RouterPolicy)(nil)
	_ PermPolicy = (*UserPolicy)(nil)
	_ PermPolicy = (*NullPolicy)(nil)
)

type RouterPolicy struct {
	router container.Trie[[]rule.Rule]
}

func NewRouterPolicy() *RouterPolicy {
	return &RouterPolicy{
		router: container.WordTrieOf[[]rule.Rule](),
	}
}

func (p *RouterPolicy) Auth(ctx *metadata.MD, method string) error {
	pass := true
	for _, token := range p.router.Find("_G") {
		err := token.Check(*ctx)
		if err != nil {
			pass = false
			break
		}
	}
	if pass {
		for _, token := range p.router.Find(method) {
			err := token.Check(*ctx)
			if err != nil {
				pass = false
				break
			}
		}
	}
	return nil
}

func (p *RouterPolicy) SetRule(key string, tokens ...rule.Rule) {
	p.router.Insert(key, tokens)
}

func (p *RouterPolicy) AppendRule(key string, tokens ...rule.Rule) {

}

func (p *RouterPolicy) String() string {
	return ""
}

func (*RouterPolicy) Policy() Policy {
	return PolicyUser
}

var method2Perm = map[string]string{
	"/Yockd.YockDaemon/Ping":     "",
	"/Yockd.YockDaemon/Upload":   "write",
	"/Yockd.YockDaemon/Download": "read",
}

type UserPolicy struct {
	// user -> token[]
	tokens    map[string][]rule.Rule
	userGroup *user.UserGroup
}

func NewUserPolicy() *UserPolicy {
	return &UserPolicy{
		tokens:    make(map[string][]rule.Rule),
		userGroup: user.NewUserGroup(),
	}
}

func (*UserPolicy) Policy() Policy {
	return PolicyUser
}

func (p *UserPolicy) Auth(ctx *metadata.MD, method string) error {
	username := ""
	if u := ctx.Get("user"); len(u) > 0 {
		username = u[0]
	}

	u := p.userGroup.Get(username)

	perm, ok := method2Perm[method]
	if !ok {
		// return errors.New("invalid perm")
		ycho.Warnf("lack perm to set")
	}
	// try to login in user account
	if len(p.tokens) != 0 && len(username) == 0 {
		return errors.New("invalid user")
	}
	for _, token := range p.tokens[username] {
		err := token.Check(*ctx)
		if err != nil {
			return errors.New("authentication failed")
		}
	}

	// check whether user contains perm
	if !u.Contains(perm) {
		return util.ErrPermDenied
	}

	if len(username) != 0 {
		username = "user." + username
	}

	ycho.Infof("%s visit %s", username, method)
	return nil
}

func (p *UserPolicy) SetRule(key string, tokens ...rule.Rule) {
	p.tokens[key] = tokens
}

func (p *UserPolicy) AppendRule(key string, tokens ...rule.Rule) {
	p.tokens[key] = append(p.tokens[key], tokens...)
}

func (p *UserPolicy) String() string {
	return ""
}

type NullPolicy struct{}

func (null NullPolicy) Policy() Policy {
	return PolicyNULL
}

func (NullPolicy) Auth(*metadata.MD, string) error {
	return nil
}

func (NullPolicy) SetRule(string, ...rule.Rule) {}

func (NullPolicy) AppendRule(string, ...rule.Rule) {}

func (NullPolicy) String() string {
	return ""
}
