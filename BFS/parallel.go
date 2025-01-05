package main

import (
	"container/list"
	"fmt"
	"sync"
)

// When receiving from main thread
// -1 = no more data, stop computing
// 0 = data inserted, continue computing
// 1 = second thread found value, stop computing

// when sending to main thread
// 0 = done computing, feed more
// 1 = value found

type Queue struct {
	mu    sync.Mutex
	value *list.List
}

func (q *Queue) SplitHalf(q1 *Queue) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q1.mu.Lock()
	defer q1.mu.Unlock()
	for i := 0; i < q.value.Len(); i += 2 {
		q1.PushBack(q.value.Front().Value.(int))
		q.value.Remove(q.value.Front())
	}
}

func (q *Queue) Len() int {
	// q.mu.Lock()
	// defer q.mu.Unlock()
	return q.value.Len()
}

func (q *Queue) Remove(val *list.Element) {
	// q.mu.Lock()
	// defer q.mu.Unlock()
	q.value.Remove(val)
}

func NewQueue() *Queue {
	return &Queue{
		value: list.New(),
	}
}

func (q *Queue) PushBack(val int) {
	// q.mu.Lock()
	// defer q.mu.Unlock()
	q.value.PushBack(val)
}

func (q *Queue) Front() *list.Element {
	// q.mu.Lock()
	// defer q.mu.Unlock()
	return q.value.Front()
}

func ifStop(c chan int) bool {
	// Check if the main thread sent a signal to stop computing
	select {
	case val := <-c:
		if val == 1 || val == -1 {
			return true
		}
	default:
		return false
	}
	return false
}

func BFSParallel(tree *map[int][]int, node int, find int, in chan int) int {

	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		if ifStop(in) {
			// Doing this to minimize extra computation
			return 2
		}
		current := queue.Front()
		cur := current.Value.(int)
		queue.Remove(current)
		for _, v := range (*tree)[cur] {
			// time.Sleep(500 * time.Millisecond)
			if v == find {
				return 1
			}
			queue.PushBack(v)
		}
	}
	return 0
}

func thread(q *Queue, in chan int, out chan int, tree *map[int][]int, find int) {

	// Checking if the current queue consist of the find value
	for {
		index := 0
		for q.Len() > index {
			val := q.Front()

			v := (*val).Value.(int)
			if v == find {
				out <- 1
				return
			}
			q.Remove(val)
			q.PushBack(v)
			index++
		}

		// BFSing through the queue
		for q.Len() > 0 {
			if ifStop(in) {
				return
			}
			val := q.Front()
			v := (*val).Value.(int)
			ans := BFSParallel(tree, v, find, in)
			if ans == 2 {
				return
			} else if ans == 1 {
				out <- 1
				return
			}
			q.Remove(val)
		}
		out <- 0
		ifContinue := <-in
		if ifContinue == -1 || ifContinue == 1 {
			return
		}
	}
}

func parallelSearch(tree *map[int][]int, search int) {
	// Implement parallel search here
	q1 := NewQueue()
	q2 := NewQueue()
	// keeping buffer of 1 so that the goroutines can handle interrupts, because
	// the main thread will be sending signals to stop computing and the goroutine will
	// be sending signals to feed more data, so the buffer of 1 will allow the goroutine to
	// move on to check if the main thread sent something
	c1main := make(chan int, 1)
	c1thread := make(chan int, 1)
	c2main := make(chan int, 1)
	c2thread := make(chan int, 1)
	defer close(c1main)
	defer close(c2main)
	defer close(c1thread)
	defer close(c2thread)

	// Splitting the data between the two queues
	for i, v := range (*tree)[1] {
		if i%2 == 0 {
			q1.PushBack(v)
		} else {
			q2.PushBack(v)
		}
	}

	go thread(q1, c1main, c1thread, tree, search)
	go thread(q2, c2main, c2thread, tree, search)

	for {
		select {
		// if the first thread returned something
		case val1 := <-c1thread:
			if val1 == 1 {
				// If it found the value then notify the second thread to stop
				c2main <- 1
				fmt.Println(true)
				return
			} else if val1 == 0 { // If it didn't find the value
				select {
				// Check if the second thread found the value
				case val2 := <-c2thread:
					if val2 == 1 {
						// If the second thread found the value then stop the first thread
						c1main <- 1
						fmt.Println(true)
						return
					} else if val2 == 0 {
						// If the second thread didn't find the value then print false
						fmt.Println(false)
						c1main <- -1
						return
					}
				default:
					// If the second thread didn't finish computing then split the data between the two threads
					q2.SplitHalf(q1)
					c1main <- 0
				}
			}

		case val2 := <-c2thread:
			// Similar to the first case but for the second thread
			if val2 == 1 {
				c1main <- 1
				fmt.Println(true)
				return
			} else if val2 == 0 {
				select {
				case val1 := <-c1thread:
					if val1 == 1 {
						c2main <- 1
						fmt.Println(true)
						return
					} else if val1 == 0 {
						fmt.Println(false)
						c2main <- -1
						return
					}
				default:
					q1.SplitHalf(q2)
					c2main <- 0
				}
			}
		}
	}
}
