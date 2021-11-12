package leetcode

import "sort"

// diffculty: easy
// problem link: https://leetcode.com/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	i, j := 0, 0

	sort.Ints(g)
	sort.Ints(s)

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}

	return i
}
