// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"encoding/json"
	"fmt"
	"testing"
)

const jwtStr = `{"type": "jwt", "token": {"exp":10,"key":"yockd_key","sub":"yock"}}`

func TestJWT(t *testing.T) {
	jwt := NewJWTRule("default", map[string]any{
		"key": "yockd_key",
		"exp": 10,
		"sub": "yock",
	})
	fmt.Println(jwt)
}

func TestUnMarshal(t *testing.T) {
	var jwt JWTRule
	err := json.Unmarshal([]byte(jwtStr), &jwt)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
