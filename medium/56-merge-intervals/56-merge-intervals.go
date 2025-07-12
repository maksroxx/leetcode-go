package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	curr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= curr[1] {
			curr[1] = max(curr[1], intervals[i][1])
		} else {
			result = append(result, curr)
			curr = intervals[i]
		}
	}
	result = append(result, curr)
	return result
}
