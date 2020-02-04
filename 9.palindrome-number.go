// package main
// 
// 
// import "fmt"
// 
// 
func isPalindrome(x int) bool {
	if x < 0 { return false }
	arr := []int{}
	for x > 0 {
		arr = append(arr, x % 10)
		x = x / 10
	}
	lena := len(arr)
	for i, j := 0, lena-1; i < j; {
		if arr[i] != arr[j] { return false }
		i ++
		j --
	}
	return true
}
// 
// 
// func main() {
// 	integer := 121
// 	fmt.Println(isPalindrome(integer))
// 	integer = 10
// 	fmt.Println(isPalindrome(integer))
// }
// 
// 
