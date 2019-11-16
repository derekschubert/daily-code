package main

import "fmt"

/*
	Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
*/

// Build max heights from both directions (left -> right & right -> left), then for each index find the min height and subtract it from the true height (min(left[i], right[i]) - height[i]).
func trap(height []int) int {
	xhlen := len(height)
	if xhlen <= 2 {
		return 0
	}

	lh := make([]int, xhlen)
	rh := make([]int, xhlen)

	lh[0] = height[0]
	rh[xhlen-1] = height[xhlen-1]

	for i, h := range height {
		// lh
		if i != 0 {
			if h > lh[i-1] {
				lh[i] = h
			} else {
				lh[i] = lh[i-1]
			}

			ri := xhlen - 1 - i
			if height[ri] > rh[ri+1] {
				rh[ri] = height[ri]
			} else {
				rh[ri] = rh[ri+1]
			}
		}
	}

	total := 0
	for i, h := range height {
		if lh[i] < rh[i] {
			total += lh[i] - h
		} else {
			total += rh[i] - h
		}
	}

	return total
}

// going to begin leaving the test cases in, instead of deleting - may prove useful to others or myself upon any future revisits
func test() {
	ex := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	o := trap(ex)
	fmt.Println(o, "== 6")
	fmt.Println()
	fmt.Println()

	ex2 := []int{2, 0, 2}
	o2 := trap(ex2)
	fmt.Println(o2, "== 2")
	fmt.Println()
	fmt.Println()

	ex3 := []int{4, 2, 3}
	o3 := trap(ex3)
	fmt.Println(o3, "== 1")
	fmt.Println()
	fmt.Println()
}
