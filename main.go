package main

import (
	"fmt"
	"strconv"
)

// Simple usage example that inserts the numbers 0, 1, 2 into a queue and then
// removes them one by one, printing them to the standard output.
func main() {
	// Create a queue an push some data in
	q := New()
	for i := 1; i < 6; i++ {
		q.Push(strconv.Itoa(i))
	}

	fmt.Println("LPOP:", q.Pop())
	// Pop out the queue contents and display them
	for !q.Empty() {
		fmt.Println(q.Pop())
	}
	// Output:
	// 0
	// 1
	// 2
}
