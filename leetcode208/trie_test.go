package leetcode

import (
	"testing"
)

func TestTrie(t *testing.T) {
	word := "apble"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search("apble")
	param_3 := obj.Search("apb")
	param_4 := obj.StartsWith("apb")
	param_5 := obj.StartsWith("abbel")
	param_6 := obj.StartsWith("apbel")
	if param_2 != true {
		t.Errorf("search: apple in trie, actual :", param_2)
	}
	if param_3 != false {
		t.Errorf("search: app not in trie, actual :", param_3)
	}
	if param_4 != true {
		t.Errorf("start with: app in trie, actual :", param_4)
	}
	if param_5 != false {
		t.Errorf("start with: appel in trie, actual :", param_5)
	}
	if param_6 != false {
		t.Errorf("start with: appel in trie, actual :", param_6)
	}
	obj.Insert("apb")
	p7 := obj.Search("apb")
	if p7 != true {
		t.Errorf("start with: appel in trie, actual :", p7)
	}
}

func TestTrieEmpty(t *testing.T) {
	word := ""
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search("apple")
	param_3 := obj.Search("app")
	param_4 := obj.StartsWith("app")
	param_5 := obj.StartsWith("abbel")
	param_6 := obj.StartsWith("appel")
	if param_2 != false {
		t.Errorf("search: apple in trie, actual :", param_2)
	}
	if param_3 != false {
		t.Errorf("search: app not in trie, actual :", param_3)
	}
	if param_4 != false {
		t.Errorf("start with: app in trie, actual :", param_4)
	}
	if param_5 != false {
		t.Errorf("start with: appel in trie, actual :", param_5)
	}
	if param_6 != false {
		t.Errorf("start with: appel in trie, actual :", param_6)
	}
}

func TestTrieLong(t *testing.T) {
	obj := Constructor()
	obj.Insert("app")
	obj.Insert("apple")
	obj.Insert("beer")
	obj.Insert("add")
	obj.Insert("jam")
	obj.Insert("rental")
	r2 := obj.Search("app")

	if r2 != true {
		t.Errorf("start with: appel in trie, actual :", r2)
	}
}
