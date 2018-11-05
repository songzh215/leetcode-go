package leetcode

import ()

type Trie struct {
	Val      byte
	subNodes map[byte]*Trie
	end      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := Trie{}
	root.subNodes = make(map[byte]*Trie)
	return root
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	if this.subNodes[word[0]] == nil {
		node := Constructor()
		node.Val = word[0]
		this.subNodes[word[0]] = &node
	}
	if len(word) == 1 {
		this.subNodes[word[0]].end = true
	}
	this.subNodes[word[0]].Insert(word[1:len(word)])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		if len(this.subNodes) == 0 {
			return true
		} else {
			return false
		}
	}
	if len(word) == 1 && this.subNodes[word[0]] != nil && this.subNodes[word[0]].end {
		return true
	}
	if this.subNodes[word[0]] == nil {
		return false
	} else {
		if !this.subNodes[word[0]].Search(word[1:len(word)]) {
			return false
		}
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if len(prefix) == 1 && this.subNodes[prefix[0]] != nil {
		return true
	}
	if this.subNodes[prefix[0]] == nil {
		return false
	} else {
		if !this.subNodes[prefix[0]].StartsWith(prefix[1:len(prefix)]) {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
