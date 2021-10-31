package Linked_Lists

type Node struct {
	Data     int
	NextNode *Node
}

func newNode(data int) *Node {
	return &Node{Data: data}
}

// SinglyLinkedList implements a Data structure that represents a sequence of nodes.
// It wraps Node struct, pointing on the first Node in sequence
// The benefit of a linked list is that you can add
// and remove items from the beginning of the list in constant time.
// For specific applications, this can be useful.
type SinglyLinkedList struct {
	Head *Node
}

func (sll SinglyLinkedList) appendToTail(data int) {
	currentNode := sll.Head
	// Traverse from Head till Tail
	for currentNode != nil {
		currentNode = currentNode.NextNode
	}

	currentNode.NextNode = newNode(data)
}

// Deleting a Node from a Singly Linked List
func (sll SinglyLinkedList) deleteNode(data int) {
	// moved Head
	if sll.Head.Data == data {
		sll.Head = sll.Head.NextNode
	}

	currentNode := sll.Head
	for currentNode.NextNode != nil {
		if currentNode.NextNode.Data == data {
			currentNode.NextNode = currentNode.NextNode.NextNode
			return
		}
		currentNode = currentNode.NextNode
	}
	return
}
