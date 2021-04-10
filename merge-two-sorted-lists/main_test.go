package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func printWalk(n *ListNode) {
	for n != nil {
		fmt.Printf("%d->", n.Val)
		n = n.Next
	}
	fmt.Print("\n")
}

func TestMergeTwoLists(t *testing.T) {

	// 1->2->3->4->5->6->
	a1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	b1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
	//1->2->3->4->5->6->
	m1 := mergeTwoLists(a1, b1)
	assert.NotNil(t,  m1)
	assert.Equal(t, 1, m1.Val)

	assert.NotNil(t,  m1.Next)
	assert.Equal(t, 2, m1.Next.Val)

	assert.NotNil(t,  m1.Next.Next)
	assert.Equal(t, 3, m1.Next.Next.Val)

	assert.NotNil(t, m1.Next.Next.Next)
	assert.Equal(t, 4, m1.Next.Next.Next.Val)

	assert.NotNil(t, m1.Next.Next.Next.Next)
	assert.Equal(t, 5, m1.Next.Next.Next.Next.Val)

	assert.NotNil(t, m1.Next.Next.Next.Next.Next)
	assert.Equal(t, 6, m1.Next.Next.Next.Next.Next.Val)

	// edge case one lis tis nil
	a2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	var b2 *ListNode
	// 1->3->5->
	m2 := mergeTwoLists(a2, b2)
	assert.NotNil(t,  m2)
	assert.Equal(t, 1, m2.Val)

	assert.NotNil(t,  m2.Next)
	assert.Equal(t, 3, m2.Next.Val)

	assert.NotNil(t,  m2.Next.Next)
	assert.Equal(t, 5, m2.Next.Next.Val)


	// edge case different lengths
	a3 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	b3 := &ListNode{Val: 2}
	m3 := mergeTwoLists(a3, b3)
	assert.NotNil(t,  m3)
	assert.Equal(t, 1, m3.Val)

	assert.NotNil(t,  m2.Next)
	assert.Equal(t, 2, m3.Next.Val)

	assert.NotNil(t,  m2.Next.Next)
	assert.Equal(t, 3, m3.Next.Next.Val)

}



