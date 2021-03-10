package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		//sending and receiving are blocking operations, they wait until they are sync
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
