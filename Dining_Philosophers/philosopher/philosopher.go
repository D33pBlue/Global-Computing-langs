package philosopher

import (
  "fmt"
  "time"
  "math/rand"
  "../fork"
)

type Philosopher struct{
  name string
  left,right *fork.Fork
}

func New(name string,left,right *fork.Fork) *Philosopher  {
  return &Philosopher{name,left,right}
}

func (self *Philosopher) Run(){
  for{
    fmt.Printf("Philosopher %s is thinking",self.name)
    time.Sleep(time.Duration(rand.Int63n(2*1e9)))
    self.left.Take <- 1
    time.Sleep(time.Duration(rand.Int63n(2*1e9)))
    self.right.Take <- 1
    fmt.Printf("Philosopher %s is eating",self.name)
    time.Sleep(time.Duration(rand.Int63n(2*1e9)))
    self.left.Leave <- 1
    self.right.Leave <- 1
  }
}
