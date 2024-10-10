// https://medium.com/@ninucium/golang-concurrency-patterns-for-select-done-errgroup-and-worker-pool-645bec0bd3c9
package main

import(
	"fmt"
	"context"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func someTask (){
	fmt.Println("some task got called")
}

func periodicTask(ctx context.Context){
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			someTask()
		case <- ctx.Done():
			fmt.Println("context got cancelled")
			ticker.Stop()
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()
	go periodicTask(ctx)
	sys := make(chan os.Signal, 1)
	signal.Notify(sys, syscall.SIGTERM)
	<-sys
	
}