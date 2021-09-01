/*
Check Permutation:
Given two strings, write a method to decide if one is a permutation of the other.
 */
package main

import "fmt"

func main()  {
	word := "cat "
	permutation := " tAc"
	fmt.Println(isPermutation(word,permutation))
	// fmt.Println('a', 'z')
	// fmt.Println('A', 'Z')
}
/*
I will assume for this problem that the comparison is case sensitive
and whitespace is significant. So, "god " is different from "dog".
 */
func isPermutation(str, perm string) bool {
	if len(str) != len(perm) {
		return false
	}

    countLetters := make([]int, 128)
    for _, letter := range str {
    	countLetters[letter]++
	}

	for _, letter := range perm {
		countLetters[letter]--
		if countLetters[letter] < 0 {
			return false
		}
	}
	return true
}