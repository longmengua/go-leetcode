package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var toReturn *ListNode
	var ans *ListNode
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		newV := l1.Val + l2.Val + carry
		carry = newV / 10
		newL := &ListNode{
			Val: newV % 10,
		}
		l1 = l1.Next
		l2 = l2.Next
		if ans == nil {
			ans = newL
			toReturn = newL
		} else {
			ans.Next = newL
			ans = ans.Next
		}
	}

	return toReturn
}
