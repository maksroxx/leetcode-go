package subarraysum

func subarraySum(nums []int, k int) int {
	subs := make(map[int]int)
	subs[0] = 1
	count := 0
	curr := 0
	for _, num := range nums {
		curr += num
		if value, exist := subs[curr-k]; exist {
			count += value
		}
		subs[curr]++
	}
	return count
}
