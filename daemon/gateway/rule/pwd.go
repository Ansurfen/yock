// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type PwdRule struct {
	index string
	pwd   string
	name  string
}

var _ Rule = (*PwdRule)(nil)

func NewPwdRule(name, pwd string) *PwdRule {
	return &PwdRule{
		name: name,
		pwd:  pwd,
	}
}

func (t *PwdRule) Kind() string {
	return "pwd"
}

func (t *PwdRule) Index() string {
	return t.index
}

func (t *PwdRule) Name() string {
	return t.name
}

func (t *PwdRule) String() string {
	return fmt.Sprintf(`{"type": "pwd","name": "%s","token": "%s"}`, t.name, t.pwd)
}

func (t *PwdRule) Release() {
	t.index = "pwd"
}

func (t *PwdRule) Check(ctx metadata.MD) error {
	v, ok := ctx[t.index]
	if !ok {
		return errors.New("token not found")
	}
	token := ""
	if len(v) > 0 {
		token = v[0]
	} else {
		return errors.New("token not found")
	}
	if token != t.pwd {
		return errors.New("pwd error")
	}
	return nil
}
