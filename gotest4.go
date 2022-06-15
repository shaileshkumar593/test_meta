/*Q4

 */

package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m chan int) {
	_ = <-m

	x = x + 1

	wg.Done()
}

// channel in go routine is blocking call. So it will block go rotine to be accessed by other
// go routine.
func main() {
	var w sync.WaitGroup
	var c chan int = make(chan int)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		c <- i
		go increment(&w, c)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
