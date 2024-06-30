package main

func main() {
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// This function inserts the greatest common divisor between the current node's value and the next node's value
// into the list as a new node. It then continues to do this for the rest of the list until it reaches the end.
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	// Initialize a pointer to the head of the list
	cur := head

	// Loop through the list until the end is reached
	for cur != nil && cur.Next != nil {
		// Calculate the greatest common divisor between the current node's value and the next node's value
		divisor := greatestCommonDivisor(cur.Val, cur.Next.Val)

		// Save the next node
		next := cur.Next

		// Insert a new node with the greatest common divisor value into the list
		cur.Next = &ListNode{Val: divisor, Next: next}

		// Move the pointer to the next node
		cur = next
	}

	// Return the head of the list
	return head
}

// This function calculates the greatest common divisor between two integers using the Euclidean algorithm.
func greatestCommonDivisor(a, b int) int {
	// If the second integer is greater than the first, swap them
	if a < b {
		a, b = b, a
	}

	// If the second integer is 0, return the first integer
	if b == 0 {
		return a
	}

	// Recursively call the function with the second and remainder of the first divided by the second
	return greatestCommonDivisor(b, a%b)
}
