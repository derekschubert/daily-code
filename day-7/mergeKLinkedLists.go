package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	var head *ListNode
	curNode := &ListNode{}

	for {
		hit := false    // ensure we haven't finished going through all lists
		var nl int      // next lowest val
		var indList int // index of list w/ next lowest val
		for i, l := range lists {
			if l == nil {
				continue
			} else if l.Val < nl || !hit {
				nl = l.Val
				indList = i
				hit = true
			}
		}

		// no more numbers to sort, end
		if !hit {
			return head
		}

		// increment used node to next node in its given list
		lists[indList] = lists[indList].Next

		// if first iteration - initialize head node
		if head == nil {
			head = &ListNode{Val: nl}
		} else { // on 2nd+ iteration
			if head.Next == nil { // on 2nd iteration
				head.Next = curNode
			} else if curNode != nil { // on 4th+ iteration
				curNode.Next = &ListNode{}
				curNode = curNode.Next
			}

			curNode.Val = nl
		}
	}
}

func tests() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	l := []*ListNode{l1, l2, l3}
	o := mergeKLists(l) // should be 1->1->2->3->4->4->5->6

	for o != nil {
		fmt.Println(o)
		o = o.Next
	}
}
