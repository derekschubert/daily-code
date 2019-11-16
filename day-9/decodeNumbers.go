package main

import (
	"fmt"
	"strconv"
)

/*
	Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.

	For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
	You can assume that the messages are decodable. For example, '001' is not allowed.
*/

func test() {
	o := decode("101")
	fmt.Println(o, "== 1")
	o = decode("111")
	fmt.Println(o, "== 3")
	o = decode("12")
	fmt.Println(o, "== 2")
	o = decode("27")
	fmt.Println(o, "== 1")
}

func decode(msg string) int {
	if len(msg) < 2 {
		return len(msg)
	}
	return _d(msg) + 1
}

func _d(msg string) int {
	if len(msg) < 2 {
		return 0
	}

	total := 0

	ci, _ := strconv.Atoi(string(msg[0]))
	ni, _ := strconv.Atoi(string(msg[1]))
	com, _ := strconv.Atoi(fmt.Sprintf("%v%v", ci, ni))

	if ci == 0 {
		// skip, means was counted last time as "10" or "20" (assuming msg is always decodable)
	} else if ni == 0 {
		_d(msg[2:])
	} else {
		if com <= 26 {
			total++
			total += _d(msg[2:])
		}
		total += _d(msg[1:])
	}

	return total
}
