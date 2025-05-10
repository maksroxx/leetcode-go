package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		comp := target - num
		if j, ok := m[comp]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
