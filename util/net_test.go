// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestNet(t *testing.T) {
	urlStrings := []string{"https://www.example.com", "http://www.example.com", "www.example.com", "example.com", "example"}

	for _, urlString := range urlStrings {
		fmt.Println(IsURL(urlString))
	}
}
