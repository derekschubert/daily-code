package main

import (
	"testing"
)

type test struct {
	l1 *node
	l2 *node
	a  int
}

func buildTests() []test {
	l1a := &node{0, &node{1, &node{2, &node{3, &node{4, nil}}}}}
	l1b := &node{3, &node{5, &node{6, &node{7, &node{8, nil}}}}}
	a1 := 3
	t1 := test{l1a, l1b, a1}

	l2a := &node{5, &node{3, &node{4, &node{9, &node{1, nil}}}}}
	l2b := &node{15, &node{1, &node{5, &node{12, &node{1, nil}}}}}
	a2 := 1
	t2 := test{l2a, l2b, a2}

	return []test{t1, t2}
}

func TestOrderedList(t *testing.T) {
	tests := buildTests()

	for _, v := range tests {
		a := unorderedListIntersection(v.l1, v.l2)
		if a == nil {
			t.Error("Expected", v.a, "but got nil")
		} else if a.Val != v.a {
			t.Error("Expected", v.a, "but got", a.Val)
		}
	}
}
