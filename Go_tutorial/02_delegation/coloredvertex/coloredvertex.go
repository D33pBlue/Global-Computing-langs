package coloredvertex

import "../vertex"

type ColoredVertex struct{
  vertex.Vertex // embedding
  Color int
}

func New (x,y float64,color int) *ColoredVertex{
  // delegation
  return &ColoredVertex{*vertex.New(x,y),color}
}
