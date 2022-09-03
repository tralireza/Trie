package Trie

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
}

type Trie26 struct {
	IsWord   bool
	Children [26]*Trie26
}

func (t *Trie26) String() string {
	var bfr bytes.Buffer
	bfr.WriteString("{")
	for i, c := range t.Children {
		if c != nil {
			fmt.Fprintf(&bfr, "%c", 'a'+i)
		} else {
			fmt.Fprintf(&bfr, "-")
		}
	}
	if t.IsWord {
		bfr.WriteString(" *}")
		return bfr.String()
	}
	bfr.WriteString(" /}")
	return bfr.String()
}

func (t *Trie26) Insert(word string) {
	for i := 0; i < len(word); i++ {
		c := t.Children[word[i]-'a']
		if c == nil {
			c = &Trie26{}
			t.Children[word[i]-'a'] = c
		}
		t = c
	}
	t.IsWord = true
}

func (t *Trie26) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		c := t.Children[word[i]-'a']
		if c == nil {
			return false
		}
		t = c
	}
	log.Print(t)
	return t.IsWord
}

// 648m Replace Words
func replaceWords(dictionary []string, sentence string) []string {
	type Trie struct {
		IsWord bool
		Child  [26]*Trie
	}

	Insert := func(r *Trie, w string) {
		for i := 0; i < len(w); i++ {
			c := r.Child[w[i]-'a']
			if c == nil {
				c = &Trie{}
				r.Child[w[i]-'a'] = c
			}
			r = c
		}
		r.IsWord = true
	}

	Search := func(r *Trie, w string) bool {
		for i := 0; i < len(w); i++ {
			c := r.Child[w[i]-'a']
			if c == nil {
				return false
			}
			r = c
		}
		return r.IsWord
	}

	trie := &Trie{}
	for _, w := range dictionary {
		Insert(trie, w)
	}

	R := strings.Split(sentence, " ")
	for i, w := range R {
		for l := 1; l < len(w); l++ {
			if Search(trie, w[:l]) {
				R[i] = w[:l]
				break
			}
		}
	}
	return R
}
