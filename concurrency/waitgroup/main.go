package main

import(
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	strs := []string {"apple", "ball", "cat", "dog"}

	for it, i := range (strs){
		wg.Add(1)
		go func(it int, s string){
			defer wg.Done()
			fmt.Printf("goroutine %v value is: %s \n",it,  i)
		}(it, i)
	}
	wg.Wait()
}