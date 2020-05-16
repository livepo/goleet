package main

import "fmt"


func mySqrt(x int) int {
	start := 4.0
	for {
		output := (start + float64(x) / start) / 2.0
		if output - start >= -0.2 && output - start <= 0.2 {
			return int(output)
		}
		start = output
	}
}



func main() {
	output := mySqrt(4)
	fmt.Println(output)
}
