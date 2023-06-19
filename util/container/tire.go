// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import "errors"

type Tire[T any] interface {
	Insert(word string, v T) error
	Delete(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
	Find(word string) T
}

type WordTrie[T any] struct {
	isWord   bool
	children [26]*WordTrie[T]
	payload  T
}

func NewWordTrie[T any]() *WordTrie[T] {
	return &WordTrie[T]{}
}

func (t *WordTrie[T]) Insert(word string, v T) error {
	cur := t
	for i, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			cur.children[n] = NewWordTrie[T]()
		}
		cur = cur.children[n]
		if i == len(word)-1 {
			if cur.isWord {
				return errors.New("duplicate key")
			}
			cur.isWord = true
			cur.payload = v
		}
	}
	return nil
}

func (t *WordTrie[T]) Delete(word string) {
	cur := t
	for _, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return
		}
		cur = cur.children[n]
	}
	if cur != nil && cur.isWord {
		cur.isWord = false
	}
}

func (t *WordTrie[T]) Search(word string) bool {
	cur := t
	for _, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return cur.isWord
}

func (t *WordTrie[T]) StartsWith(prefix string) bool {
	cur := t
	for _, c := range prefix {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return true
}

func (t *WordTrie[T]) Find(word string) T {
	var v T
	return v
}
