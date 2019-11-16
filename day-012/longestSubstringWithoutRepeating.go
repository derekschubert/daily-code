package main

// finds longest w/o mirrored/repeating SUBSTRINGS
// ie: abcabcbb = 5 (abcab), because abcabc repeats (abc|abc)
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 { // for 0 & 1 len strings
		return len(s)
	}
	longest := 1
	cur := []int32{}

	for _, v := range s {
		cur = append(cur, v)

		if len(cur) > 1 {
			for i := 1; i <= (len(cur) / 2); i++ {
				if cur[len(cur)-1] == cur[len(cur)-1-i] {
					pass := true

					for j := 1; j < len(cur)-i-1; j++ {
						t := len(cur) - j - 1
						t2 := len(cur) - i - j - 1

						if cur[t] != cur[t2] {
							pass = false
							break
						}
					}

					if pass == true {
						cur = []int32{v}
						break
					}
				}
			}
		}
		if len(cur) > longest {
			longest = len(cur)
		}
	}

	return longest
}

// finds longest w/o repeating CHARACTERS
// ie: abcabcbb = 3 (abc), because abca has 2 'a' chars
// can be further optimized, but i need to do some other productive things today as well :D
func lengthOfLongestSubstringCharacters(s string) int {
	if len(s) < 2 { // for 0 & 1 length strings
		return len(s)
	}

	var ans, i, j int
	n := len(s)
	m := make(map[byte]int, n)

	for i < n && j < n {
		if _, ok := m[s[j]]; !ok {
			m[s[j]] = 1
			j++
			ans = max(ans, j-i)
		} else {
			delete(m, s[i])
			i++
		}
	}

	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
