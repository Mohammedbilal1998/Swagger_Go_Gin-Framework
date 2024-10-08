package main

import(
	"fmt"
)

func main(){
	data := 0

	go func() {
		data++
	}()
	if(data == 0){
		fmt.Println(data)
	}
}