package Linked_Lists

import Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"

type singlyLinkedList Linked_Lists.SinglyLinkedList

// I suppose that to in that case means from end

func (sll singlyLinkedList) KthToLasr(k int) int {
	p1 := sll.Head
	p2 := sll.Head

	// Traverse p1 on k elems
	for i := 0; i < k; i++ {
		if p1 == nil {
			return 0 // Out of bounds
		}
		p1 = p1.NextNode
	}
	// Move them at the same pace.
	// When p1 hits the end, p2 will be at the right element.
	for p1 != nil {
		p1 = p1.NextNode
		p2 = p2.NextNode
	}
	return p2.Data
}
