package main

import "fmt"

func MaxValue(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}


func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	lenc := len(candies)
	retarr := make([]bool, 0)
	for i := 0; i < lenc; i ++ {
		max = MaxValue(max, candies[i])
	}

	for i := 0; i < lenc; i ++ {
		isMax := candies[i] + extraCandies >= max
		retarr = append(retarr, isMax)
	}
	return retarr
}


func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
