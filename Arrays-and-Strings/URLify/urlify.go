/*
URLify:
Write a method to replace all spaces in a string with '%2e:
You may assume that the string has sufficient space at the end to hold the additional characters,
and that you are given the "true" length of the string.
(Note: if implementing in Java, please use a character array so that you can perform this operation in place.)
EXAMPLE
Input: "Mr John Smith JJ, 13
Output: "Mr%20John%20Smith"
*/
package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("\"Mr John Smith     \", 13")
	fmt.Println(URLify("\"Mr John Smith     \", 13"))
}

func URLify(notUrl string) string {
	notUrlSplit := strings.Split(notUrl,",")
	propperPart := notUrlSplit[0]
	almostURL := strings.TrimSpace(propperPart[:len(propperPart)-1])+"\""
	return strings.ReplaceAll(almostURL," ", "%20")
}