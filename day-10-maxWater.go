package main

import "fmt"

/*
	Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

	Note: You may not slant the container and n is at least 2.
*/

/*
	Based on day-3-doTwoNumbersAddUpToK (sorted) solution.
		point to outer 'walls', and move inwards. Works because as width gets smaller (ie: i:0 j:10: width=10 vs i:4 j:6: width = 4), we know the smallest height of either index must be greater than previously found min-height wall (h)
*/
func maxArea(height []int) int {
	water := 0
	i := 0
	j := len(height) - 1

	for i < j {
		h := min(height[i], height[j])
		water = max(water, h*(j-i))
		fmt.Println(water, h, i, j)
		for height[i] <= h && i < j {
			fmt.Println("  - i++")
			i++
		}
		for height[j] <= h && i < j {
			fmt.Println("  - j--")
			j--
		}
	}

	return water
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
