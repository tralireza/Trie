package Trie

import (
	"bytes"
	"fmt"
	"io"
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
		bfr.WriteString(" /}")
		return bfr.String()
	}
	bfr.WriteString(" *}")
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
	return t.IsWord
}

func (t *Trie26) Graphviz(wr io.Writer) {
	io.Copy(wr, strings.NewReader("digraph {\nnode [shape=rect];\nedge [arrowhead=vee];\n"))

	var Walk func(n, p *Trie26, w []byte)
	Walk = func(n, p *Trie26, w []byte) {
		if n.IsWord {
			io.Copy(wr, strings.NewReader(fmt.Sprintf("%q [color=red shape=note];\n", w)))
		}
		for i, c := range n.Children {
			if c != nil {
				io.Copy(wr, strings.NewReader(fmt.Sprintf("%q -> %q\n", w, append(w, byte(i)+'a'))))
				Walk(c, n, append(w, byte(i)+'a'))
			}
		}
	}
	Walk(t, nil, []byte{})

	wr.Write([]byte{'}'})
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
