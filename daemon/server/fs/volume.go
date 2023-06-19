// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/container"
)

const (
	FileTruncate = iota
	FileExpand
)

type Volume struct {
	files container.Tire[File]
	meta  VolumeMeta
	name  string
}

type VolumeMeta struct {
	createAt time.Time
	truncate bool
}

type volumeMetaSequence struct {
	Step  int
	Field string
	Type  int
}

var volumeMetaSequences = []volumeMetaSequence{
	{util.Int64Size, "createAt", util.Int64Type},
	{util.ByteSize, "truncate", util.BoolType},
}

func (meta *VolumeMeta) String() string {
	var buf []byte
	rv := reflect.ValueOf(*meta)
	for _, seq := range volumeMetaSequences {
		buf = append(buf, util.ByteTransfomer.AutoToBytes(rv.FieldByName(seq.Field))...)
	}
	return string(buf)
}

func newVolume(name string, info VolumeMeta) *Volume {
	return &Volume{
		files: container.NewWordTrie[File](),
		meta:  info,
		name:  name,
	}
}

func (v *Volume) Name() string {
	return v.name
}

func (v *Volume) Mount() {

}

func (v *Volume) Unmount() {

}

func (v *Volume) SwapIn() {
	buf := v.meta.String()
	fmt.Println(buf)
}

func (v *Volume) SwapOut() {

}
