// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"encoding/binary"
	"errors"
	"math"
	"math/rand"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

const (
	BigEndian    = true
	LittleEndian = false

	ByteSize    = 1
	BoolSize    = 1
	Int16Size   = 2 * ByteSize
	Int32Size   = 4 * ByteSize
	Int64Size   = 8 * ByteSize
	Float32Size = 4 * ByteSize
	Float64Size = 8 * ByteSize

	ByteType = iota
	BoolType
	IntType
	Int16Type
	Int32Type
	Int64Type
	Float32Type
	Float64Type
	StringType
)

var ByteTransfomer BytesTransfomer

func SetBytesMode(_type bool) {
	ByteTransfomer.BytesEncode = NewBytesEncode(_type)
	ByteTransfomer.BytesDecode = NewBytesDecode(_type)
}

type BytesTransfomer struct {
	BytesEncode
	BytesDecode
}

func (transfomer BytesTransfomer) AutoToBytes(data any) []byte {
	switch v := data.(type) {
	case byte:
		return []byte{v}
	case bool:
		if v {
			return []byte{'1'}
		}
		return []byte{'0'}
	case []byte:
		return v
	case int:
		return transfomer.Int32ToBytes(int32(v))
	case int16:
		return transfomer.Int16ToBytes(v)
	case int32:
		return transfomer.Int32ToBytes(v)
	case int64:
		return transfomer.Int64ToBytes(v)
	case float32:
		return transfomer.Float32ToBytes(v)
	case float64:
		return transfomer.Float64ToBytes(v)
	case string:
		return []byte(v)
	}
	return nil
}

func (transfomer BytesTransfomer) AutoToType(raw []byte, _type int) any {
	switch _type {
	case BoolType:
		if raw[0] == '1' {
			return true
		}
		return false
	case ByteType:
		return raw
	case Int16Type:
		return transfomer.BytesDecode.BytesToInt16(raw)
	case Int32Type:
		return transfomer.BytesDecode.BytesToInt32(raw)
	case Int64Type:
		return transfomer.BytesDecode.BytesToInt64(raw)
	case Float32Type:
		return transfomer.BytesDecode.BytesToFloat32(raw)
	case Float64Type:
		return transfomer.BytesDecode.BytesToFloat64(raw)
	default:
		return string(raw)
	}
}

type BytesEncode struct {
	order binary.ByteOrder
}

func NewBytesEncode(_type bool) BytesEncode {
	encoder := BytesEncode{}
	if _type {
		encoder.order = &binary.BigEndian
	} else {
		encoder.order = &binary.LittleEndian
	}
	return encoder
}

func (encode BytesEncode) Float32ToBytes(f float32) []byte {
	bits := math.Float32bits(f)
	var buf = make([]byte, Float32Size)
	encode.order.PutUint32(buf, bits)
	return buf
}

func (encode BytesEncode) Float64ToBytes(f float64) []byte {
	bits := math.Float64bits(f)
	var buf = make([]byte, Float64Size)
	encode.order.PutUint64(buf, bits)
	return buf
}

func (encode BytesEncode) Int16ToBytes(i int16) []byte {
	var buf = make([]byte, Int16Size)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func (encode BytesEncode) Int32ToBytes(i int32) []byte {
	var buf = make([]byte, Int32Size)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func (encode BytesEncode) Int64ToBytes(i int64) []byte {
	var buf = make([]byte, Int64Size)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

type BytesDecode struct {
	order binary.ByteOrder
}

func NewBytesDecode(_type bool) BytesDecode {
	decoder := BytesDecode{}
	if _type {
		decoder.order = &binary.BigEndian
	} else {
		decoder.order = &binary.LittleEndian
	}
	return decoder
}

func (decode BytesDecode) BytesToFloat32(buf []byte) float32 {
	return math.Float32frombits(decode.order.Uint32(buf))
}

func (decode BytesDecode) BytesToFloat64(buf []byte) float64 {
	return math.Float64frombits(decode.order.Uint64(buf))
}

func (decode BytesDecode) BytesToInt16(buf []byte) int16 {
	return int16(binary.BigEndian.Uint16(buf))
}

func (decode BytesDecode) BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

func (decode BytesDecode) BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, Int64Size)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, Int32Size)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// Powered by github.com/syyongx/php2go
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0
		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}
		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}
		return sum
	}
	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

type ByteWalk struct {
	buf    []byte
	cursor int
}

func NewByteWalk(buf []byte) *ByteWalk {
	return &ByteWalk{
		cursor: 0,
		buf:    buf,
	}
}

func (bw *ByteWalk) Size() int {
	return len(bw.buf)
}

func (bw *ByteWalk) IsEnd() bool {
	return len(bw.buf) == bw.cursor
}

func (bw *ByteWalk) Buf() []byte {
	return bw.buf
}

func (bw *ByteWalk) Cursor() int {
	return bw.cursor
}

func (bw *ByteWalk) Next(step int) ([]byte, error) {
	if len(bw.buf) <= step+bw.cursor {
		return nil, errors.New("range out of index")
	}
	ret := bw.buf[bw.cursor : bw.cursor+step]
	bw.cursor += step
	return ret, nil
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandString return string of length n
func RandString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}
