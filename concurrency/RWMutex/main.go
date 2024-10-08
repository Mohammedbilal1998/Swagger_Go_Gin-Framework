package main

import (
	"fmt"
	"sync"
)

func main() {
	// l := &sync.Mutex{}
	l := &sync.RWMutex{}
	m := make(map[int] int, 100)

	go reader(l, m)
	go writer(l, m)
	go reader(l, m)

	go reader(l, m)
	go reader(l, m)
	go reader(l, m)
	go reader(l, m)
	go reader(l, m)

	ch := make(chan struct{})
	<-ch
}

func reader(l *sync.RWMutex, m map[int]int) {
	for{
		l.Lock()
		for key, value := range(m) {
			fmt.Printf("key :%v, value :%v\n", key, value)
		}
		l.Unlock()
	}
}

func writer(l *sync.RWMutex, m map[int]int) {
	for{
		
		for i:=0; i<100; i++ {
			l.Lock()
			m[i]=i
			l.Unlock()
		}
		
	}
}