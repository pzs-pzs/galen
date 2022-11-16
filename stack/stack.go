package stack

func isValid(s string) bool {
	stk := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stk = append(stk, v)
			continue
		}
		if len(stk) == 0 {
			return false
		}
		if v == ')' && stk[len(stk)-1] == '(' {
			stk = stk[:len(stk)-1]
			continue
		}
		if v == ']' && stk[len(stk)-1] == '[' {
			stk = stk[:len(stk)-1]
			continue
		}
		if v == '}' && stk[len(stk)-1] == '{' {
			stk = stk[:len(stk)-1]
			continue
		}
		return false
	}
	return len(stk) == 0
}
