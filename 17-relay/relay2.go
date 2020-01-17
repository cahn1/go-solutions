package main

import "fmt"
import "time"

// Add a bool channel for runner() to tell main() it is done
var done = make(chan bool)

func runner(baton chan int) {
	num := <-baton // wait to find out which runner I am
	fmt.Println("Runner", num, "running with the baton")

	if num < 4 {
		time.Sleep(1 * time.Second)
		fmt.Println("\tRunner", num+1, "moves to the starting line")
		go runner(baton)
	}
	time.Sleep(5 * time.Second) // this is me running
	if num < 4 {
		fmt.Println("\tRunner", num, "passing baton to runner", num+1)
		time.Sleep(1 * time.Second)
		baton <- num + 1
	} else {
		fmt.Println("Race over!")
		done <- true // tell main() we are done
	}
}

func main() {
	baton := make(chan int) // create channel
	go runner(baton)        // start the first runner
	baton <- 1              // send "1" to runner
	<-done                  // wait for last runner()
}
