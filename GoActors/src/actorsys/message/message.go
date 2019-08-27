package message

type Anything interface{}

type Message struct{
  Value Anything
  From int64
  To int64
}

func New(from,to int64,value Anything) *Message{
  msg := new(Message)
  msg.From = from
  msg.To = to
  msg.Value = value
  return msg
}
