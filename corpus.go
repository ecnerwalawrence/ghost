package main

import (
	"math/rand"
	"os"
	"strings"
)

type Corpus interface {
	FindAWord(prefix string) string
	IsAWord(prefix string) bool
}

type CorpusV1 struct {
	Dict map[string]string
}

var corpusSingleton Corpus

func GetCorpus() Corpus {
	if corpusSingleton != nil {
		return corpusSingleton
	}

	if os.Getenv("CORPUS") == "1" {
		corpusSingleton = NewCorpusV1()
	} else {
		corpusSingleton = NewCorpusV2()
	}

	return corpusSingleton
}

func NewCorpusV1() *CorpusV1 {
	words := []string{"apple", "appletree", "word", "work"}

	dict := map[string]string{}
	for _, w := range words {
		dict[w] = w
	}
	return &CorpusV1{Dict: dict}
}

func (c CorpusV1) FindAWord(prefix string) string {
	if prefix == "" {
		i := rand.Intn(len(c.Dict))
		for _, w := range c.Dict {
			if i != 0 {
				i--
			} else {
				return w
			}
		}
	}

	for _, w := range c.Dict {
		if strings.HasPrefix(w, prefix) {
			return w
		}
	}
	return ""
}

func (c CorpusV1) IsAWord(prefix string) bool {
	return (c.Dict[prefix] != "")
}

type CorpusV2 struct {
	trie  map[string]interface{}
	words []string
}

func NewCorpusV2() *CorpusV2 {
	corpus := CorpusV2{trie: map[string]interface{}{}, words: []string{}}
	words := []string{"apple", "appletree", "word", "work"}
	for _, w := range words {
		corpus.BuildTrie(w)
	}
	return &corpus
}

// Assume words are sorted
func (c *CorpusV2) BuildTrie(word string) {
	for i := 0; i < (len(word) - 1); i++ {
		prefixKey := string(word[0 : i+1])
		prefixVal := string(word[0:(i + 2)])
		nextLetter := word[i+1]

		m := c.trie[prefixKey]
		if m == nil {
			c.trie[prefixKey] = map[byte]string{nextLetter: prefixVal}
		} else if mLvl2, ok := m.(map[byte]string); ok {
			mLvl2[nextLetter] = prefixVal
		}
	}
	if c.trie[word] == nil {
		c.trie[word] = map[byte]string{0: word}
	} else if mLvl2, ok := c.trie[word].(map[byte]string); ok {
		mLvl2[0] = word
	}
	c.words = append(c.words, word)
}

// Recurses through the Trie
func (c CorpusV2) FindAWord(prefix string) string {
	if prefix == "" {
		index := rand.Intn(len(c.words) - 1)
		return c.words[index]
	}

	w := c.trie[prefix]
	if w == nil {
		return ""
	}

	var mLvl2 map[byte]string
	var ok bool
	if mLvl2, ok = w.(map[byte]string); !ok {
		return ""
	}

	i := rand.Intn(len(mLvl2))
	for k, v := range mLvl2 {
		if i != 0 {
			i--
			continue
		}
		if k == 0 {
			return v
		}
		return c.FindAWord(v)
	}

	return ""
}

func (c CorpusV2) IsAWord(prefix string) bool {
	w := c.trie[prefix]

	if w == nil {
		return false
	}

	_, word := c.getByteMap(w)
	if word == "" {
		return false
	}

	return true
}

func (c CorpusV2) getByteMap(v interface{}) (map[byte]string, string) {
	var mLvl2 map[byte]string
	var ok bool
	if mLvl2, ok = v.(map[byte]string); !ok {
		return mLvl2, ""
	}

	if mLvl2[0] == "" {
		return mLvl2, ""
	}

	return mLvl2, mLvl2[0]
}
