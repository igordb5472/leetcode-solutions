package main

func longestPalindrome(s string) string {
	var start, end int
	for i := 0; i < len(s)-1; i++ {
		l, r := i, i
		for l >= 1 && r < len(s)-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l+1 > end-start+1 {
			start, end = l, r
		}
		if s[i] == s[i+1] {
			l, r = i, i+1
			for l >= 1 && r < len(s)-1 && s[l-1] == s[r+1] {
				l--
				r++
			}
			if r-l+1 > end-start+1 {
				start, end = l, r
			}
		}
	}
	return s[start : end+1]
}
