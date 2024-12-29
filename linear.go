package main

import (
	"container/list"
	"fmt"
)

func BFS(tree *map[int][]int, node int, find int) bool {

	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		current := queue.Front()
		cur := current.Value.(int)
		queue.Remove(current)
		for _, v := range (*tree)[cur] {
			// time.Sleep(500 * time.Millisecond)
			if v == find {
				return true
			}
			queue.PushBack(v)
		}
	}
	return false
}

func linearSearch(tree *map[int][]int, search int) {
	val := BFS(tree, 1, search)
	fmt.Println(val)
}
