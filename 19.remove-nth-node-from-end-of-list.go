package main
import "fmt"

 type ListNode struct {
     Val int
     Next *ListNode
 }


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	first := pre
	for i := 0; i < n; i ++ {
		pre = pre.Next // 将要删除的前一个节点
	}

	second := first

	for pre.Next != nil {
		second = second.Next
		pre = pre.Next
	}
	second.Next = second.Next.Next
	return first.Next
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

	ret := removeNthFromEnd(l1, 2)

	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
