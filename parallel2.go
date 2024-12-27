package main

import (
	"container/list"
	"fmt"
)

func ifResultFound(c chan bool) int {
	select {
	case val := <-c:
		if val {
			return 1
		} else {
			return 0
		}
	default:
		return -1
	}
}

func masterBFS(tree *map[int][]int, queue *list.List, find int) bool {
	ch := make(chan bool, 1)
	ch <- false
	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		cur := current.Value.(int)
		if ifResultFound(ch) == 0 {
			go slaveBFS(tree, cur, find, ch)
			continue
		} else if ifResultFound(ch) == 1 {
			return true
		}
		for _, v := range (*tree)[cur] {
			if v == find {
				return true
			}
			queue.PushBack(v)
		}
	}
	return <-ch
}
func slaveBFS(tree *map[int][]int, node int, find int, in chan bool) {
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		cur := current.Value.(int)
		for _, v := range (*tree)[cur] {
			if v == find {
				in <- true
				return
			}
			queue.PushBack(v)
		}
	}
	in <- false
}

func lazyThread(tree *map[int][]int, search int) {
	queue := list.New()
	for _, v := range (*tree)[1] {
		queue.PushBack(v)
	}
	fmt.Println(masterBFS(tree, queue, search))
}
