/**---------------------------------------------------------
 * SOURCE: https://leetcode.com/problems/add-two-numbers/
 * ---------------------------------------------------------
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Example:
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 * Runtime result: 8 ms, faster than 89.88% of Go online submissions for Add Two Numbers.
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nl := ListNode{0, nil}
	tn := &nl

	c := 0
	hit := false

	for l1 != nil || l2 != nil {
		v := 0
		if l1 != nil && l2 != nil {
			v = (l1.Val + l2.Val + c)

			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			v = (l1.Val + c)
			l1 = l1.Next
		} else {
			v = (l2.Val + c)
			l2 = l2.Next
		}

		if v >= 10 {
			c = 1
		} else {
			c = 0
		}

		if hit {
			tn.Next = &ListNode{0, nil}
			tn = tn.Next
		} else {
			hit = true
		}

		tn.Val = v % 10
	}

	if c == 1 {
		tn.Next = &ListNode{1, nil}
	}

	return &nl
}
