/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var p *ListNode
	c := head
	for c != nil {
		if c.Val == val {
			if c == head {
				c = c.Next
				head = c
				continue
			} else {
				p.Next = c.Next
				c = c.Next
				continue
			}
		}
		p = c
		c = c.Next
	}
	return head
}

