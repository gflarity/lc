package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "on", longestCommonPrefix([]string{"one", "one hundred", "once"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"one", "two", "three"}))
}
