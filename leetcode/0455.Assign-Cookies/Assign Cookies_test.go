package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	g []int
	s []int
}

func TestSolution(t *testing.T) {

	paramsArr := []params{
		{
			g: []int{1, 2, 3},
			s: []int{1},
		},

		{
			g: []int{1, 2},
			s: []int{1, 2, 3},
		},
	}

	for _, params := range paramsArr {
		g, s := params.g, params.s
		fmt.Printf("Input:\ng: %v\ns: %v\n", g, s)
		fmt.Printf("Output:\n%d\n", findContentChildren(g, s))
		fmt.Printf("---------------\n")
	}
}
