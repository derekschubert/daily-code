package main

import "testing"

type test struct {
	Steps  int
	Climb  []int // different amount of stairs you can climb in one move
	Answer int
}

var tests = []test{
	{Steps: 4, Climb: []int{1, 2}, Answer: 5},
	{Steps: 2, Climb: []int{1, 2}, Answer: 2},
	{Steps: 3, Climb: []int{1, 2}, Answer: 3},
	{Steps: 3, Climb: []int{1, 2, 3}, Answer: 4},
	{Steps: 5, Climb: []int{1, 2, 3}, Answer: 13},
}

func TestWaysUpStaircase(testing *testing.T) {
	for _, t := range tests {
		a := waysUpStaircase(t.Steps, t.Climb)
		if a != t.Answer {
			testing.Error(
				"Fail: for", t.Steps,
				"expected:", t.Answer,
				"but got:", a,
			)
		}
	}
}
