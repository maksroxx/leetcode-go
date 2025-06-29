package longestsubarray

func longestSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	maxLength := 0
	count := 0
	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}
		for count > 1 {
			if nums[left] == 0 {
				count--
			}
			left++
		}
		maxLength = max(maxLength, right-left)
		right++
	}
	return maxLength
}
