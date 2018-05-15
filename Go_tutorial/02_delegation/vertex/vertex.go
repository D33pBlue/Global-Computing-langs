package vertex

import (
  "math"
)

type Vertex struct{
  X,Y float64
}

func (self *Vertex) Abs() float64 {
  return math.Sqrt(self.X*self.X+self.Y*self.Y)
}

func New (x,y float64) *Vertex{
  return &Vertex{x,y}
}
