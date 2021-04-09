package main

func maxProfit(prices []int) int {
	// edge cases, 0 or 1 trade
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	b := 0
	mp := prices[1] - prices[b]
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[b] {
			b = i
			continue
		}
		if prices[i] - prices[b] > mp {
			mp = prices[i] - prices[b]
		}
	}

	if mp < 0 {
		return 0
	}
	return mp
}
