package fork

import "fmt"

type Fork struct{
  name string
  Take chan int
  Leave chan int
}

func New(n string) *Fork{
  return &Fork{n,make(chan int),make(chan int)}
}

func (self *Fork) Run(){
  for{
    <-self.Take
    fmt.Printf("Fork %s taken\n",self.name)
    <-self.Leave
    fmt.Printf("Fork %s left\n",self.name)
  }
}
