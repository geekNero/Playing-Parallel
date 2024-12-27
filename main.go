package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	stringTree := readFile()
	var search int
	fmt.Scan(&search)
	n, _ := strconv.Atoi(stringTree[0])
	tree := make(map[int][]int)
	// Storing the tree as a map where parents are keys and children are values
	for i := 1; i < n; i++ {
		edge := strings.Split(stringTree[i], " ")
		node, _ := strconv.Atoi(edge[0])
		adj, _ := strconv.Atoi(edge[1])
		tree[node] = append(tree[node], adj)
	}
	start := time.Now()
	linearSearch(&tree, search)
	elapsed := time.Since(start)
	fmt.Printf("Linear: %s\n", elapsed)

	start = time.Now()
	parallelSearch(&tree, search)
	elapsed = time.Since(start)
	fmt.Printf("Concurrent: %s\n", elapsed)

	start = time.Now()
	lazyThread(&tree, search)
	elapsed = time.Since(start)
	fmt.Printf("Lazy: %s\n", elapsed)
}
