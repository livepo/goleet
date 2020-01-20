// 这方法没有用字典的方法好，只是练习一下语法才这么写的
// package main
// 
// import (
// 	"fmt"
// 	"sort"
// )
// 
type RankData struct {value, idx int}
type RDSlice []RankData

func (r RDSlice) Len() int { return len(r) }
func (r RDSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r RDSlice) Less(i, j int) bool { return r[i].value < r[j].value }


func twoSum(nums []int, target int) []int {
	m := []RankData{}
	for idx, n := range nums {
		m = append(m, RankData{value: n, idx: idx})
	}
	sort.Sort(RDSlice(m))
	for idx, n := range m {
		if i, ok := bsearch(m, target - n.value, idx+1, len(m)-1); ok {
			return []int{m[idx].idx, m[i].idx}
		}
	}
	return []int{-1, -1}
}


func bsearch(m []RankData, target, left, right int) (int, bool) {
	for left <= right {
		mid := (left + right) / 2
		if m[mid].value == target {
			return mid, true
		} else if m[mid].value > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1, false
}
// 
// 
// func main() {
// 	a := twoSum([]int{0,4,3,0}, 7)
// 	fmt.Println(a)
// }
