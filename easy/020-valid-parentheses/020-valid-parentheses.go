package validparentheses

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '{', '[', '(':
			stack = append(stack, c)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}
