/*
Q2
*/

package main

import (
	"fmt"
)

func intIter(c chan int) {
	Seq := [...]int{10, 20, 35, 100, 200, 502}

	for i := 0; i < len(Seq); i++ {
		c <- Seq[i]
	}
	close(c)
}

func main() {
	var c chan int = make(chan int)

	go intIter(c)

	for val := range c {
		fmt.Println(val)
	}
}
