package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	assert.Equal(t, -1, firstUniqChar("aa"))
	assert.Equal(t, -1, firstUniqChar(""))
	assert.Equal(t, 6, firstUniqChar("1231234"))
}
