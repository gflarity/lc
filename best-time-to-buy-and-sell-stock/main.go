package main

func maxProfitBrute(prices []int) int {
	// brute force calculate the delta between each pair that is allowe => n^2

	// edge cases, 0 or 1 trade
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	// track the max profits
	mp := 0

	// outer loop tracks potential sell, inner tracks potential buy
	for i := len(prices) - 1; i > -1; i-- {
		for j := i - 1; j > -1; j-- {
			lp := prices[i] - prices[j]
			if lp > mp {
				mp = lp
			}
		}
	}

	return mp

}

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
