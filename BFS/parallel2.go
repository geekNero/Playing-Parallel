package main

import (
	"container/list"
	"fmt"
	"sync/atomic"
)

func masterBFS(tree *map[int][]int, queue *list.List, find int) bool {
	ch := make(chan bool, 1)
	slave_queue := list.New()
	thread_free := int32(0)
	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		cur := current.Value.(int)
		slave_status := atomic.LoadInt32(&thread_free)
		if slave_status == 0 && len((*tree)[cur]) > 20 {
			slave_queue.PushBack(cur)
			atomic.StoreInt32(&thread_free, -1)
			go slaveBFS(tree, slave_queue, find, ch, &thread_free)
			continue
		} else if slave_status == 1 {
			return true
		}
		for _, v := range (*tree)[cur] {
			// time.Sleep(500 * time.Millisecond)
			if v == find {
				return true
			}
			queue.PushBack(v)
		}
	}
	if atomic.LoadInt32(&thread_free) == 0 {
		return false
	}
	return <-ch
}
func slaveBFS(tree *map[int][]int, queue *list.List, find int, in chan bool, out *int32) {
	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		cur := current.Value.(int)
		for _, v := range (*tree)[cur] {
			// time.Sleep(500 * time.Millisecond)
			if v == find {
				atomic.StoreInt32(out, 1)
				in <- true
				return
			}
			queue.PushBack(v)
		}
	}
	atomic.StoreInt32(out, 0)
	in <- false
}

func lazyThread(tree *map[int][]int, search int) {
	queue := list.New()
	for _, v := range (*tree)[1] {
		queue.PushBack(v)
	}
	fmt.Println(masterBFS(tree, queue, search))
}
