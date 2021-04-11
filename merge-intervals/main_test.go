package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	// [[]]
	assert.ElementsMatch(t, [][]int{}, merge([][]int{}))

	// [[1,3],[2,6],[8,10],[15,18]] => [[1,6],[8,10],[15,18]]
	assert.ElementsMatch(t, [][]int{{1,6}, {8,10}, {15,18}}, merge([][]int{{1,3},{2,6},{8,10},{15,18}}))

	// [[1,4],[4,5]] => [[1,5]]
	assert.ElementsMatch(t, [][]int{{1,5}}, merge([][]int{{1,4},{4,5}}))

	// [[1,1],[2,2], [3,4]] => [[1,4]
	assert.ElementsMatch(t, [][]int{{1,1}, {2,2}, {3,4}}, merge([][]int{{1,1}, {2,2}, {3,4}}))

	// [[1,2], [3,4], [5,6] [0,7]] => [0, 7]
	assert.ElementsMatch(t, [][]int{{0,7}}, merge([][]int{{1,2}, {3,4}, {5,6}, {0,7}}))
}

func TestMergeInterval(t *testing.T) {
	// eq [1,3] [2, 4] or [1,2] [3, 4] or [1,2], [4,5]
	// [1,1] [2,2], [3,3]
	assert.ElementsMatch(t, []int{1,4}, mergedInterval([]int{1,3}, []int{2,4}), "[]int{1,3}, []int{2,4} => []int{1,4}")
	assert.ElementsMatch(t, []int{1,4}, mergedInterval([]int{2,4}, []int{1,3}), "[]int{2,4}, []int{1,3}, => []int{1,4}")
	assert.ElementsMatch(t, []int{1,4}, mergedInterval([]int{1,3}, []int{3,4}), "[]int{1,3}, []int{3,4} => []int{1,4}")
	assert.ElementsMatch(t, []int{1,4}, mergedInterval([]int{3,4}, []int{1,3}), "[]int{3,4}, []int{1,3}, => []int{1,4}")

	assert.Nil(t, mergedInterval([]int{1,2}, []int{3,4}), "[]int{1,2}, []int{3,4} => nil")
	assert.Nil(t,  mergedInterval([]int{3,4}, []int{1,2}), "[]int{3,4}, []int{1,2} => nil")
	assert.Nil(t,  mergedInterval([]int{4,5}, []int{1,2}), "[]int{4,5}, []int{1,2} => nil")
	assert.Nil(t,  mergedInterval([]int{1,2}, []int{4,5}), "[]int{1,2}, []int{4,5} => Nil")
	assert.Nil(t, mergedInterval([]int{1,1}, []int{2,2}), "[]int{1,1}, []int{2,2} => Nil")
}