package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	r := reverseList(nil)
	assert.Nil(t, r)

	list := &ListNode{Val: 1}
	r = reverseList(list)
	assert.NotNil(t, r)
	assert.Equal(t, 1, r.Val)

	list = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	r = reverseList(list)
	assert.Equal(t, 3, r.Val)
	assert.NotNil(t, r.Next)
	assert.Equal(t, 2, r.Next.Val)
	assert.NotNil(t, r.Next.Next)
	assert.Equal(t, 1, r.Next.Next.Val)
}

