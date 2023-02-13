package main

func main() {

}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sell, buy, prevSell, prevBuy := 0, -prices[0], 0, 0
	for _, price := range prices {
		prevBuy = buy
		buy = max(prevSell-price, prevBuy)
		prevSell = sell
		sell = max(prevBuy+price, prevSell)
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
