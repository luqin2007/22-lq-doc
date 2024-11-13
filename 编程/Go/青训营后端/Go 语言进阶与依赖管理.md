# 语言进阶

## 并发

>[!note] 并发：多线程程序在一个核上调度运行

> [!note] 并行：多线程应用在多个核的 CPU 上运行

> [!note] 协程：用户态调度的轻量级线程，由 Go 本身完成调度，栈在 KB 级别

### 开启 goroutine

调用函数之前添加一个 `go` 关键字即可创建一个协程运行
### 通信

推荐使用 Channel 进行通信

> [!note] 通过通信共享内存，而不是通过共享内存实现通信

使用 `make(chain 类型[, 缓冲大小])` 创建通道 Channel。
- 使用 `ch <- v` 将变量  v 发送给通道 `ch`
- 使用 `for v := range ch` 可以循环从通道 `ch` 读取数据
- 通道需要关闭，使用 `defer close(ch)` 关闭

通道可分为有缓冲通道和无缓冲通道，无缓冲通道会使通信时的两个协程同步化，因此又称同步通道。

### 锁

go 仍保留通过共享内存实现通信的方法，通常需要加锁
- 锁：`lock sync.Mutex`
- 加锁：`lock.Lock()`
- 解锁：`lock.Unlock()`
### 并发同步

使用 `WaitGroup` 实现并发同步，实质是通过计数器计数
- `WaitGroup.Add(n)`：添加 n 个并发任务
- `WaitGroup.Done()`：完成一个任务
- `WaitGroup.Wait()`：阻塞等待并发完成
# 依赖管理

由于不同项目依赖版本不同，为了控制依赖版本，Go  依赖管理经过了 `GOPATH`，`GoVendor`，`GoModule` 三个阶段
## GOPATH

`GOPATH` 环境变量即当前项目的目录，项目代码直接依赖 `src` 下的代码，使用 `go get` 直接下载到 `src`
- `bin`：编译生成的二进制文件
- `pkg`：编译中间产物，加速编译
- `src`：项目源码
问题：存在共同的依赖但版本不同时，无法实现 package 多版本控制

## Go Vendor

所有项目及子项目目录下增加 `vendor` 目录，各自的依赖在各自的 `vender` 目录下，也无法控制依赖版本

## Go Module

- 配置文件：`go.mod`，配置文件，描述依赖
- 中心仓库：Proxy，管理依赖库
- 本地工具：`go get/mod`
### go.mod

`go.mod` 文件由三部分组成 
- 依赖版本管理的基本单元，即项目的目录 `module ...`
- go 原生库版本 `go xxx`
- 单元依赖 `require ( ... )`
	- `//indirect`：非直接依赖
	- `incompatible`：主版本 `2+` 的模块但没有增加 `+vN` 后缀，表示可能不兼容

Go 版本管理规则有两种规则：
- 语义化版本
- 基于 commit 的伪版本 （通常为语义化版本+提交时间 + commit 的 hash）

Go 会选择最低的兼容版本

### 依赖分发

Go 支持从以下几种位置下载：
- 直接从 Github 等托管平台下载，但无法保证稳定性和可用性，并增加平台压力
- Proxy：会缓存源站的内容，保证稳定可靠
### 工具

- `go get 包 @<param>`
	- `@update`：默认，获取最新版本
	- `@none`：删除包
	- `@<version>`：获取指定版本
	- `@<hash>`：获取特定 commit
	- `@master`：获取 `master`  分支最新 commit
- `go mod init/download/tidy`
	- `init`：初始化项目，创建 `go.mod`
	- `download`：下载所有依赖
	- `tidy`：增加必要依赖，删除无用依赖
