package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestCorpus", func() {
	var corpus Corpus

	Describe("FindAWord", func() {
		var prefix string

		Context("When using version 1", func() {
			BeforeEach(func() {
				corpus = NewCorpusV1()
			})

			Context("When prefix is matching", func() {
				BeforeEach(func() {
					prefix = "app"
				})

				It("returns a word that matches", func() {
					word := corpus.FindAWord(prefix)
					Expect(word).To(HavePrefix(prefix))
				})
			})

			Context("When prefix is blank", func() {
				BeforeEach(func() {
					prefix = ""
				})

				It("returns a word", func() {
					word := corpus.FindAWord(prefix)
					Expect(word).NotTo(BeEmpty())
				})
			})

			Context("When prefix is not matching", func() {
				BeforeEach(func() {
					prefix = "xxxx"
				})

				It("returns no word", func() {
					word := corpus.FindAWord(prefix)
					Expect(word).To(BeEmpty())
				})
			})

		})

		Context("When using corpus2", func() {
			BeforeEach(func() {
				corpus = NewCorpusV2()
			})

			Context("When prefix is matching", func() {
				BeforeEach(func() {
					prefix = "app"
				})

				It("returns a word that matches", func() {
					word := corpus.FindAWord(prefix)
					Expect(word).To(HavePrefix(prefix))
				})
			})

			Context("When prefix is blank", func() {
				BeforeEach(func() {
					prefix = ""
				})

				It("returns a word", func() {
					word := corpus.FindAWord(prefix)
					Expect(word).NotTo(BeEmpty())
				})
			})

			Context("When prefix is not matching", func() {
				BeforeEach(func() {
					prefix = "xxxx"
				})

				It("returns no word", func() {
					word := corpus.FindAWord(prefix)
					Expect(word).To(BeEmpty())
				})
			})

		})
	})

	Describe("IsAWord", func() {
		var word string

		Context("When word exists", func() {
			BeforeEach(func() {
				corpus = GetCorpus()
				word = "apple"
			})

			It("returns true", func() {
				Expect(corpus.IsAWord(word)).To(BeTrue())
			})
		})

		Context("When using version 2", func() {
			BeforeEach(func() {
				corpus2 := CorpusV2{trie: map[string]interface{}{}}
				corpus2.BuildTrie("apple")
				corpus2.BuildTrie("appletree")
				corpus = &corpus2
			})

			It("returns true", func() {
				Expect(corpus.IsAWord(word)).To(BeTrue())
			})
		})

		Context("When word not exists", func() {
			Context("When using version 2", func() {
				BeforeEach(func() {
					corpus2 := CorpusV2{trie: map[string]interface{}{}}
					corpus2.BuildTrie("apple")
					corpus = &corpus2
					word = "peanuts"
				})

				It("returns false", func() {
					Expect(corpus.IsAWord(word)).To(BeFalse())
				})
			})
		})
	})

	Describe("BuildTrie", func() {
		var expectedTrie map[string]interface{}
		var corpus2 CorpusV2

		BeforeEach(func() {
			corpus2 = CorpusV2{trie: map[string]interface{}{}}
			expectedTrie = map[string]interface{}{
				"a": map[byte]string{
					'b': "ab",
					'p': "ap",
				},
				"ab": map[byte]string{
					'l': "abl",
				},
				"abl": map[byte]string{
					'e': "able",
				},
				"able": map[byte]string{
					0: "able",
				},
				"ap": map[byte]string{
					'p': "app",
				},
				"app": map[byte]string{
					'l': "appl",
				},
				"appl": map[byte]string{
					'e': "apple",
				},
				"apple": map[byte]string{
					0:   "apple",
					't': "applet",
				},
				"applet": map[byte]string{
					'r': "appletr",
				},
				"appletr": map[byte]string{
					'e': "appletre",
				},
				"appletre": map[byte]string{
					'e': "appletree",
				},
				"appletree": map[byte]string{
					0: "appletree",
				},
			}
		})

		It("builds a trie", func() {
			corpus2.BuildTrie("able")
			corpus2.BuildTrie("apple")
			corpus2.BuildTrie("appletree")
			Expect(corpus2.trie).To(BeEquivalentTo(expectedTrie))
		})
	})
})
