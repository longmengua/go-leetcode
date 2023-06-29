package leetcode

import (
	"testing"
)

/*
Given a string s, find the length of the longest substring without repeating characters.
*/
func TestQ3(t *testing.T) {
	s := "abcabcbb"
	result := LengthOfLongestSubstring(s)
	Assert(t, result, 3)

	s = "bbbbb"
	result = LengthOfLongestSubstring(s)
	Assert(t, result, 1)

	s = "pwwkew"
	result = LengthOfLongestSubstring(s)
	Assert(t, result, 3)

	s = " "
	result = LengthOfLongestSubstring(s)
	Assert(t, result, 1)

	s = ""
	result = LengthOfLongestSubstring(s)
	Assert(t, result, 0)

	s = "au"
	result = LengthOfLongestSubstring(s)
	Assert(t, result, 2)

	s = "dvdf"
	result = LengthOfLongestSubstring(s)
	Assert(t, result, 3)

	s = "abba"
	result = LengthOfLongestSubstring(s)
	Assert(t, result, 2)
}
