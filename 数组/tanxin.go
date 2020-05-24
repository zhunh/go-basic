package main

import "math"

/*
贪心算法
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	res := math.MinInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res = max(sum, res)
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

func main() {

}
