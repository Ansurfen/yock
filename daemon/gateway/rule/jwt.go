// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ansurfen/yock/ycho"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

const defaultJWTKey = "yockd_key"

var _ Rule = (*JWTRule)(nil)

type JWTRule struct {
	index string
	name  string
	token token
}

func NewJWTRule(name string, v map[string]any) *JWTRule {
	key := defaultJWTKey
	if k, ok := v["key"].(string); ok {
		key = k
	}
	method := ""
	if m, ok := v["method"].(string); ok {
		method = m
	}
	return &JWTRule{
		name:  name,
		index: key,
		token: token{
			claims: v,
			method: method,
		},
	}
}

func (t *JWTRule) String() string {
	b, err := json.Marshal(&t.token)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`{"type": "jwt","name": "%s","token": %s}`, t.name, b)
}

func (t *JWTRule) Kind() string {
	return "jwt"
}

func (t *token) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.claims)
}

func (t *JWTRule) UnmarshalJSON(data []byte) error {
	var jwt map[string]any
	err := json.Unmarshal(data, &jwt)

	var tt map[string]any
	if v, ok := jwt["token"].(map[string]any); ok {
		tt = v
	}
	method := ""
	if m, ok := tt["method"].(string); ok {
		method = m
	}

	*t = JWTRule{token: token{claims: tt, method: method}}
	return err
}

func (t *JWTRule) Index() string {
	return t.index
}

func (t *JWTRule) Name() string {
	return t.name
}

func (t *JWTRule) Release() error {
	var err error
	token, err := t.token.Parse().SigningString()
	if err != nil {
		ycho.Error(err)
	}
	ycho.Infof("jwt.%s token: %s", t.name, token)
	t.index = "token-x"
	return err
}

func (t *JWTRule) Check(ctx metadata.MD) error {
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
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return t, nil })
	if err != nil {
		ycho.Error(err)
		return err
	}
	return nil
}

type token struct {
	claims jwt.MapClaims
	method string
}

func (t *token) Method() string {
	return t.method
}

func (t *token) SigningMethod() jwt.SigningMethod {
	return t.parseMethod()
}

func (t *token) Claims() jwt.Claims {
	return t.claims
}

func (t *token) UnmarshalJSON(data []byte) error {
	var claims map[string]any
	err := json.Unmarshal(data, &claims)
	*t = token{claims: claims, method: claims["method"].(string)}
	return err
}

func (t *token) String() string {
	var buf string
	for k, v := range t.claims {
		tmp := ""
		switch k {
		case "key":
			tmp = fmt.Sprintf(`key = "%s"`, v)
		case "exp":
			tmp = fmt.Sprintf(`exp = %d`, v)
		case "sub":
			tmp = fmt.Sprintf(`sub = "%s"`, v)
		case "iss":
			tmp = fmt.Sprintf(`iss = "%s"`, v)
		case "aud":
			tmp = fmt.Sprintf(`aud = %v`, v)
		case "nbf":
			tmp = fmt.Sprintf(`nbf = %d`, v)
		case "method":
			tmp = fmt.Sprintf(`method = "%s"`, v)
		}
		buf += tmp + "\n"
	}
	return buf
}

func (t *token) parseMethod() jwt.SigningMethod {
	var method jwt.SigningMethod
	switch t.method {
	case "hs256":
		method = jwt.SigningMethodHS256
	case "hs384":
		method = jwt.SigningMethodHS384
	case "hs512":
		method = jwt.SigningMethodHS512
	case "es256":
		method = jwt.SigningMethodES256
	case "es384":
		method = jwt.SigningMethodES384
	case "es512":
		method = jwt.SigningMethodES512
	case "ps256":
		method = jwt.SigningMethodPS256
	case "ps384":
		method = jwt.SigningMethodPS384
	case "ps512":
		method = jwt.SigningMethodPS512
	case "rs256":
		method = jwt.SigningMethodRS256
	case "rs384":
		method = jwt.SigningMethodRS384
	case "rs512":
		method = jwt.SigningMethodRS512
	default:
		method = jwt.SigningMethodNone
	}
	return method
}

func (t *token) Parse() *jwt.Token {
	method := t.parseMethod()
	nowTime := time.Now()
	var (
		exp int64
		nbf int64
	)

	if v, ok := t.claims["exp"].(float64); ok {
		exp = int64(v)
	}

	if v, ok := t.claims["nbf"].(float64); ok {
		nbf = int64(v)
	}

	tmp := t.claims
	tmp["iat"] = nowTime
	if exp != 0 {
		tmp["exp"] = float64(nowTime.Add(time.Duration(exp * int64(time.Hour))).Unix())
	}
	if nbf != 0 {
		tmp["nbf"] = float64(nowTime.Add(time.Duration(nbf * int64(time.Hour))).Unix())
	}
	return jwt.NewWithClaims(method, tmp)
}
