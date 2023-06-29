package leetcode

import (
	"reflect"
	"testing"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
func TestQ2(t *testing.T) {
	result := ListNodeToInt(AddTwoNumbers(
		IntToListNode([]int{2, 4, 3}),
		IntToListNode([]int{5, 6, 4}),
	))
	isEqual := reflect.DeepEqual(result, []int{7, 0, 8})
	Assert(t, isEqual, true)
}

func IntToListNode(nums []int) *ListNode {
	var toReturn *ListNode
	if len(nums) == 0 {
		return nil
	}

	var l *ListNode
	for i, v := range nums {
		if l == nil {
			l = &ListNode{
				Val:  v,
				Next: &ListNode{},
			}
			toReturn = l
		} else {
			l = l.Next
			l.Val = v
			if i+1 < len(nums) {
				l.Next = &ListNode{}
			}
		}
	}

	return toReturn
}

func ListNodeToInt(l *ListNode) []int {
	toReturn := []int{}
	if l == nil {
		return nil
	}

	var p = l
	for p != nil {
		toReturn = append(toReturn, p.Val)
		p = p.Next
	}

	return toReturn
}
