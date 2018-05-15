package main

import (
  "fmt"
  "./vertex"
  "./coloredvertex"
)

func main(){
  v := vertex.New(3,4)
  fmt.Println(v.Abs())
  v2 := coloredvertex.New(3,4,7)
  fmt.Println(v2.X)
  fmt.Println(v2.Abs())
  fmt.Println(v2.Color)
}
