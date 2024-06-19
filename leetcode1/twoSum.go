package main

import "fmt"

func twoSum(arr []int, target int) []int {
	valMap := make(map[int]int)

	// create a map such that it holds each value with its respective index
	// append only the values which on not finding the diff in the map
	// if found the diff in the map return an arr[arrayIdx[value],map[diffValue]]
	for i, v := range arr {
		diff := target - v
		if idx, ok := valMap[diff]; ok {
			return []int{idx, i}
		} else {
			valMap[v] = 1
		}
	}
	return []int{}
}

func main() {
	arr := []int{2, 1, 5, 3}
	ans := twoSum(arr, 4)
	fmt.Println(ans)
}
