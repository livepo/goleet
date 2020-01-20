// package main
// 
// import "fmt"
// 
// 
// type ListNode struct {
//     Val int
//     Next *ListNode
// }
// 
// 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pri := &ListNode{}
	cur := pri
	out, c := 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			out = l1.Val + l2.Val + c
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			out = l1.Val + c
			l1 = l1.Next
		} else {
			out = l2.Val + c
			l2 = l2.Next
		}
		if out >= 10 {
			out = out - 10
			c = 1
		} else {
			c = 0
		}
		cur.Next = &ListNode{Val: out}
		cur = cur.Next
	}
	if c == 1 {
		cur.Next = &ListNode{Val: c}
	}
	return pri.Next
}
// 
// 
// func main() {
// 	n1 := &ListNode{Val: 2}
// 	n2 := &ListNode{Val: 4}
// 	n3 := &ListNode{Val: 3}
// 	n1.Next, n2.Next = n2, n3
// 
// 	n4 := &ListNode{Val: 5}
// 	n5 := &ListNode{Val: 6}
// 	n6 := &ListNode{Val: 4}
// 	n4.Next, n5.Next = n5, n6
// 
// 	p := addTwoNumbers(n1, n4)
// 	for p != nil {
// 		fmt.Println(p.Val)
// 		p = p.Next
// 	}
// }
