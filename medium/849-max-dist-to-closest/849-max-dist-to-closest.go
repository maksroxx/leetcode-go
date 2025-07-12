package maxdisttoclosest

func maxDistToClosest(seats []int) int {

	maxDist := 0
	lastOcupired := -1

	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if lastOcupired != -1 {
				dist := (i - lastOcupired) / 2
				maxDist = max(dist, maxDist)
			} else {
				maxDist = i
			}
			lastOcupired = i
		}
	}
	maxDist = max(maxDist, len(seats)-1-lastOcupired)
	return maxDist
}
