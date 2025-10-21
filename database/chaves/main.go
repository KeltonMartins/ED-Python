package main

import (
	"fmt"
)

func main() {
	gt1, gt2 := 0, 0
	queue := NewQueue[string]()
	for i := range 16 {
		queue.Enqueue(string('A' + i))
	}

	for queue.items.Len() > 1 {
		fmt.Scan(&gt1, &gt2)
		if gt1 > gt2 {
			volta := queue.Dequeue()
			queue.Enqueue(volta)
			queue.Dequeue()
		} else {
			queue.Dequeue()
			volta := queue.Dequeue()
			queue.Enqueue(volta)
		}
	}
	fmt.Println(queue)
}
