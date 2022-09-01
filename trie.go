package Trie

import (
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
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
