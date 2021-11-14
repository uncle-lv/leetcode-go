package leetcode

import (
	"fmt"
	"testing"
)

type param struct {
	flowerbed []int
	n         int
}

func TestCanPlaceFlowers(t *testing.T) {
	params := []param{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
		},

		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
		},

		{
			flowerbed: []int{0, 0, 1, 0, 1},
			n:         1,
		},

		{
			flowerbed: []int{0},
			n:         1,
		},
	}

	for _, param := range params {
		fmt.Printf("Input:\nflowerbed: %v\nn: %d\n", param.flowerbed, param.n)
		fmt.Printf("Output:\n%t\n", canPlaceFlowers(param.flowerbed, param.n))
		fmt.Printf("---------------\n")
	}
}
