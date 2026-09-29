package main

func generateParenthesis(n int) []string {
	combs := new([]string)
	addParenthesis(combs, "", n, n, n*2)
	return *combs
}

func addParenthesis(combs *[]string, comb string, forOpen, forClose, maxLen int) {
	if len(comb) == maxLen {
		*combs = append(*combs, comb)
	}
	if forOpen > 0 {
		addParenthesis(combs, comb+"(", forOpen-1, forClose, maxLen)
	}
	if forOpen < forClose {
		addParenthesis(combs, comb+")", forOpen, forClose-1, maxLen)
	}
}
