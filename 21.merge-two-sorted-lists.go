package main

import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return head.Next
}


func main() {
	var n1 *ListNode = &ListNode{Val: 0}
	var n2 *ListNode = &ListNode{Val: 7}
	var n3 *ListNode = &ListNode{Val: 8}
	var n7 *ListNode = &ListNode{Val: 9}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n7
	l1 := n1

	var n4 *ListNode = &ListNode{Val: 3}
	var n5 *ListNode = &ListNode{Val: 4}
	var n6 *ListNode = &ListNode{Val: 9}
	n4.Next = n5
	n5.Next = n6
	l2 := n4

	p := mergeTwoLists(l1, l2)

	head := p
	for ; head != nil; {
		fmt.Println(head.Val)
		head = head.Next
	}
}
