package main


import "fmt"


func productExceptSelf(nums []int) []int {
	// input: [1, 2, 3, 4]
	// ouput: [24, 12, 8, 6]
	lenn := len(nums)
	right := make([]int, lenn)
	if lenn == 0 { return right }
	i := lenn - 1
	right[i] = 1
	for i := lenn-2; i >= 0; i -- {
		right[i] = nums[i+1] * right[i+1]
	}

	left := 1
	for i := 0; i < lenn; i ++ {
		right[i] = left * right[i]
		left = left * nums[i]
	}
	return right
}


func main() {
	ret := productExceptSelf([]int{3})
	fmt.Println(ret)
}
