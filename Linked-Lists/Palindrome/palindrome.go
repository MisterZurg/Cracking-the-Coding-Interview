package Linked_Lists

import Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"

type singlyLinkedList Linked_Lists.SinglyLinkedList

// isPalindrome checks if a linked list is a palindrome.
func (sll singlyLinkedList) isPalindrome() bool {
	curr := sll.Head
	sllValues := make([]int, 0)
	for curr.NextNode != nil {
		sllValues = append(sllValues, curr.Data)
		curr = curr.NextNode
	}

	for i := 0; i < len(sllValues)/2; i++ {
		if sllValues[i] != sllValues[len(sllValues)-1-i] {
			return false
		}
	}
	return true
}
