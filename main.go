
package main

import "fmt"
import "time"


func producer(ch chan int){

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        // TODO: push real value to buffer
        ch <- i
    }
    close(ch)
}

func consumer(ch chan int){

    time.Sleep(1 * time.Second)
    for i := range ch { //TODO: get real value from buffer
        fmt.Printf("[consumer]: %d\n", i)
        time.Sleep(50 * time.Millisecond)
    }
    
}


func main(){
    
    // TODO: make a bounded buffer
    buffer := make(chan int, 5)
    
    go consumer(buffer)
    go producer(buffer)
    
    select {}
}