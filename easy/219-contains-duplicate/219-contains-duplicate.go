package containsduplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, exist := seen[nums[i]]; exist {
			n := i - index
			if n < 0 {
				n *= -1
			}
			if n <= k {
				return true
			}
		}
		seen[nums[i]] = i
	}
	return false
}
