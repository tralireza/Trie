package Trie

import (
	"log"
	"strings"
	"testing"
)

func init() {
	log.Print("> Trie")
}

func TestTrie26(t *testing.T) {
	trie := &Trie26{}
	for _, w := range []string{"cat", "rat", "bat", "battle"} {
		trie.Insert(w)
	}
	for _, w := range []string{"fox", "bat"} {
		log.Print(" -> ", trie.Search(w))
	}
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
