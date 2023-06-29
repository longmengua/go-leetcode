package leetcode

/*
compare to longest length number when hit the duplicated digit everytime
*/
func LengthOfLongestSubstring(str string) int {
	longest := 0
	currLen := 0
	m := make(map[int32]int)
	// rune type is int32, which means unicode.
	for index, v := range str {
		if _, exist := m[v]; exist {
			if longest < currLen {
				longest = currLen
			}
			if index > len(str) {
				break
			}
			currLen = 1
		} else {
			currLen++
		}
		m[v] = index
	}
	return longest
}
