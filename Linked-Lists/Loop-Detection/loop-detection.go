package Loop_Detection

import Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"

type singlyLinkedList Linked_Lists.SinglyLinkedList

// LoopDetection implements an algorithm that returns the node at the beginning of the loop.
func (sll singlyLinkedList) LoopDetection() Linked_Lists.Node {
	parsedNodes := make(map[Linked_Lists.Node]bool)

	curr := sll.Head

	for curr != nil {
		if _, ok := parsedNodes[*curr]; !ok {
			parsedNodes[*curr] = true
		} else {
			return *curr
		}
		curr = curr.NextNode
	}
	// In case that there was no loop in Linked List
	// We simply return empty Node
	return Linked_Lists.Node{}
}
