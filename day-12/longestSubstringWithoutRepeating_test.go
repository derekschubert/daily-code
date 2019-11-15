package main

import (
	"testing"
)

type test12 struct {
	s string // input String
	l int    // response Length
}

var tests12Substrings = []test12{
	// repeating substrings (ie: abcab passes, abcabc fails [abc|abc])
	{"abcabcbb", 5}, // abcab
	{"bbbbb", 1},    // b
	{"pwwkew", 4},   // wkew
}

var tests12Characters = []test12{
	// repeating characters (ie: abca invalid [2 'a's])
	{"abcabcbb", 3}, // abc
	{"bbbbb", 1},    // b
	{"pwwkew", 3},   // wke
	{"aab", 2},      // ab
	{"dvdf", 3},     // vdf
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, pair := range tests12Substrings {
		r := lengthOfLongestSubstring(pair.s)
		if r != pair.l {
			t.Error(
				"For", pair.s,
				"expected", pair.l,
				"got:", r,
				"\n",
			)
		}
	}
}

func TestLengthOfLongestSubstringCharacters(t *testing.T) {
	for _, pair := range tests12Characters {
		r := lengthOfLongestSubstringCharacters(pair.s)
		if r != pair.l {
			t.Error(
				"For", pair.s,
				"expected", pair.l,
				"got:", r,
				"\n",
			)
		}
	}
}
