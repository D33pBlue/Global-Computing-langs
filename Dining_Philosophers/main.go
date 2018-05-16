package main

import (
  "fmt"
  "./fork"
  "./philosopher"
)

func main(){
  const NPhil = 5
  var forks [NPhil] fork.Fork
  var phils [NPhil] philosopher.Philosopher
  // build and run Forks
  for i:=0; i<NPhil; i++ {
    forks[i] = *fork.New(fmt.Sprintf("F%d",i))
    go forks[i].Run()
  }
  // build and run Philosophers
  for i:=0; i<NPhil-1; i++ {
    phils[i] = *philosopher.New(fmt.Sprintf("P%d",i),&forks[i],&forks[i+1])
    go phils[i].Run()
  }
  // last Philosopher is left handed
  // => so there is no deadlock!
  phils[NPhil-1] = *philosopher.New(
    fmt.Sprintf("P%d",NPhil-1),
    // with all Philosophers right handed there is deadlock
    // &forks[NPhil-1],
    // &forks[0])
    &forks[0],
    &forks[NPhil-1])
  phils[NPhil-1].Run()
}
