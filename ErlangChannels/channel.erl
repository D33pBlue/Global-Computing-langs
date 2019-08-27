-module(channel).
-export([ch_make/0,ch_send/2,ch_receive/1,ch_close/1,channel/2]).

channel(In,Out)->
  receive
    {req_send,Pid,Val} ->
      case Out of
        []          -> channel(In++[{Pid,Val}],Out);
        [Recv|TOut] -> Recv ! {chval,Val},
                      Pid ! send_ack,
                      channel(In,TOut)
      end;
    {req_receive,Pid} ->
      case In of
        []              -> channel(In,Out++[Pid]);
        [{Snd,Val}|TIn] -> Pid ! {chval,Val},
                          Snd ! send_ack,
                          channel(TIn,Out)
      end;
    stop -> io:format("~30p~n",[{closing,self()}])
  end.

ch_make() ->
  spawn(channel,channel,[[],[]]).

ch_close(Channel) ->
  Channel ! stop.

ch_send(Channel,Val) ->
  Channel ! {req_send,self(),Val},
  receive
    send_ack -> ok
  end.

ch_receive(Channel) ->
  Channel ! {req_receive,self()},
  receive
    {chval,Value} -> Value
  end.
