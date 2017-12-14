package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// Create a queue an push some data in
	q := New()
	var wg sync.WaitGroup

	for i := 1; i < 6; i++ {
		wg.Add(1)
		go push(q, &wg)
	}

	// fmt.Println("LPOP:", q.Pop())
	// Pop out the queue contents and display them
	var lwg sync.WaitGroup
	for i := 1; i < 6; i++ {
		lwg.Add(1)
		go lpop(q, &lwg)
	}

	wg.Wait()
	lwg.Wait()
}

func push(q *Queue, wg *sync.WaitGroup) {
	for i := 1; i < 6; i++ {
		q.Push(strconv.Itoa(i))
	}
	wg.Done()
}

func lpop(q *Queue, lwg *sync.WaitGroup) {
	fmt.Println(q.Pop())
	lwg.Done()
}
