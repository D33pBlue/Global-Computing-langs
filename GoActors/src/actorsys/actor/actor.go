package actor

import(
  "actorsys/message"
)

type Actor struct{
  Id int64
  Ch chan message.Message
  sys chan message.Message
}

func (self *Actor) Receive() message.Anything{
  msg := <-self.Ch
  return msg.Value
}

func (self *Actor) Send(actorId int64,value message.Anything){
  msg := message.New(self.Id,actorId,value)
  self.sys <- *msg
}

func(self *Actor) GetAddress()int64{
  return self.Id
}

type Actorsys struct{
  nextID int64
  actors map[int64] chan message.Message
  incoming chan message.Message
}

func (self *Actorsys) run(){
  for{
    msg := <-self.incoming
    if ch,ok := self.actors[msg.To]; ok {
      ch <- msg
    }
  }
}

func StartActorSys() *Actorsys{
  s := new(Actorsys)
  s.nextID = 0
  s.actors = make(map[int64] chan message.Message)
  s.incoming = make(chan message.Message,500)
  go s.run()
  return s
}

func (self *Actorsys) newActor() *Actor{
  a := new(Actor)
  a.Id = self.nextID
  self.nextID = self.nextID+1
  a.Ch = make(chan message.Message,100)
  self.actors[a.Id] = a.Ch
  a.sys = self.incoming
  return a
}

func New(sys *Actorsys) *Actor{
  return sys.newActor()
}
