package main

/*
  Given a list of numbers and a number k,
  return whether any two numbers from the list add up to k.

  For example, given [10, 15, 3, 7] and k of 17,
	return true since 10 + 7 is 17.

	---

	Two solutions: one if the array is sorted, one if it is unsorted.
*/

/*
	If array is sorted, have two index pointers (high & low).
	Keep moving the pointers inwards based off of xi[high] + xi[low] values
	If the pointers are looking at the same place or pass eachother, no pair exists.
*/
func doTwoNumbersAddUp_Sorted(xi []int, k int) bool {
	l := 0           // low
	h := len(xi) - 1 // high

	for l < h {
		t := xi[l] + xi[h] // total
		if t == k {
			return true
		} else if t < k {
			l++
		} else { // t > k
			h--
		}
	}

	return false
}

/*
	For each value in the slice, store the complement in a map
	(complement == value that adds up to k: ie if k = 9 && xi[i] = 4, complement = 5 (5 + 4 = 9))
	Check to see if current value is a complement to any previous value
*/
func doTwoNumbersAddUp_Unsorted(xi []int, k int) bool {
	m := make(map[int]bool)

	for i := range xi {
		c := k - xi[i] // complement
		if _, ok := m[xi[i]]; ok {
			return true
		}
		m[c] = true
	}

	return false
}
