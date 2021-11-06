package Linked_Lists

import Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"

type singlyLinkedList Linked_Lists.SinglyLinkedList

func Intersection(first, second singlyLinkedList) Linked_Lists.Node {
	firstNodes := make([]*Linked_Lists.Node, 0)
	secondNodes := make([]*Linked_Lists.Node, 0)

	curr := first.Head
	for curr.NextNode != nil {
		firstNodes = append(firstNodes, curr)
		curr = curr.NextNode
	}

	curr = second.Head
	for curr.NextNode != nil {
		secondNodes = append(secondNodes, curr)
		curr = curr.NextNode
	}
	var sharedNode *Linked_Lists.Node = &Linked_Lists.Node{}
	for i, j := len(firstNodes)-1, len(secondNodes)-1; i > 0 && j > 0 && firstNodes[i] == firstNodes[j]; {
		sharedNode = firstNodes[i]
		i--
		j--
	}

	return *sharedNode
}
