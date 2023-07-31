// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	u := newUser("admin1")
	u.Grant("all")
	u.Revoke("read")
	fmt.Println(u)
}

func TestUserGroup(t *testing.T) {
	group := NewUserGroup("./user.toml")
	fmt.Println(group)
}
