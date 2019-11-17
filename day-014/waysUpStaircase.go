package main

/*
	There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time. Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.

	For example, if N is 4, then there are 5 unique ways:

		1, 1, 1, 1
		2, 1, 1
		1, 2, 1
		1, 1, 2
		2, 2

	What if, instead of being able to climb 1 or 2 steps at a time, you could climb any number from a set of positive integers X? For example, if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time.
*/

func waysUpStaircase(steps int, climb []int) int {
	if steps == 0 {
		return 0
	}

	return _up(steps, &climb) + 1
}

func _up(steps int, climb *[]int) int {
	if steps <= 1 {
		return 0
	}

	t := 0

	for i, v := range *climb {
		if steps-v >= 0 {
			if i != 0 {
				t++
			}

			t += _up(steps-v, climb)
		}

	}

	return t
}
