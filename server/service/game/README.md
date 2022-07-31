# player actor(由game管理的本地actor)
# 1.player actor通信

```
player设计为本地actor不走服务发现，player可以通过request、requestWait、Send方式和远端的服务通信，
远端服务只能通过向玩家所在game通信并由game将消息转发给玩家，途中game会判断玩家是否还在内存，不在内存的
玩家需要重新激活player actor
```

## 2.player actor退出

```
退出actor，会关闭actor占用的内存以及事件循环使用的chan和goroutine
满足以下所有条件的玩家，player actor退出：
1.长时间未处理消息
2.长时间未上线
```

## 3.player actor激活

```
通过给玩家发送消息，由game触发激活玩家actor，并将消息投递到所属邮箱
```

## FAQ:
### 1.为什么player设计为actor，而不是所有玩家在一个单线程服务器？

```
答：养成玩法把玩家做成actor，所有玩家各自独立计算，可以更方便的进行并发处理，比如耗时的循环操作、和其他服务的通信等，
试想单线程模式，玩家服务器几乎不可能用requestWait这样的方法向其他服务通信，会大量影响服务器吞吐量，
一些对全服玩家的循环操作，耗时太长也需要考虑分批完成。相反，把单个玩家设计为独立的actor，可以无需关系这些问题，
也可以大胆使用requestWait进行同步编程。
```


### 2.player actor为什么需要退出？

```
答：非活跃玩家需要(缓存\落地)处理，不能长期留在内存，否则内存和goroutine会随着玩家注册量膨胀，由此
非活跃玩家(缓存\落地)后，actor需要退出，退出后收到消息，会由game先激活player actor再将消息投递。
```


### 3.什么是本地actor？

```
答：不会向cluster注册，不会被其他远端服务所发现,只有本地game能直接Send发送的actor是本地actor，但可以主动像其他服务发送消息，
本地actor通过request、requestWait、Send向其他远端服务通信，远端服务都能正常接收并且可以通过respone返回给发送者，
但是无法通过接收到的sourceId或其他途径得到的actorId直接和本地actor通信，类似于客户端和web服务器通过http通信的关系。
```


### 4.player actor为什么需要是本地？

```
答：因为player actor退出、激活会大量反复触发服务发现，增加etcd和所有服务节点的压力，而且一般的处理逻辑是玩家主动和其他服务交互，
大多数服务不会主动跟player actor交互，需主动跟玩家交互的可以通过将消息发至玩家所在game，由game直接转发或激活转发。
需主动广播给玩家推送消息，比如联盟公告、跑马灯等，可以通过gSession直接走网关和客户端通信
```
