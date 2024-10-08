package main

import(
	"fmt"
	"sync"
)

func main(){
	var (
		lock sync.Mutex 
		math sync.WaitGroup
	)
		data := 0

		increment := func() {
			lock.Lock()
			defer lock.Unlock()
			data++
			fmt.Printf("Incrementing: %d\n", data)
		}

		decrement := func() {
			lock.Lock()
			defer lock.Unlock()
			data--
			fmt.Printf("Incrementing: %d\n", data)
		}

		for i:=0; i<5; i++ {
			math.Add(1)
			func() {
				defer math.Done()
				go increment()
			}()
		}

		for i:=0; i<5; i++ {
			math.Add(1)
			func() {
				defer math.Done()
				go decrement()
			}()
		}

		math.Wait()
		fmt.Printf("work Done: %d\n", data)


}