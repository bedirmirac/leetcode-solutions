/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	firstCurr := head
	secondCurr := reverseList(slow)
	for secondCurr != nil {
		if firstCurr.Val != secondCurr.Val {
			return false
		}
		firstCurr = firstCurr.Next
		secondCurr = secondCurr.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	return prev
}

