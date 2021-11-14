package leetcode

import (
	"sort"
)

// diffculty: medium
// problem link: https://leetcode.com/problems/non-overlapping-intervals/

type Intervals [][]int

func (intervals Intervals) Len() int {
	return len(intervals)
}

func (intervals Intervals) Less(i, j int) bool {
	return intervals[i][1] < intervals[j][1]
}

func (intervals Intervals) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	min := 0

	if len(intervals) < 2 {
		return min
	}

	sort.Sort(Intervals(intervals))

	pre, skip := 0, 0
	for i := 1; i < len(intervals); i++ {
		if intervals[pre][1] > intervals[i][0] {
			min++
			skip++
		} else {
			pre += skip + 1
			skip = 0
		}
	}

	return min
}
