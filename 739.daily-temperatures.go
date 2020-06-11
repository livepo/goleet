package main

import "fmt"


type element struct {Val, Idx int}


func dailyTemperatures(T []int) []int {
	ret := make([]int, len(T))
	stack := make([]element, 0)

	for idx, t := range T {
		for len(stack) > 0 && stack[len(stack)-1].Val < t { // 如果栈顶有元素并且小于当前元素，开始计算
			el := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[el.Idx] = idx - el.Idx
		}
		stack = append(stack, element{Val: t, Idx: idx})
	}

	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret[el.Idx] = 0
	}
	return ret
}


func main() {
	ret := dailyTemperatures([]int{73,74,75,71,69,72,76,73})
	fmt.Println(ret)
}
