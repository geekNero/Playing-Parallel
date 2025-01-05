package main

import "fmt"

func kadane(arr []int, start int, end int, add int, ch chan int) {

	maxSum := 0
	currentSum := 0
	for i := start; i != end; i += add {
		// Implement Kadane's algorithm here
		currentSum += arr[i]
		if currentSum < 0 {
			currentSum = 0
		} else if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	ch <- maxSum
	ch <- currentSum
}

func EatEachSideApproach(arr []int) int {
	n := len(arr)
	ch1 := make(chan int)
	ch2 := make(chan int)
	middle := n / 2
	go kadane(arr, 0, middle, 1, ch1)
	go kadane(arr, n-1, middle-1, -1, ch2)
	leftSumMax := <-ch1
	leftSumMiddle := <-ch1
	rightSumMax := <-ch2
	rightSumMiddle := <-ch2
	fmt.Println(max(leftSumMax, rightSumMax, leftSumMiddle+rightSumMiddle))
	return 0
}
