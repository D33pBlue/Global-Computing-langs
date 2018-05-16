package main

import (
  "./ctm"
  "./cs"
)

func main(){
  machine := ctm.New()
  machine.SetPrice("Tea",2)
  machine.SetPrice("Coffee",3)
  go machine.Run()
  cs1 := cs.New("1")
  go cs1.Run(machine)
  cs2 := cs.New("2")
  cs2.Run(machine)
}
