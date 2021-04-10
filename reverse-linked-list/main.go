package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// [1,2,3,4,5]: 1->2->3->4->5
// 1<-2<-3<-4<-5
// 1, 2->1, 3->2->1
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	newHead := copy(cur)
	for cur.Next != nil {
		// copy the next node to be inserted
		next := copy(cur.Next)
		// new node's next is the old newHead
		next.Next = newHead
		// newHead is now the next
		newHead = next
		// move forward
		cur = cur.Next
	}

	return newHead
}

func copy(node *ListNode) *ListNode {
	return &ListNode{
		Val: node.Val,
	}
}

