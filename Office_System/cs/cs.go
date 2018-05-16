package cs

import (
  "fmt"
  "math/rand"
  "time"
  "../ctm"
)

type CS struct {
  Money int
}

func New() *CS {
  return &CS{0}
}

func (self *CS)Run(machine *ctm.CTM){
  for {
    // publish and gain some money
    gain := rand.Intn(5)
    self.Money += gain
    machine.Pay <- self.Money
    time.Sleep(time.Duration(rand.Intn(5))*time.Second)
    select {
    case machine.Tea <- 1:
    case machine.Coffee <- 1:
    case <-machine.TOut:
    }
    var m int
    select {
    case beverage := <- machine.Drink:
      fmt.Printf("CS: drink %s\n",beverage)
      m = <-machine.Change
    case m = <-machine.Change:
    }
    self.Money = m
    fmt.Printf("CS: got back %d money\n",m)
  }
}
