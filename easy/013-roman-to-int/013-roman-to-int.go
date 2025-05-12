package romantoint

func romanToInt(s string) int {
	f := func(c byte) int {
		switch c {
		case 'I':
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		default:
			return 0
		}
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		value := f(s[i])
		if i < len(s)-1 && value < f(s[i+1]) {
			sum -= value
		} else {
			sum += value
		}
	}
	return sum
}
