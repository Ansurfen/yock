// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

// type FIFO[T util.Comparable] struct {
// 	cap  int
// 	size int
// 	data container.HashMap[T, Entity]
// 	find map[T]Entity
// }

// func NewFIFO[T util.Comparable](cap int) *FIFO[T] {
// 	return &FIFO[T]{
// 		cap: cap,
// 		data: container.NewHashMap[T, Entity](func(key T, cap int) T {
// 			return key
// 		}),
// 		size: 0,
// 	}
// }

// func (fifo *FIFO[T]) Get(k T) Entity {
// 	if v, ok := fifo.find[k]; !ok {
// 		return nil
// 	} else {
// 		return v
// 	}
// }

// func (fifo *FIFO[T]) Put(k T, v Entity) {
// 	if fifo.cap == 0 {
// 		return
// 	}
// 	if v, ok := fifo.find[k]; ok {
// 		fifo.data.Del(k)
// 		fifo.data.Put(k, v)
// 	} else {
// 		if fifo.cap == fifo.size {
// 			// fifo.data.Del() // 删掉最后一个元素
// 			// delete(fifo.find, ) // del后要把元素的key给我
// 			fifo.find[k] = v
// 		}
// 	}
// }

// func (fifo *FIFO[T]) Cap() int {
// 	return fifo.cap
// }
