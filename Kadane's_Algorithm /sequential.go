package main

import "fmt"

func Sequential(arr []int) {
	maxSum := 0
	currentSum := 0
	for i := 0; i < len(arr); i++ {
		// Implement Kadane's algorithm here
		currentSum += arr[i]
		if currentSum < 0 {
			currentSum = 0
		} else if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	fmt.Println(maxSum)

}
