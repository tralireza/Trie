package Trie

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strings"
	"testing"
)

func init() {
	log.Print("> Trie")
}

func TestTrie26(t *testing.T) {
	trie := &Trie26{}
	for _, w := range []string{"at", "art", "zoo", "zebra", "zeta", "lion", "liquor", "bat", "battle", "batman", "batsman", "battleship", "let", "letter", "lettings"} {
		trie.Insert(w)
	}
	for _, w := range []string{"fox", "bat"} {
		log.Printf("? %q -> %t", w, trie.Search(w))
	}

	var Graph func(n, p *Trie26, word []byte)
	Graph = func(n, p *Trie26, word []byte) {
		if n.IsWord {
			log.Printf("%q", word)
		}
		for i, c := range n.Children {
			if c != nil {
				log.Printf("%v %v %q | %q -> %q", n, c, word, word, append(word, byte(i)+'a'))
				Graph(c, n, append(word, byte(i)+'a'))
			}
		}
	}
	Graph(trie, nil, []byte{})

	f, err := os.OpenFile("trie26.gv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, fs.FileMode(0640))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	trie.Graphviz(w)
	w.Flush()
}

// 648m Replace Words
func Test648(t *testing.T) {
	HashMap := func(dictionary []string, sentence string) []string {
		D := map[string]struct{}{}
		for _, r := range dictionary {
			D[r] = struct{}{}
		}

		R := strings.Split(sentence, " ")
		for i, w := range R {
			for l := 1; l < len(w); l++ {
				if _, ok := D[w[:l]]; ok {
					R[i] = w[:l]
					break
				}
			}
		}
		return R
	}

	for _, f := range []func([]string, string) []string{HashMap, replaceWords} {
		log.Print(" -> ", f([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
		log.Print(" -> ", f([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
	}
}
