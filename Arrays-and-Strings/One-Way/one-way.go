/*
One Away:
There are three types of edits that can be performed on strings:
	insert a character,
	remove a character,
	or replace a character.

Given two strings, write a function to check if they are one edit (or zero edits) away.
 */
package One_Way

import "math"

func oneway(first, second string) bool{
	// Length checks.
	if math.Abs(float64(len(first)- len(second))) > 1 {
		return false
	}

	if len(first) == len(second) {
		return oneEditReplace(first, second)
	}

	if len(first) - 1 == len(second) {
		return oneEditInsert(second, first)
	}

	if len(first) + 1 == len(second) {
		return oneEditInsert(first, second)
	}

	return false
}

func oneEditReplace(first, second string) bool {
	foundDifferense := false
	for i:=0;i< len(first);i++{
		if first[i] != second[i]{
			if foundDifferense {
				return false
			}
			foundDifferense = true
		}
	}
	return true
}

func oneEditInsert(first, second string) bool {
	i, j := 0, 0
	for i < len(first) && j < len(second){
		if first[i] != second[j]{
			if i != j {
				return false
			}
			j++
		} else {
			i++
			j++
		}
	}
	return true
}