// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import "errors"

type Trie[T any] interface {
	Insert(word string, v T) error
	Delete(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
	Find(word string) T
	FindNode(word string) (Trie[T], bool)
	FindChildren(word string) []string
	Keys() []string
	Value() T
	dump(word string) []string
}

var (
	_ Trie[nilType] = (*WordTrie[nilType])(nil)
	_ Trie[nilType] = (*MapTrie[nilType])(nil)
)

type WordTrie[T any] struct {
	isWord   bool
	children [26]*WordTrie[T]
	payload  T
}

func WordTrieOf[T any]() *WordTrie[T] {
	return &WordTrie[T]{}
}

func (t *WordTrie[T]) Value() T {
	return t.payload
}

func (t *WordTrie[T]) Insert(word string, v T) error {
	cur := t
	for i, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			cur.children[n] = WordTrieOf[T]()
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
	cur := t
	var v T
	for _, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return v
		}
		cur = cur.children[n]
	}
	if cur.isWord {
		return cur.payload
	}
	return v
}

func (t *WordTrie[T]) FindNode(word string) (Trie[T], bool) {
	cur := t
	for _, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return nil, false
		}
		cur = cur.children[n]
	}
	return cur, true
}

func (t *WordTrie[T]) dump(word string) (ret []string) {
	cur := t
	for i := 'a'; i <= 'z'; i++ {
		ch := i - 'a'
		if cur.children[ch] != nil {
			word += string(ch)
			cur = cur.children[ch]
			ret = append(ret, cur.dump(word)...)
			word = word[:len(word)-1]
			cur = t
		}
	}
	if cur.isWord {
		ret = append(ret, word)
	}
	return ret
}

func (t *WordTrie[T]) FindChildren(word string) []string {
	cur, ok := t.FindNode(word)
	if ok {
		return cur.dump("")
	}
	return nil
}

func (t *WordTrie[T]) Keys() []string { return nil }

type MapTrie[T any] struct {
	isWord   bool
	children map[rune]*MapTrie[T]
	payload  T
}

func MapTrieOf[T any]() *MapTrie[T] {
	return &MapTrie[T]{
		isWord:   false,
		children: make(map[rune]*MapTrie[T]),
	}
}

func (t *MapTrie[T]) Insert(word string, v T) error {
	if len(word) == 0 {
		return nil
	}
	cur := t
	for _, c := range word {
		_, ok := cur.children[c]
		if !ok {
			cur.children[c] = &MapTrie[T]{isWord: false, children: make(map[rune]*MapTrie[T])}
		}
		cur = cur.children[c]
	}
	if cur.isWord {
		return errors.New("duplicate key")
	}
	cur.isWord = true
	cur.payload = v
	return nil
}

func (t *MapTrie[T]) Delete(word string) {
	cur := t
	for _, c := range word {
		if cur.children[c] == nil {
			return
		}
		cur = cur.children[c]
	}
	if cur != nil && cur.isWord {
		cur.isWord = false
	}
}

func (t *MapTrie[T]) Search(word string) bool {
	cur := t
	for _, c := range word {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return cur.isWord
}

func (t *MapTrie[T]) StartsWith(prefix string) bool {
	cur := t
	for _, c := range prefix {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return true
}

func (t *MapTrie[T]) Find(word string) T {
	cur := t
	var v T
	for _, c := range word {
		if cur.children[c] == nil {
			return v
		}
		cur = cur.children[c]
	}
	if cur.isWord {
		return cur.payload
	}
	return v
}

func (t *MapTrie[T]) FindNode(word string) (Trie[T], bool) {
	cur := t
	for _, c := range word {
		if cur.children[c] == nil {
			return nil, false
		}
		cur = cur.children[c]
	}
	if cur == nil {
		return nil, true
	}
	return cur, true
}

func (t *MapTrie[T]) FindChildren(word string) (ret []string) {
	cur, ok := t.FindNode(word)
	if ok {
		ret = cur.dump("")
	}
	return
}

func (t *MapTrie[T]) dump(word string) (ret []string) {
	cur := t
	for ch := range t.children {
		if cur.children[ch] != nil {
			word += string(ch)
			cur = cur.children[ch]
			ret = append(ret, cur.dump(word)...)
			word = word[:len(word)-1]
			cur = t
		}
	}
	if cur.isWord {
		ret = append(ret, word)
	}
	return ret
}

func (t *MapTrie[T]) Keys() []string {
	return t.FindChildren("")
}

func (t *MapTrie[T]) Value() T {
	return t.payload
}
