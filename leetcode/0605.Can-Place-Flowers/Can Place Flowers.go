package leetcode

// diffculty: easy
// problem link: https://leetcode.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	if len(flowerbed) == 0 {
		return 0 >= n
	}

	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return 1 >= n
	}

	if len(flowerbed) == 1 && flowerbed[0] == 1 {
		return 0 >= n
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		count++
	}

	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count++
		}
	}

	if flowerbed[len(flowerbed)-2] == 0 && flowerbed[len(flowerbed)-1] == 0 {
		count++
	}

	return count >= n
}
