// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kernel

import "github.com/ansurfen/yock/util"

type SingalStream struct {
	userSingals *util.SafeMap[bool]
}

func newSingalStream() *SingalStream {
	return &SingalStream{
		userSingals: util.NewSafeMap[bool](),
	}
}

func (stream *SingalStream) Wait(sig string) bool {
	v, ok := stream.userSingals.Get(sig)
	if !ok {
		stream.userSingals.SafeSet(sig, false)
	}
	return ok && v
}

func (stream *SingalStream) Notify(sig string) {
	stream.userSingals.SafeSet(sig, true)
}
