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
  scientist := cs.New()
  scientist.Run(machine)
}
