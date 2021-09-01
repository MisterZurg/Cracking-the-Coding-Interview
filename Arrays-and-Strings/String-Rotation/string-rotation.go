/*
String Rotation:
Assume you have a method isSubstring which checks if one word is a substring of another.
Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one
call to isSubstring (e.g., "waterbottle" is a rotation of "erbottlewat").
*/
package String_Rotation

import "strings"

func IsRotation(s1, s2 string) bool{
	// Search for hot point like when s1[0] = s2[Z]
	// Split string and compare parts
	checkString := s2 + s2
	return strings.Contains(checkString, s1)
}
