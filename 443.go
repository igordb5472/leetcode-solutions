package main

import "strconv"

func compress(chars []byte) int {
	newChars, overwritten := make([]byte, 0, len(chars)), chars[:0:len(chars)]
	prev, cnt := chars[0], 1
	for i := 1; i < len(chars) && len(newChars) < cap(newChars); i++ {
		if chars[i] != prev {
			newChars = append(newChars, prev)
			if cnt > 1 {
				cntBytes := []byte(strconv.Itoa(cnt))
				for _, b := range cntBytes {
					if len(newChars) == cap(newChars) {
						overwritten = append(overwritten, newChars...)
						return len(newChars)
					}
					newChars = append(newChars, b)
				}
			}
			prev, cnt = chars[i], 1
		} else {
			cnt++
		}
	}
	if len(newChars) == cap(newChars) {
		overwritten = append(overwritten, newChars...)
		return len(newChars)
	}
	newChars = append(newChars, chars[len(chars)-1])
	if cnt > 1 {
		cntBytes := []byte(strconv.Itoa(cnt))
		for _, b := range cntBytes {
			if len(newChars) == cap(newChars) {
				overwritten = append(overwritten, newChars...)
				return len(newChars)
			}
			newChars = append(newChars, b)
		}
	}
	overwritten = append(overwritten, newChars[:min(len(newChars), len(chars))]...)
	return len(newChars)
}
