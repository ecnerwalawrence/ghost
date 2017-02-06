package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player interface {
	Play(prefix string) byte
	PlayChallenge(prefix string) bool
	GetType() string
}

type Human struct{}

func (h *Human) GetType() string { return "Human" }

func (h *Human) Play(prefix string) byte {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("[Enter Letter]")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "challenge" && len(prefix) > 0 {
			fmt.Println("[Human will challenge]")
			return 0
		} else if len(input) != 1 {
			fmt.Printf("[Invalid Input:%s]\n", input)
			continue
		}
		return input[0]
	}
}

func (h *Human) PlayChallenge(prefix string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("[Enter Word]")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "bluff" || input == "" || !strings.HasPrefix(input, prefix) {
		return false
	}
	fmt.Printf("[word:%s]\n", input)
	return true
}

type Computer struct {
	targetWord string
	bluff      bool
}

func (c *Computer) GetType() string { return "Computer" }

func (c *Computer) Play(prefix string) byte {
	corpus := GetCorpus()
	if !strings.HasPrefix(c.targetWord, prefix) || prefix == "" {
		c.targetWord = corpus.FindAWord(prefix)
	}

	fmt.Printf("{computer's target word: `%s`}\n", c.targetWord)
	if c.targetWord == "" {
		// Computer can either challenge or bluff with a fake letter
		fmt.Println("computer will challenge")
		return 0
	} else {
		c.bluff = false
	}

	return c.targetWord[len(prefix)]
}

func (c *Computer) PlayChallenge(prefix string) bool {
	return c.bluff
}
