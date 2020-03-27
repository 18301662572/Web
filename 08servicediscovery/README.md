# Service Discovery 服务发现
临时节点、监视功能、故障时的自动摘除功能，实现一套 服务发现 以及 故障节点摘除 的基本元件

#### endpoint: ip+port 的组合。
#### 服务发现: 通过“订单服务”去找到这些 endpoint 的过程。<br>
#### 负载均衡: 选择把请求发送给哪一台机器，以最大化利用下游机器的过程<br>

### 服务发现：怎么通过服务名字找到 endpoints?
```text
类似于dns服务。

我们自己来设计的话，只需要有一个 kv 存储就可以了。（redis的set或者存在MySQL更可靠）

拿 redis 举例，我们可以用 set 来存储 endpoints 列表:
redis-cli> sadd order_service.http 100.10.1.15:1002
redis-cli> sadd order_service.http 100.10.2.11:1002
redis-cli> sadd order_service.http 100.10.5.121:1002
redis-cli> sadd order_service.http 100.10.6.1:1002
redis-cli> sadd order_service.http 100.10.10.1:1002
redis-cli> sadd order_service.http 100.10.100.11:1002

获取 endpoint 列表也很简单:
127.0.0.1:6379> smembers order_service.http
1) "100.10.1.15:1002"
2) "100.10.5.121:1002"
3) "100.10.10.1:1002"
4) "100.10.100.11:1002"
5) "100.10.2.11:1002"
6) "100.10.6.1:1002"
```

### 故障节点摘除
```text
两种思路：
    1.客户端主动的故障摘除
    2.客户端被动故障摘除。
主动摘除： 我作为依赖其他人的上游，在下游一台机器挂掉的时候，我可以自己主动把它从依赖的节点列表里摘掉。
    两种手段:
        一种是靠应用层心跳，
        还有一种靠请求投票。
被动故障摘除： 依赖出问题了要别人通知我。这个通知一般通过服务注册中心发给我。
```

### watch 数据变化 （监视功能）

### 消息通知


   



