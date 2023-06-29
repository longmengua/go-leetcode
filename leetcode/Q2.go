package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var toReturn *ListNode
	var curr *ListNode
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		newV := l1.Val + l2.Val + carry
		carry = newV / 10
		newL := &ListNode{
			Val: newV % 10,
		}
		l1 = l1.Next
		l2 = l2.Next
		if curr == nil {
			curr = newL
			toReturn = newL
		} else {
			curr.Next = newL
			curr = curr.Next
		}
	}

	return toReturn
}
