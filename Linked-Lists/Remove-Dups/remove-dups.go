package Linked_Lists

import Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"

type singlyLinkedList Linked_Lists.SinglyLinkedList

// RemoveDupsWithBuffer
func (sll singlyLinkedList) RemoveDupsWithBuffer() {
	currNode := sll.Head
	chachedElems := make(map[int]bool)
	for currNode.NextNode != nil {
		if _, ok := chachedElems[currNode.Data]; ok {
			currNode.NextNode = currNode.NextNode.NextNode
		}
		chachedElems[currNode.Data] = true
		currNode = currNode.NextNode
	}
	return
}

// RemoveDupsNoBuffer No Buffer Allowed then two pointers technique
func (sll singlyLinkedList) RemoveDupsNoBuffer() {
	current := sll.Head
	for current != nil {
		runner := current
		for runner.NextNode != nil {
			if current.Data == runner.NextNode.Data {
				runner.NextNode = runner.NextNode.NextNode
			}
			runner = runner.NextNode
		}
		current = current.NextNode
	}
}
