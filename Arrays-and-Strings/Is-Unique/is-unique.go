/*
Is Unique:
Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/
package main

import (
	"fmt"
	"sort"
)

func main()  {
	unique := "abcde"
	fmt.Println(isUnique(unique))
	notUnique := "abcderfghtqa"
	fmt.Println(isUnique(notUnique))
	notMap := "cdesatba"
	fmt.Println(isUniqueNoMap(notMap))
}

// Case 1 : We allowed to use map:
func isUnique(message string) bool {
	alphabet := make(map[rune]bool)
	for _, letter := range message {
		if !alphabet[letter]{
			alphabet[letter] = true
		} else {
			return false
		}
	}
	return true
}

// Case 2 : We are not allowed to use map
func isUniqueNoMap(message string) bool {
	messageSliceTokens := []rune(message)
	sort.Slice(messageSliceTokens, func(i, j int) bool {
		return messageSliceTokens[i] < messageSliceTokens[j]
	})
	for i := 1;i< len(messageSliceTokens); i++{
		if messageSliceTokens[i - 1] == messageSliceTokens[i]{
			return false
		}
	}
	return true
}