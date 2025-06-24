package happynumber

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	seen := make(map[int]bool)
	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		n = findSum(n)
	}
	return true
}

func findSum(number int) int {
	sum := 0
	for number > 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	return sum
}
