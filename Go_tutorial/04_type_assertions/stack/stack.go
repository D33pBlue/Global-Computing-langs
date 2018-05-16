package stack

type Stack [] interface{}

// There are no generics in go.
// Use interface{} as base type
// to build multi-type containers
// and then check the type before
// use with type assertions.

func (self *Stack) Push(x interface{}){
  *self = append(*self,x)
}
