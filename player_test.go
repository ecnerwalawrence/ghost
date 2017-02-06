package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestPlayer", func() {
	var player Player

	// Testing applies to Computer Player only
	Describe("Play", func() {

		// Testing for Human Player would require some mocking
		// Given the time limit, I did not put the effort.

		Context("When player is a computer", func() {
			var computer Computer
			var oldTargetWord, prefix string
			var letter byte

			Context("When prefix has changed from computer's targetword", func() {
				BeforeEach(func() {
					oldTargetWord = "again"
					computer = Computer{targetWord: oldTargetWord}
					player = &computer
					prefix = "ap"
					letter = player.Play(prefix)
				})

				It("randomly picks a new target word", func() {
					Expect(computer.targetWord).NotTo(Equal(oldTargetWord))
				})

				It("randomly picks a new target word", func() {
					Expect(computer.targetWord).To(ContainSubstring(string(letter)))
				})
			})
		})
	})
})
