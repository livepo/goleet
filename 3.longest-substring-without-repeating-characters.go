// package main
// 
// import "fmt"
// 
// 
func Max(i, j int) int {
	if i <= j {
		return j
	} else {
		return i
	}
}


func lengthOfLongestSubstring(s string) int {
	lens := len(s)
	m := make(map[rune]int)
	max_length := 0
	for i, j := 0, 0; j < lens; {
		if v, _ := m[rune(s[j])]; v != 0 {
			for s[i] != s[j] {
				m[rune(s[i])] = 0
				i ++
			}
			i ++
		} else {
			m[rune(s[j])] ++
			max_length = Max(j-i+1, max_length)
		}
		j ++
	}
	return max_length
}
// 
// 
// func main() {
// 	str := "pwwkew"
// 	fmt.Println(lengthOfLongestSubstring(str))
// }
