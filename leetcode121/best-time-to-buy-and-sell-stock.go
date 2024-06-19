package main

import "fmt"

func profit(arr []int) int {
	l := arr[0]
	r := arr[1]
	maxProfit := 0
	for r < len(arr) {
		if arr[l] < arr[r] {
			profit := arr[r] - arr[l]
			maxProfit = max(maxProfit, profit)
		} else {
			l = r
		}
		r += 1
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 2, 3, 4, 6, 2}
	fmt.Println(profit(prices))
}
