package leetcode

/*
compare to longest length number when hit the duplicated digit everytime
*/
func LengthOfLongestSubstring(s string) int {
	longest := 0
	head := 0
	m := make(map[int32]int)

	// rune type is int32, which means unicode.
	for index, v := range s {
		if _, exist := m[v]; exist {
			if m[v] >= head {
				head = m[v] + 1
			}
		}
		m[v] = index
		if index-head+1 > longest {
			longest = index - head + 1
		}
	}
	return longest
}
