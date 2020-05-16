package main
import "fmt"


func maxSubArray(nums []int) int {
	lenn := len(nums)
	if lenn == 1 {
		return nums[0]
	}

	DP := make([]int, lenn)
	DP[0] = nums[0]
	max := DP[0]

	for i := 1; i < lenn; i++ {
		if DP[i-1] <= 0 {
			DP[i] = nums[i]
		} else {
			DP[i] = DP[i-1] + nums[i]
		}
		if DP[i] > max {
			max = DP[i]
		}
	}
	return max
}


func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	output := maxSubArray(nums)
	fmt.Println(output)
}
