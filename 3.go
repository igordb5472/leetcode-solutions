package main

func lengthOfLongestSubstring(s string) int {
	maxLen, lastSeen := 0, make([]int, 128)
	for l, r := 1, 1; r <= len(s); r++ {
		l = max(l, lastSeen[s[r-1]]+1)
		lastSeen[s[r-1]] = r
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
