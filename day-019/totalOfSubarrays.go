package main

import "fmt"

/*
	Given an array of integers and a number k, where 1 <= k <= length of the array,
	compute the sum of each subarray of length k.

	You can modify the given array in place.
*/

func main() {
	a := getTotals([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(a)
	a = getTotals([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(a)
	a = getTotals([]int{1, 2, 3, 4, 5}, 1)
	fmt.Println(a)
}

func getTotals(xi []int, k int) []int {
	ans := make([]int, len(xi)-k+1)
	p := 0

	for i := range xi {
		if i == 0 {
			continue
		}

		xi[i] += xi[i-1]

		if i >= k {
			xi[i] -= xi[i-k]
			if i >= k+1 {
				xi[i] += xi[i-k-1]
				xi[i-k] -= xi[i-k-1]
			}
		}

		if i >= k-1 {
			ans[p] = xi[i]
			p++
		}
	}

	return ans
}
