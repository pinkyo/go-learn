package main

func main() {
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		divisor := greatestCommonDivisor(cur.Val, cur.Next.Val)
		next := cur.Next
		cur.Next = &ListNode{Val: divisor, Next: next}
		cur = next
	}
	return head
}

func greatestCommonDivisor(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return greatestCommonDivisor(b, a%b)
}
