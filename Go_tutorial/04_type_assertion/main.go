package main

import (
  "fmt"
  "./stack"
)

type MyInt int

func (self MyInt) Triple() MyInt{
  return 3*self
}

func main()  {
  fmt.Println("start")
  var s stack.Stack
  var m MyInt = 3
  s.Push(3)
  s.Push(m)
  for i := 0; i < len(s); i++ {
    //fmt.Println(s[i].Triple()) error!
    // need type assertion
    if x,ok := s[i].(MyInt); ok {
      fmt.Printf("here is my int: %d\n",x.Triple())
    } else{
      fmt.Println("not MyInt")
    }
  }
}
