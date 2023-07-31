// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"strings"

	"github.com/ansurfen/yock/util"
	"github.com/spf13/viper"
)

type UserGroup struct {
	users map[string]*User
	conf  *viper.Viper
}

func NewUserGroup(file ...string) *UserGroup {
	group := &UserGroup{users: map[string]*User{}}
	if len(file) > 0 {
		group.Load(file[0])
	}
	return group
}

func (group *UserGroup) Load(file string) {
	conf, err := util.OpenConf(file)
	if err != nil {
		panic(err)
	}
	group.conf = conf
	for user, info := range conf.AllSettings() {
		u := newUser(user)
		if infos, ok := info.(map[string]any); ok {
			if perm, ok := infos["perm"].(string); ok {
				u.Grant(strings.Split(perm, ",")...)
			}
		}
		group.Add(u)
	}
}

func (group *UserGroup) Add(u *User) {
	group.users[u.Name()] = u
}

func (group *UserGroup) Get(name string) *User {
	if u, ok := group.users[name]; ok {
		return u
	}
	return &User{}
}

func (group *UserGroup) String() string {
	buf := ""
	i := 0
	userc := len(group.users)
	for _, user := range group.users {
		buf += user.String()
		if i != userc-1 {
			buf += "\n\n"
		}
		i++
	}
	return buf
}

type User struct {
	name  string
	perms [PermCount]bool
}

func newUser(name string) *User {
	return &User{name: name}
}

func (user *User) Clone() *User {
	return &User{}
}

func (user *User) Name() string {
	return user.name
}

// Contains returns true when user possess perm to be specified
func (user *User) Contains(perms ...string) bool {
	for _, perm := range perms {
		if len(perm) == 0 {
			continue
		}
		p := findPerm[perm]
		if !user.perms[p] {
			return false
		}
	}
	return true
}

func (user *User) String() string {
	buf := fmt.Sprintf("[%s]\nperm: ", user.name)
	perms := []string{}
	for perm, has := range user.perms {
		if has {
			perms = append(perms, Perm(perm).String())
		}
	}
	buf += strings.Join(perms, ", ")
	return buf
}

func (user *User) Grant(perms ...string) {
	for _, perm := range perms {
		perm = strings.TrimSpace(perm)
		if i, ok := findPerm[perm]; ok {
			if i == PermCount {
				for i := 0; i < PermCount; i++ {
					user.perms[i] = true
				}
				return
			}
			user.perms[i] = true
		}
	}
}

func (user *User) Revoke(perms ...string) {
	for _, perm := range perms {
		perm = strings.TrimSpace(perm)
		if i, ok := findPerm[perm]; ok {
			if i == PermCount {
				for i := 0; i < PermCount; i++ {
					user.perms[i] = false
				}
				return
			}
			user.perms[i] = false
		}
	}
}
