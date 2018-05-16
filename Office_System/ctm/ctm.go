package ctm

import (
  "fmt"
  "time"
)

type CTM struct{
  Pay,Coffee,Tea,TOut,Change chan int
  Drink chan string
  Price map[string]int
  Gain int
}

func (self *CTM) SetPrice(item string,price int){
  self.Price[item] = price
}

func New() *CTM{
  return &CTM{
    make(chan int),
    make(chan int),
    make(chan int),
    make(chan int),
    make(chan int),
    make(chan string),
    make(map[string]int),
    0}
}

func (self *CTM) Run(){
  var choice string
  var money int
  for{
    money = <-self.Pay
    select {
    case <-self.Tea:
      choice = "Tea"
    case <-self.Coffee:
      choice = "Coffee"
    case <-time.After(5*time.Second):
      choice = "nothing"
      fmt.Println("Time out")
      self.TOut <- 1
      }
    if choice!="nothing" {
      fmt.Printf("CTM: select %s\n",choice)
      if self.Price[choice]<=money {
        self.Drink <- choice
        self.Gain += self.Price[choice]
        money = money-self.Price[choice]
      }else{
        fmt.Printf("CTM: %d money not enough\n",money)
      }
    }
    self.Change <- money //give back the Change
    fmt.Printf("CTM: current balance %d money\n",self.Gain)
  }
}
