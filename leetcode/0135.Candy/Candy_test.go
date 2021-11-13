package leetcode

import (
	"fmt"
	"testing"
)

type param struct {
	ratings []int
}

func TestCandy(t *testing.T) {
	params := []param{
		{
			ratings: []int{1, 0, 2},
		},
		{
			ratings: []int{1, 2, 2},
		},
	}

	for _, param := range params {
		fmt.Printf("Input:\n%v\n", param.ratings)
		fmt.Printf("Output:\n%d\n", candy(param.ratings))
		fmt.Printf("---------------\n")
	}
}
