/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenHeadA, lenHeadB := headA, headB
	lenA, lenB := 0, 0
	currA, currB := headA, headB
	for lenHeadA != nil {
		lenA++
		lenHeadA = lenHeadA.Next
	}
	for lenHeadB != nil {
		lenB++
		lenHeadB = lenHeadB.Next
	}
	different := lenA - lenB
	if different < 0 {
		for different != 0 {
			currB = currB.Next
			different++
		}
	}
	if different > 0 {
		for different != 0 {
			currA = currA.Next
			different--
		}
	}
	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
	}
	return nil
}

