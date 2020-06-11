package main

import "fmt"


func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	row := len(matrix)
	if row == 0 { return ret }
	col := len(matrix[0])

	i, j := 0, 0
	left, right, up, down := 0, col, 0, row
	for {
		if left + 1 == right && up + 1 == down {
			break
		}
		fmt.Println("left to right")
		for k := left; k < right; k ++ {
			ret = append(ret, matrix[i][j])
			j ++
		}
		i ++
		up ++
		fmt.Println("up to down")
		for k := up; k < down; k ++ {
			ret = append(ret, matrix[i][j])
			i ++
		}
		j --
		right --
		fmt.Println("right to left")
		for k := right-1; k >= left; k -- {
			ret = append(ret, matrix[i][j])
			j --
		}
		i --
		down ++
		fmt.Println("down to up")
		for k := down-1; k >= up; k -- {
			ret = append(ret, matrix[i][j])
			i --
		}
	}
	return ret
}


func main() {
	matrix := make([][]int, 0)
	arr1 := []int{1,  2,  3,  4}
	arr2 := []int{5,  6,  7,  8}
	arr3 := []int{9, 10, 11, 12}
	matrix = append(matrix, arr1)
	matrix = append(matrix, arr2)
	matrix = append(matrix, arr3)
	ret := spiralOrder(matrix)
	fmt.Println(ret)
}
