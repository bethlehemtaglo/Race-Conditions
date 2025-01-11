// Use `go run foo.go` to run your program
//this was under shared variables 'go'

package main

import (
    . "fmt"
    "runtime"
    "time"
    "sync"
)

var i = 0
var mu sync.Mutex



func incrementing() {
    //TODO: increment i 1000000 times
    for j:=0; j < 1000000; j++ {
        mu.Lock()
        i++ 
        mu.Unlock()
    }
}

func decrementing() {
    //TODO: decrement i 1000000 times
    for j:=0; j < 1000000; j++ {
        mu.Lock()
        i--
        mu.Unlock()
    }
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
        // GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting. 
        // It defaults to the value of [runtime.NumCPU]. If n < 1, it does not change the current setting.
        // This call will go away when the scheduler improves. It can therefore be used to limit or increase concurrency
    runtime.GOMAXPROCS(2)    
	
    // TODO: Spawn both functions as goroutines
    
    go incrementing()
    go decrementing()

	
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
