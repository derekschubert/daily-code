package main

/*
	Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical.
	For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return the node with value 8.
*/

type node struct {
	Val  int
	Next *node
}

// point at beginning of each list
// walk the pointers down based on which value is higher
// if we find a match, return mem address to node
// if we ___ & still no match, return nil
// - reach end of a list && the val in the other list is HIGHER than (last) val in current list
// - reach end of both lists
func orderedListIntersection(l1 *node, l2 *node) *node {
	n1 := *l1
	n2 := *l2

	for {
		if n1.Val == n2.Val {
			return &n1
		} else if n1.Val < n2.Val && n1.Next != nil {
			n1 = *n1.Next
		} else if n1.Val > n2.Val && n2.Next != nil {
			n2 = *n2.Next
		} else {
			return nil
		}
	}
}

// just make a hash map that keeps track of values in l1
// then iterate through l2 and check if key exists
func unorderedListIntersection(l1 *node, l2 *node) *node {
	m := make(map[int]bool)
	n1 := l1

	for n1 != nil {
		m[n1.Val] = true
		n1 = n1.Next
	}

	n2 := l2
	for n2 != nil {
		if _, ok := m[n2.Val]; ok {
			return n2
		}
		n2 = n2.Next
	}

	return nil
}
