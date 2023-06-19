// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package agent

import "github.com/ansurfen/yock/daemon/server/gateway/rule"

type RuleAgent interface {
	Del(name string)
	Get(name string) rule.Rule
	Release(name string, v map[string]any) rule.Rule
}
