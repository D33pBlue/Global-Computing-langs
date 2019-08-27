package main

import(
  "actorsys/actor"
  "actorsys/simpleactor"
)

func main(){
  sys := actor.StartActorSys()
  a := simpleactor.New(sys,"Pippo")
  b := simpleactor.New(sys,"Pluto")
  go a.Act(b.GetAddress())
  b.Act(a.GetAddress())
}
