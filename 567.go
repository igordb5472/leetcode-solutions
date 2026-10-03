package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	diff := make(map[byte]int, len(s2))
	for i := 0; i < len(s1); i++ {
		diff[s1[i]]--
		if diff[s1[i]] == 0 {
			delete(diff, s1[i])
		}
		diff[s2[i]]++
		if diff[s2[i]] == 0 {
			delete(diff, s2[i])
		}
	}
	if len(diff) == 0 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		diff[s2[i-len(s1)]]--
		if diff[s2[i-len(s1)]] == 0 {
			delete(diff, s2[i-len(s1)])
		}
		diff[s2[i]]++
		if diff[s2[i]] == 0 {
			delete(diff, s2[i])
		}
		if len(diff) == 0 {
			return true
		}
	}
	return false
}
