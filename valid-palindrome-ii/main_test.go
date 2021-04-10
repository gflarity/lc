package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("a 1 a"))
	assert.True(t, isPalindrome("a 1   1 a"))
	assert.False(t, isPalindrome("a 1 1 b"))
}

