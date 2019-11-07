package main

/*
  Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

  For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

  Follow-up: what if you can't use division?
*/

// with division O(2n)
func productOfAllButSelfBasic(xi []int) []int {
	t := 1 // get total of all values
	for _, v := range xi {
		t *= v
	}

	// divide each value by total to get product of all other numbers except self
	rx := make([]int, len(xi)) // return slice
	for i, v := range xi {
		rx[i] = t / v
	}

	return rx
}

// long no division O(n^2)
func productOfAllButSelfLongNoDivision(xi []int) []int {
	rx := make([]int, len(xi))

	for i, x := range xi {
		for j := range xi {
			if i == j { // don't multiply self
				continue
			}
			if rx[j] == 0 { // make sure value is initialized for multiplication
				rx[j] = 1
			}

			rx[j] *= x
		}
	}

	return rx
}

// no division, fast O(2n) - iterate through array 2x
// not my idea - but my implementation
func productOfAllButSelfMagic(xi []int) []int {
	// Create 2 arrays of products - one starting from beginning fibonacci of products - and inverse (starting from end, going down)
	u := make([]int, len(xi)) // starting from beginning going Up
	d := make([]int, len(xi)) // starting from end going Down

	lx := len(xi) - 1 // i keep using this value, lets store it :)

	for i, v := range xi {
		// initialize first val for both Up & Down slices
		if i == 0 {
			u[i] = v
			d[lx] = xi[lx]
		} else {
			// build up products
			u[i] = u[i-1] * v

			// build down products - start from end of array, work towards beginning
			d[lx-i] = d[lx-i+1] * xi[lx-i]
		}
	}

	// with math hacks, we can multiply index before of prefix & index after of suffix & get the final answer
	// (with some edge cases for first and last index)
	rx := make([]int, len(xi)) // result slice
	for i := range xi {
		if i == 0 {
			rx[i] = d[i+1]
		} else if i == lx {
			rx[i] = u[i-1]
		} else {
			// magic :)
			rx[i] = u[i-1] * d[i+1]
		}
	}

	return rx
}
