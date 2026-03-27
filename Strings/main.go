package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseWords(s string) string {
	// strings.Split will split according to the provided separator but strings.Fields splits it with any whitespace
	words := strings.Fields(s)
	fmt.Println(words)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func isPalindrome(s string) bool {
	var filtered []rune
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}
	fmt.Println("filtered", filtered)
	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	maxLength, start := 0, 0
	for i, r := range s {
		if idx, found := charIndex[r]; found && idx >= start {
			start = idx + 1
		}
		charIndex[r] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}

func main() {
	// fmt.Println(reverseWords("Go is fun"))                      // Output: "fun is Go"
	// fmt.Println(isPalindrome("A man, a plan, a canal, Panama")) // Output: true
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))           // Output: 3
	symbol := map[string]string{
		"USD": "1",
		"EUR": "2",
		"GBP": "3",
		"RMB": "4",
	}
	// fmt.Print("symbol", symbol)
	// fmt.Println("RMB", symbol["RMB"])

	for i, v := range symbol {
		fmt.Printf("Key: %s, Value: %s\n", i, v)
	}

	r := [...]int{99: -1}

	for i, v := range r {
		fmt.Printf("Key: %d, Value: %d\n", i, v)
	}
}
