/*
Palindrome Permutation:
Given a string, write a function to check if it is a permutation of a palindrome.
A palindrome is a word or phrase that is the same forwards and backwards.
A permutation is a rearrangement of letters.
The palindrome does not need to be limited to just dictionary words.
EXAMPLE
Input: Tact Coa
Output: True (permutations: "taco cat". "atco cta". etc.)
*/
package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Tact Coatt")
	fmt.Println(palindromePermutation("Tact Coatt"))
	fmt.Println(palindromePermutation("A man, a plan, a canal: Panama"))
}

func palindromePermutation(palindrome string) bool{
	lowecase := strings.ToLower(palindrome)
	letters := make(map[rune]int)
	for _, letter := range lowecase{
		if letter == ' '{
			continue
		}
		letters[letter]++
	}
	oddToken := true
	for ind := range letters{
		if letters[ind] % 2 == 0 {
			continue
		} else if oddToken{
			oddToken = !oddToken
		} else {
			return false
		}
	}
	return true
}