Redis 是一个将数据缓存在内存中，再定期写入文件中的键值对数据库

> [!warning] 官方版本 Windows 下需要 WSL 支持，也可以使用他人编译的[非官方版本](https://github.com/zkteco-home/redis-windows)

```cardlink
url: https://redis.io/docs/latest/get-started/
title: "Community Edition"
description: "Get started with Redis Community Edition"
host: redis.io
favicon: https://redis.io/docs/latest/images/favicons/favicon-196x196.png
```
# 键值操作

## 键值操作
## String

字符串，基本数据类型，二进制安全，因此可以表示任何文本数据和二进制数据

`````col
````col-md
flexGrow=1
===
```redis
set name "BUPT"
get name
```
````
````col-md
flexGrow=1
===
![[../../../_resources/images/Pasted image 20241106190457.png]]
````
`````
## List

> [!attention] 一个 List 最多存储 $2^{32}-1$ 个数据

简单双向链表，支持从头部和尾部插入、删除数据，下标从 0 开始（但控制台输出标号从 1 开始）
`````col
````col-md
flexGrow=1
===
`<key>` 表示列表名
- `LPUSH/RPUSH <key> <value...>`：插入数据
- `LPOP/RPOP <key> <count=1>`：弹出数据
- `LRANGE <key> <start> <end>`：列出列表数据
````
````col-md
flexGrow=1
===
![[../../../_resources/images/Redis 2024-11-06 19.11.55.excalidraw]]
````
`````
## Set

无序集合，基于哈希，操作时间复杂度 `O(1)`

> [!attention] 一个 Set 最多存储 $2^{32}-1$ 个数据

`<key>` 表示 Set 名
- `SADD <key> <value...>`：将一个值添加到集合，若集合中没有该元素则返回 1，否则返回 0
- `SMEMBERS <key>`：列出集合数据
## ZSet

类似 `Set`，但每个键都关联一个 `score` 值（浮点类型），内部数据按关联 `score` 大小排序
- `ZADD <key> <score> <value>`：添加元素
- `ZRANGEBYSCORE <key> <min> <max>`：根据 `score` 大小获取部分值
## Hash

> [!attention] 一个 Hash 最多存储 $2^{32}-1$ 个数据

哈希表，一个以 String 为键，任意类型为值的表。

Redis 称一个 Hash 中数据的键为 `field`（属性标签），`<key>` 表示该 Hash 名
- `HMSET <key> <field> <value> [<field2> <value2> ...]`：添加、修改键值
- `HGET <key> <field>`：获取值
## Stream

实现生产者-消费者模型，多用于消息处理
# 持久化管理

>[!note] 持久化：将数据从内存同步到硬盘
>- 全量数据格式：将内存中的数据写入硬盘，方便下次读取文件时加载
>- 增量请求数据：将内存中的数据序列化为请求，读文件重新请求以得到数据

Redis 支持 RDB（Redis DataBase）和 AOF（AppendOnly File）两种持久化方式，二者可共存
- RDB：将数据压缩保存到硬盘中，本质是一种快照备份机制，默认文件为 `dump.db`
- AOF：以日志形式记录服务器每一个写操作，默认文件为 `appendonly.aof`
持久化方式及触发条件等配置如下：
- `save <sec> <update>`：每 `sec` 中，触发 `update` 次写操作，触发数据持久化
- `appendonly [yes/no]`：是否开启 AOF 持久化，默认为 `no`
- `appendfsync [no/always/everysec]`：AOF 持久化方式
	- `no`：不写入
	- `always`：每接收到一次更新操作，立即将日志写入硬盘
	- `everysec`：每秒写入硬盘一次
# 数据分区

将数据分片，用于横向扩展。根据分区算法，可以是：
- 范围分区：按范围分区，一般要求 Key 为 `object_name:<id>` 的形式
	- 缺点：需要维护一个记录键区间到数据库实例的映射表
- 哈希分区：根据键的哈希值分配

根据执行数据分区的角色不同，分区方式可以是：
- 客户端分区：客户端根据分区算法计算出数据库实例节点
- 代理分区：客户端将键值发给代理，由代理决定数据库实例节点
- 查询路由：客户端请求任何一个数据库实例，数据库实例转发到正确节点

Redis Cluster 是一种**查询路由**与**客户端分区**混合的查询路由，在客户端的辅助下重定向到正确节点
# 集群监控

Redis 主从架构通过哨兵机制监控运行状态。
- 哨兵 Sentinel 是一个分布式系统，一个框架下可以有多个哨兵
- 通过流言（Gossip）协议监视 Master 是否下线
- 通过投票协议（Agreement Protocols）实现故障迁移

哨兵实际是一个特殊模式下的 Redis 服务器，通过 `--sentinel` 运行即可
- 定时任务：每个哨兵维护三个定时任务
	- 向主从节点发送 INFO 命令获取最新主从结构（10s/次，主节点主观下线后 1s/次）
	- 通过发布订阅获取其他哨兵信息
	- 通过向其他节点发送 PING 进行心跳检查，判断下线（1s/次）
- 主观下线：心跳检查失败时，若超过一定时间没有回复，哨兵主观认为节点下线
- 客观下线：某哨兵认为**主节点**主观下线后，向其他哨兵节点查询主节点状态，得到主节点下线的反馈达到一定数量时认为**主节点**客观下线
- 领导者选举：主节点客观下线后，各哨兵节点进行协商，选举出新领导者并进行故障转移
