package leetcode

/*
compare to longest length number when hit the duplicated digit everytime
*/
func LengthOfLongestSubstring(s string) int {
	longest := 0
	currLen := 0
	head := 0
	m := make(map[int32]int)

	// rune type is int32, which means unicode.
	for index, v := range s {
		if _, exist := m[v]; exist {
			if m[v] > head {
				head = m[v]
			}
			currLen = index - head
		} else {
			currLen++
		}
		m[v] = index
		if currLen > longest {
			longest = currLen
		}
	}
	return longest
}
