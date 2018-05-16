package cs

import (
  "fmt"
  "math/rand"
  "time"
  "../ctm"
)

type CS struct {
  name string
  Money int
}

func New(name string) *CS {
  return &CS{name,0}
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
      fmt.Printf("CS%s: drink %s\n",self.name,beverage)
      m = <-machine.Change
    case m = <-machine.Change:
    }
    self.Money = m
    fmt.Printf("CS%s: got back %d money\n",self.name,m)
  }
}
