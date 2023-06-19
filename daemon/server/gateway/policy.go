// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gateway

import (
	"errors"

	"github.com/ansurfen/yock/daemon/server/gateway/rule"
	"github.com/ansurfen/yock/daemon/server/user"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/container"
	"google.golang.org/grpc/metadata"
)

type Policy int

const (
	PolicyRouter Policy = iota
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
)

type RouterPolicy struct {
	router container.Tire[[]rule.Rule]
}

func NewRouterPolicy() *RouterPolicy {
	return &RouterPolicy{
		router: container.NewWordTrie[[]rule.Rule](),
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
	} else {
		return util.ErrUserNotFound
	}

	u := p.userGroup.Get(username)
	if u == nil {
		return util.ErrUserNotFound
	}

	perm, ok := method2Perm[method]
	if !ok {
		return errors.New("invalid perm")
	}

	for _, token := range p.tokens[username] {
		err := token.Check(*ctx)
		if err != nil {
			return errors.New("err token")
		}
	}

	if !u.Contains(perm) {
		return util.ErrPermDenied
	}

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
