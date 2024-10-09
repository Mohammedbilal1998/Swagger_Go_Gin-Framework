// https://medium.com/@ninucium/go-interview-questions-part-1-pointers-channels-and-range-67c61345cf3c
package main

import (
	"fmt"
	"sync"
)
func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4,5,6,7,8,9,10}
	wg := sync.WaitGroup{}
	// Add the number of works equal to the number of array elements.
	wg.Add(len(array))
	go func() {
	  for _, value := range array {
		ch <- &value
	  }
	}()
	go func() {
	  for value := range ch {
		fmt.Println(*value)
		// Decrement the waitgroup counter with each iteration.
		wg.Done()
	  }
	}()
  
	wg.Wait()
  }