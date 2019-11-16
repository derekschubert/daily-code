package main

import "fmt"

/*
	Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

	For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
*/

func test() {
	e := []int{2, 4, 6, 2, 5}
	a := largestSum(e)
	fmt.Println(a, "== 13")

	e = []int{5, 1, 1, 5}
	a = largestSum(e)
	fmt.Println(a, "== 10")
}

// O(n) time - just looping straight through once.
func largestSum(nums []int) int {
	p := 0 // previous
	p2 := 0

	for _, v := range nums {
		t := p
		p = max(p2+v, p)
		p2 = t
	}

	return p
}

// Recursive version with memoization
func largestSumRecursive(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return _largestSum(nums, len(nums)-1, memo)
}

func _largestSum(nums []int, i int, memo []int) int {
	if i < 0 {
		return 0
	}
	if memo[i] >= 0 {
		return memo[i]
	}
	memo[i] = max(_largestSum(nums, i-2, memo)+nums[i], _largestSum(nums, i-1, memo))
	return memo[i]
}

// helper :)
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
