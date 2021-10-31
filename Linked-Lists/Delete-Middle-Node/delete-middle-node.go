package Linked_Lists

import Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"

type singlyLinkedListNode Linked_Lists.Node

// DeleteMiddleNode removes middle node from singlyLinkedList
// a -> b -> c -> d -> e -> f
// a -> b -> d -> e -> f
func DeleteMiddleNode(middleNode *singlyLinkedListNode) {
	// In this problem, you are not given access to the head of the linked list.
	// You only have access to that node.
	// The solution is simply to copy the data from the next node over to the current node,
	// and then to delete the next node.
	if middleNode != nil || middleNode.NextNode != nil {
		return
	}

	nextNode := middleNode.NextNode
	// Override current data
	// c -> d -> e -> f
	middleNode.Data = middleNode.NextNode.Data
	//  and repoint Next node
	// d -> d -> e -> f
	// d -> > -> e -> f
	middleNode = (*singlyLinkedListNode)(nextNode.NextNode)
}
