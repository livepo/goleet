/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	priv, cur := &new(ListNode), head
	priv.Next = cur

	for cur != nil {
		priv.Next, cur.Next, cur = cur.Next, cur, cur.Next
	}

}
