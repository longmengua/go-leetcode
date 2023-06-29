package leetcode

/**/
func AddTwoNumbersAns(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	l3 := head
	p1 := l1
	p2 := l2
	c := 0

	for p1 != nil || p2 != nil {
		if p1 != nil {
			c += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			c += p2.Val
			p2 = p2.Next
		}
		l3.Next = &ListNode{
			Val: c % 10,
		}
		l3 = l3.Next
		c /= 10
	}
	if c == 1 {
		l3.Next = &ListNode{
			Val: c,
		}
	}

	return head.Next
}
