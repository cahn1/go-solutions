package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Create a global (shared) WaitGroup
var wg sync.WaitGroup

func sleeper(num int) {
	// Defer the call to Done, so it will happen before we return
	// and in the even we panic(), we still call wg.Done()
	defer wg.Done()
	// Each sleeper will sleep for between 5 and 30 seconds
	secs := rand.Int()%26 + 5
	fmt.Printf("Sleeper %d sleeping for %d seconds\n", num, secs)
	time.Sleep(time.Duration(secs) * time.Second)
	fmt.Printf("Sleeper %d done!\n", num)
}

func main() {
	for i := 1; i < 25_000; i++ {
		wg.Add(1)
		go sleeper(i)
	}
	wg.Wait()
}
