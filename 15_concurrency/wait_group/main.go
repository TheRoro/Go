package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	//anonymous function, to delegate parameters that are not the responsibility of the count function
	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
