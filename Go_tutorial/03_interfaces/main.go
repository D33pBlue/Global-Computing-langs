package main

import (
  "fmt"
  "./abser"
  "./vertex"
)
func main (){
  var v abser.Abser
  v = vertex.New(3,4) // Vertex is an Abser
  // v = *vertex.New(3,4) gives error!
  fmt.Println(v.Abs())
}
