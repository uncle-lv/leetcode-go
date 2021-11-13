package leetcode

// diffculty: hard
// problem link: https://leetcode.com/problems/candy/

func candy(ratings []int) int {
	num := len(ratings)

	if num == 1 {
		return num
	}

	candyArr := make([]int, num)
	for i := 0; i < num; i++ {
		candyArr[i] = 1
	}

	for i := 1; i < num; i++ {
		if ratings[i] > ratings[i-1] {
			candyArr[i] = candyArr[i-1] + 1
		}
	}

	for i := num - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			candyArr[i-1] = max(candyArr[i-1], candyArr[i]+1)
		}
	}

	num = 0
	for _, value := range candyArr {
		num += value
	}

	return num
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
