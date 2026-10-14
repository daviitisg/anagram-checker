package main

import (
	"fmt"
	"strings"
)

func main() {
	var word1, word2 string

	fmt.Print("First word: ")
	fmt.Scan(&word1)
	fmt.Print("Second word: ")
	fmt.Scan(&word2)

	// Make all letters lowercase to ensure case insensitivity
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	map1 := count(word1)
	map2 := count(word2)
	
	if compare(map1, map2) {
		fmt.Println("These words are anagrams.")
	} else {
		fmt.Println("These words are NOT anagrams.")
	}
}

func count(lett string) map[rune]int {
	counts := make(map[rune]int)

	for _, i := range lett {
		counts[i]++
	}
	return counts
}

func compare(m1, m2 map[rune]int) bool {
	// If the number of unique characters differs, they cannot be anagrams
	if len(m1) != len(m2) {
		return false
	}

	// Check if m2 contains the same characters with the same frequencies
	for char, count := range m1 {
		if m2[char] != count {
			return false
		}
	}
	return true
}