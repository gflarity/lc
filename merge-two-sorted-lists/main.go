package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->3->5->7
// 2->4->6->8
// _
func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	// edge case, one is nil
	if a == nil && b == nil {
		return nil
	} else if a == nil {
		return b
	} else if b == nil {
		return a
	}

	// initialize cur to smallest of the tw
	var head *ListNode
	if a.Val < b.Val {
		head = a
		a = a.Next
	} else {
		head = b
		b = b.Next
	}

	// use this to construct the list coming from head
	cur := head

	for a != nil || b != nil {
		// check the for one of them being nil, this ends the merge as we just
		// continue with the other
		if a == nil {
			cur.Next = b
			cur = cur.Next
			break
		} else if b == nil {
			cur.Next = a
			cur = cur.Next
			break
		}
		// neither are nil if we've gotten here

		// at every step you'll have two choices, always pick the smaller
		if a.Val < b.Val {
			// pick a
			cur.Next = a
			a = a.Next
		} else {
			// pick b
			cur.Next = b
			b = b.Next
		}

		// NEXT!
		cur = cur.Next
	}
	return head
}

func printWalk(n *ListNode) {
	for n != nil {
		fmt.Printf("%d->", n.Val)
		n = n.Next
	}
	fmt.Print("\n")
}

func main() {
	a1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	b1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
	printWalk(mergeTwoLists(a1, b1))

	// edge case one lis tis nill
	a2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	var b2 *ListNode
	printWalk(mergeTwoLists(a2, b2))

	// edge case different lengths
	a3 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	b3 := &ListNode{Val: 2}
	printWalk(mergeTwoLists(a3, b3))

}
