package main

import (
  "fmt"
  "math/rand"
)

func Producer(ch chan<- int){
  for {
    n := rand.Intn(100)
    ch <- n
    fmt.Printf("Produced value\n")
  }
}

func Consumer(ch <-chan int){
  for {
    var n int = <- ch
    fmt.Printf("Received value %d\n",n)
  }
}

func main(){
  const NProd = 10
  const NCons = 3
  // Async channel with buffer
  channel := make(chan int,100)
  for i:=0; i<NCons; i++ {
    go Consumer(channel)
  }
  for i:=0; i<NProd-1; i++ {
    go Producer(channel)
  }
  Producer(channel)
}
