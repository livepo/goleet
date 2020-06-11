package main

import "fmt"


 type ListNode struct {
     Val int
     Next *ListNode
 }


func swapPairs(head *ListNode) *ListNode {
	priv := new(ListNode)
	priv.Next = head

	ret := priv
	for priv.Next != nil && priv.Next.Next != nil {
		p := priv
		first, second := priv.Next, priv.Next.Next
		tmp := second.Next
		second.Next = first
		first.Next = tmp
		p.Next = second
		priv = first
	}
	return ret.Next
}


func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	ret := swapPairs(l1)
	p := ret
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
