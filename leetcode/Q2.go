package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var toReturn *ListNode
	var ans *ListNode
	ln1 := l1
	ln2 := l2
	carry := 0

	for ln1 != nil || ln2 != nil {
		newV := ln1.Val + ln2.Val + carry
		carry = newV / 10
		newL := &ListNode{
			Val: newV % 10,
		}
		ln1 = ln1.Next
		ln2 = ln2.Next
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
