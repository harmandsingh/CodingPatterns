// The first input of the test case is an array of values representing a linked list.
// The second input is the index where the tail connects to form a cycle (or −1 if there's no cycle).
// This index is used only to construct the linked list and is not passed to the function.

package main

// Definition for a Linked List node
type ListNode struct {
	Val  int
	Next *ListNode
}

func countCycleLength(head *ListNode) int {

	slow, fast := head, head
	cycleLength := 0
	intersectionNode := &ListNode{}

	// Break when we've reached the last node and cycle does not exist
	// If a cycle is found we won't be reaching the last node and we'll break from inside the for loop
	for {
		// Move slow pointer
		if slow != nil && slow.Next != nil {
			slow = slow.Next
		} else {
			return 0
		}

		// Move fast pointer twice
		if fast != nil && fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return 0
		}

		// Check if fast pointer and slow pointer point to the same node and thus a cycle exists
		if slow == fast {
			intersectionNode = slow
			break
		}
	}

	// Find cycle length
	tempNode := intersectionNode
	for {
		tempNode = tempNode.Next
		cycleLength++
		if tempNode == intersectionNode {
			break
		}
	}

	return cycleLength
}
