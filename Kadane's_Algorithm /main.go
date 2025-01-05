package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Reading array from file
	input := readFile()

	n, _ := strconv.Atoi(input[0]) // Length of the array
	arr := make([]int, n)
	array := strings.Split(input[1], " ")
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(array[i])
	}
	start := time.Now()
	Sequential(arr)
	elapsed := time.Since(start)
	fmt.Printf("Linear: %s\n", elapsed)

	start = time.Now()
	EatEachSideApproach(arr)
	elapsed = time.Since(start)
	fmt.Printf("Eat from each side: %s\n", elapsed)
}
