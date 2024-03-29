// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import "google.golang.org/grpc/metadata"

type Rule interface {
	Kind() string
	Name() string
	Index() string
	Check(metadata.MD) error
	String() string
}
