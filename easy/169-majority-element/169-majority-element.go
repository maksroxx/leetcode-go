package majorityelement

func majorityElement(nums []int) int {
	var cand int
	count := 0

	for _, num := range nums {
		if count == 0 {
			cand = num
		}
		if num == cand {
			count++
		} else {
			count--
		}
	}
	return cand
}
