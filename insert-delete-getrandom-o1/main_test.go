package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	rs := Constructor();
	assert.True(t, rs.Insert(1))
	assert.Equal(t, 1, rs.GetRandom())
	assert.True(t, rs.Insert(2))
	assert.True(t, rs.Remove(1))
	assert.Equal(t, 2, rs.GetRandom())

	// ["RandomizedSet","remove","remove","insert","getRandom","remove","insert"]
	//[[],[0],[0],[0],[],[0],[0]]
	rs = Constructor()
	assert.False(t, rs.Remove(0))
	assert.False(t, rs.Remove(0))
	assert.True(t, rs.Insert(0))
	assert.Equal(t, 0, rs.GetRandom())
	assert.True(t, rs.Remove(0))
	assert.True(t, rs.Insert(0))

	//["RandomizedSet","insert","insert","getRandom","getRandom","insert","remove","getRandom","getRandom","insert","remove"]
	//[[],[3],[3],[],[],[1],[3],[],[],[0],[0]]
	rs = Constructor()
	assert.True(t, rs.Insert(3))
	assert.False(t, rs.Insert(3))
	assert.Equal(t, 3, rs.GetRandom())
	assert.Equal(t, 3, rs.GetRandom())
	assert.True(t, rs.Insert(1))
	assert.True(t, rs.Remove(3))
	assert.Equal(t, 1, rs.GetRandom())
	assert.Equal(t, 1, rs.GetRandom())
	assert.True(t, rs.Insert(0))
	assert.True(t, rs.Remove(0))


}
