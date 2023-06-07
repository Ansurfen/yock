// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "errors"

type Trie[T any] struct {
	isWord   bool
	children [26]*Trie[T]
	payload  T
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{}
}

func (t *Trie[T]) Insert(word string, v T) error {
	cur := t
	for i, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			cur.children[n] = NewTrie[T]()
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

func (t *Trie[T]) Delete(word string) {
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

func (t *Trie[T]) Search(word string) bool {
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

func (t *Trie[T]) StartsWith(prefix string) bool {
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
