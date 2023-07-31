// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	opts := SummaryOpts{
		Objectives: map[string]float64{
			"0.5":  0.05,
			"0.9":  0.01,
			"0.99": 0.001,
		},
		MaxAge: 5 * time.Second,
	}
	raw, err := json.Marshal(opts)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))
	recv := SummaryOpts{}
	err = json.Unmarshal(raw, &recv)
	if err != nil {
		panic(err)
	}
	fmt.Println(recv.Adapter())
}
