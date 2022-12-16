package app

import "fmt"

type Trie struct {
	Val      rune
	Leaf     bool
	Children map[rune]*Trie
}

/** Initialize your data structure here. */
func NewTrie() *Trie {
	child := make(map[rune]*Trie)
	return &Trie{Val: 0, Children: child}
}

/** Inserts a word into the trie. */
func (trie *Trie) Insert(word string) {
	ptr := trie
	for _, c := range word {
		_, exists := ptr.Children[c]
		if !exists {
			child := make(map[rune]*Trie)
			ptr.Children[c] = &Trie{Val: c, Children: child}
		}
		ptr = ptr.Children[c]
	}
	ptr.Leaf = true
}

/** Returns if the word is in the trie. */
func (trie *Trie) Search(word string) bool {
	node := trie.search(word)
	return node != nil && node.Leaf
}

/** Returns if the word is in the trie. */
func (trie *Trie) search(word string) *Trie {
	ptr := trie
	for _, c := range word {
		ct, ok := ptr.Children[c]
		if !ok {
			return nil
		}
		ptr = ct
	}
	return ptr
}

func (trie *Trie) Delete(word string) {
	var parents []*Trie

	ptr := trie
	for _, c := range word {
		parents = append(parents, ptr)
		ct, ok := ptr.Children[c]
		if !ok {
			return
		}
		ptr = ct
	}

	if !ptr.Leaf {
		return
	}

	if len(ptr.Children) != 0 {
		ptr.Leaf = false
		return
	}

	for len(parents) > 0 {
		p := parents[len(parents)-1]
		parents = parents[:len(parents)-1]

		delete(p.Children, ptr.Val)
		if len(p.Children) != 0 || p.Leaf {
			break
		}
		ptr = p
	}
}

func (trie *Trie) Walk() {
	var walk func(string, *Trie)
	walk = func(pfx string, node *Trie) {
		if node == nil {
			return
		}

		if node.Val != 0 {
			pfx += string(node.Val)
		}

		if node.Leaf {
			fmt.Println(string(pfx))
		}

		for _, v := range node.Children {
			walk(pfx, v)
		}
	}
	walk("", trie)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (trie *Trie) StartsWith(prefix string) bool {
	node := trie.search(prefix)
	return node != nil
}
