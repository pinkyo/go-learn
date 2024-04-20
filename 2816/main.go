package main

func main() {
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	head = reverse(head)
	cur := head
	var pre *ListNode
	incr := 0
	for cur != nil {
		temp := 2*cur.Val + incr
		incr = temp / 10
		cur.Val = temp % 10
		pre = cur
		cur = cur.Next
	}
	if incr > 0 {
		pre.Next = &ListNode{Val: incr}
	}
	head = reverse(head)
	return head
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
