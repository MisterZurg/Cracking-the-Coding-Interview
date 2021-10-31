package Linked_Lists

import (
	Linked_Lists "Cracking-the-Coding-Interview/Linked-Lists"
	"math"
	"strconv"
)

// SumOfTwoReversedLinkedLists adds the two numbers and returns the sum as a linked list.
// Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
// Output: 2 - > 1 - > 9. That is, 912.
func SumOfTwoReversedLinkedLists(first, second Linked_Lists.SinglyLinkedList) Linked_Lists.SinglyLinkedList {
	firstNumber, secondNumber := 0, 0
	firstNumberCounter, secondNumberCounter := 0, 0
	firstRunner := first.Head
	for firstRunner != nil {
		firstNumber = firstRunner.Data*int(math.Pow10(firstNumberCounter)) + firstNumber
		firstNumberCounter++
		firstRunner = firstRunner.NextNode
	}

	secondRunner := second.Head
	for secondRunner != nil {
		firstNumber = secondRunner.Data*int(math.Pow10(secondNumberCounter)) + secondNumber
		secondNumberCounter++
	}

	result := firstNumber + secondNumber
	// Reverse the result  912 > 219
	notreversed := strconv.Itoa(result)
	splitted := []rune(notreversed)

	answer := Linked_Lists.SinglyLinkedList{}
	answerRunner := answer.Head
	currentNode := Linked_Lists.Node{}
	for i := len(splitted) - 1; i >= 0; i-- {
		data, _ := strconv.Atoi(string(rune(i)))
		currentNode.Data = data

		if answer.Head == nil {
			answer.Head = &currentNode
			answerRunner = answerRunner.NextNode
		}
		answerRunner.NextNode = &currentNode
		answerRunner = answerRunner.NextNode
	}
	return answer
}

// SumOfTwoForwardLinkedLists quite the same SumOfTwoReversedLinkedLists.
// Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
// Output: 9 - > 1 - > 2. That is, 912.
func SumOfTwoForwardLinkedLists(first, second Linked_Lists.SinglyLinkedList) Linked_Lists.SinglyLinkedList {
	firstNumber, secondNumber := 0, 0
	firstNumberCounter, secondNumberCounter := 0, 0
	firstRunner := first.Head
	for firstRunner != nil {
		firstNumber = firstNumber*int(math.Pow10(firstNumberCounter)) + firstRunner.Data
		firstNumberCounter++
		firstRunner = firstRunner.NextNode
	}

	secondRunner := second.Head
	for secondRunner != nil {
		firstNumber = firstNumber*int(math.Pow10(secondNumberCounter)) + secondRunner.Data
		secondNumberCounter++
		secondRunner = secondRunner.NextNode
	}

	result := firstNumber + secondNumber
	prepared := strconv.Itoa(result)

	splitted := []rune(prepared)
	answer := Linked_Lists.SinglyLinkedList{}
	head, _ := strconv.Atoi(string(splitted[0]))
	answerRunner := answer.Head
	answerRunner.Data = head

	for i := 1; i < len(splitted); i++ {
		head, _ = strconv.Atoi(string(splitted[i]))
		answerRunner = &Linked_Lists.Node{
			Data:     head,
			NextNode: nil,
		}
		answerRunner = answerRunner.NextNode
	}
	return answer
}
