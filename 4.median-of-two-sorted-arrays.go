// package main
// 
// import "fmt"
// 
// 
func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	return float64(getKth(nums1, len1, nums2, len2, (len1+len2+1)>>1) +
	        getKth(nums1, len1, nums2, len2, (len1+len2+2)>>1)) / 2
}


func getKth(nums1 []int, len1 int, nums2 []int, len2 int, k int) int {
	if len1 > len2 {
		return getKth(nums2, len2, nums1, len1, k)
	}

	if len1 == 0 {
		return nums2[k-1]
	}

	if k == 1 {
		return Min(nums1[0], nums2[0])
	}
	i := Min(len1, k / 2)
	j := Min(len2, k / 2)
	if nums1[i-1] > nums2[j-1] {
		return getKth(nums1, len1, nums2[j:], len2-j, k-j)
	} else {
		return getKth(nums1[i:], len1-i, nums2, len2, k-i)
	}
}
// 
// 
// func main() {
// 	arr1 := []int{1, 3}
// 	arr2 := []int{2, 4}
// 	ret := findMedianSortedArrays(arr1, arr2)
// 	fmt.Println(ret)
// }
