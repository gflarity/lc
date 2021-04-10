package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestValidPalindrome(t *testing.T) {
	assert.True(t, validPalindrome("aba"), "aba should be true")
	assert.True(t, validPalindrome("abba"), "abba should be true")
	assert.True(t, validPalindrome("abbac"), "abbac has one deletion, it should be true")
	assert.False(t, validPalindrome("zac"), "zac has two deletions, it should be false")
	assert.False(t, validPalindrome("zabax"), "zabax requires two deletions, it should be false")
	assert.False(t, validPalindrome("dabzbac"), "dabzbac has two deletions, it should be false")
	assert.True(t, validPalindrome("lcuucul"), "should be true")
	assert.True(t, validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"), "should be true")
}
