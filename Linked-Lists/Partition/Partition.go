package Linked_Lists

import Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"

type singlyLinkedList Linked_Lists.SinglyLinkedList

func (sll singlyLinkedList) Partition(x int) Linked_Lists.Node {
	lessThanXSequence := Linked_Lists.Node{}
	greaterThanXSequence := Linked_Lists.Node{}

	mainNodeRunner := sll.Head

	for mainNodeRunner != nil {
		if mainNodeRunner.Data < x {
			// Insert node than less
			mainNodeRunner.NextNode = &lessThanXSequence
			lessThanXSequence = *mainNodeRunner
			// before all nodes greater than or equal to x.
		} else {
			greaterThanXSequence.NextNode = mainNodeRunner
			greaterThanXSequence = *mainNodeRunner
		}
		mainNodeRunner = mainNodeRunner.NextNode
	}
	greaterThanXSequence.NextNode = nil
	return lessThanXSequence
}
