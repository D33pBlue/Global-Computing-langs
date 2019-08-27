package simpleactor

import(
  "fmt"
  "time"
  "actorsys/actor"
)

type SimpleActor struct{
  Name string
  actor.Actor
}

func New(sys *actor.Actorsys,name string) *SimpleActor{
  return &SimpleActor{name,*actor.New(sys)}
}

func (self *SimpleActor) Act(other int64){
  for{
    self.Send(other,"Hello from "+self.Name)
    if m,ok := self.Receive().(string); ok{
      fmt.Println(self.Name+" received: \""+m+"\"")
    }
    time.Sleep(1*time.Second)
  }
}
