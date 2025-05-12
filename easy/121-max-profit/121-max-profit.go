package maxprofit

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
