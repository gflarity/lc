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

	// track the max profits
	mp := 0

	return mp
}

func main() {
	//	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
