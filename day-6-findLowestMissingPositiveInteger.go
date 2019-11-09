package main

/*
	Given an array of integers, find the first missing positive integer.

	In other words, find the lowest positive integer that does not exist in the array.
	The array can contain duplicates and negative numbers as well.

	For example, the input [3, 4, -1, 1] should give 2.
	The input [1, 2, 0] should give 3.
*/

func lowestMissingPositiveInteger(xi []int) int {
	li := 1 // lowest int
	m := make(map[int]bool, len(xi))

	for _, v := range xi {
		if v == li {
			li = checkMapFindNextEmpty(m, li+1)
		}
		if v > 0 {
			m[v] = true
		}
	}

	return li
}

func checkMapFindNextEmpty(m map[int]bool, li int) int {
	nli := li // new lowest int
	for {
		if _, ok := m[nli]; ok {
			nli++
		} else {
			return nli
		}
	}
}
