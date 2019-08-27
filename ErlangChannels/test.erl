-module(test).
-export([main/0,producer/1,consumer/1]).

-import(channel,[ch_make/0,ch_send/2,ch_receive/1,ch_close/1,channel/2]).
-import(timer,[sleep/1]).
-import(rand,[uniform/1]).


producer(Ch) ->
  Val = uniform(100),
  io:format("~30p~n",[{self(),produced,Val}]),
  ch_send(Ch,Val),
  io:format("~30p~n",[{self(),sent,Val}]),
  producer(Ch).

consumer(Ch) ->
  sleep(4000),
  Val = ch_receive(Ch),
  io:format("~30p~n",[{self(),received,Val}]),
  consumer(Ch).

instantiate(Fun,Ch,1) ->
  spawn(test,Fun,[Ch]);

instantiate(Fun,Ch,N) ->
  spawn(test,Fun,[Ch]),
  instantiate(Fun,Ch,N-1).

main()->
  Ch = ch_make(),
  instantiate(consumer,Ch,1),
  instantiate(producer,Ch,3).
