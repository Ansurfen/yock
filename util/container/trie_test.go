// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"testing"

	"github.com/ansurfen/yock/util/test"
)

func TestWordTrie(t *testing.T) {
	var tree Trie[int] = WordTrieOf[int]()
	tree.Insert("abc", 10)
	tree.Insert("a", 20)
	test.Assert(tree.Find("a") == 20)
	tree.Delete("a")
	test.Assert(tree.Find("a") == 0)
	test.Assert(tree.StartsWith("a"))
}

func TestMapTrie(t *testing.T) {
	var tree Trie[int] = MapTrieOf[int]()
	tree.Insert("abc", 10)
	tree.Insert("a", 20)
	test.Assert(tree.Find("a") == 20)
	test.Assert(len(tree.Keys()) == 2)
	tree.Delete("a")
	test.Assert(tree.Find("a") == 0)
	test.Assert(tree.StartsWith("a"))
	test.Assert(len(tree.Keys()) == 1)
}
