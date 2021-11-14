package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

type param struct {
	intervals [][]int
}

func TestSort(t *testing.T) {
	arr := [][]int{
		{3, 6},
		{2, 4},
		{7, 9},
		{4, 5},
		{1, 2},
	}

	sort.Sort(Intervals(arr))

	fmt.Println(arr)
}

func TestEraseOverlapIntervals(t *testing.T) {
	// [ [1,2], [2,3], [3,4], [1,3] ]
	params := []param{
		{
			intervals: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{1, 3},
			},
		},

		// [ [1,2], [1,2], [1,2] ]
		{
			intervals: [][]int{
				{1, 2},
				{1, 2},
				{1, 2},
			},
		},

		// [ [1,2], [2,3] ]
		{
			intervals: [][]int{
				{1, 2},
				{2, 3},
			},
		},

		// [[1,100],[11,22],[1,11],[2,12]]
		{
			intervals: [][]int{
				{1, 100},
				{11, 22},
				{1, 11},
				{2, 12},
			},
		},

		// [[0,2],[1,3],[2,4],[3,5],[4,6]]
		{
			intervals: [][]int{
				{0, 2},
				{1, 3},
				{2, 4},
				{3, 5},
				{4, 6},
			},
		},
	}

	for _, param := range params {
		fmt.Printf("Input: %v\n", param.intervals)
		fmt.Printf("Output: %d\n", eraseOverlapIntervals(param.intervals))
		fmt.Printf("---------------\n")
	}
}
