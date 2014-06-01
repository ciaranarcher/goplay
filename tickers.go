// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import "time"
import "fmt"
import "math/rand"

var (
	avail_toggle = 9 // 2 == easily found, 9 == not going to be found often
)

func main() {

	c := make(chan bool)
	rand.Seed(time.Now().UTC().UnixNano())

	ticker := time.NewTicker(time.Millisecond * 500)
	go func(chan bool) {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			if checkAvail() {
				c <- true
			}
		}
	}(c)

	select {
	case <-c:
		ticker.Stop()
		fmt.Println("Agent avail!")
	case <-time.After(time.Second * 5):
		fmt.Println("Too long on queue!")
	}
}

func checkAvail() bool {
	if rand.Intn(100)%avail_toggle == 0 { //
		return true
	} else {
		return false
	}
}
