# 猜谜游戏
## 随机数

- 随机数使用 `math/rand` 包生成，使用 `rand.Seed()` 设置随机数种子
- 通常使用当前时间作为随机数种子，`time.Now()` 可获取当前时间

> [!note] `rand.Seed` 已过时，使用 `rand.New(rand.NewSource(seed))` 创建 `Rand` 对象代替全局的 `rand` 系列函数

## 用户输入

通过 `os.Stdin` 读入用户控制台输入，可以使用 `bufio.NewReader` 包装一下
- `Reader.ReadString(char)`：读取输入，直到某个字符

>[!caution] 注意 `ReadString` 参数中的字符也会在读出的字符串中，如 `\n`，可以使用 `strings.TrimSuffix` 处理
# 在线词典

- 可使用[在线工具](https://curlconverter.com/go/)生成网络请求代码
- `defer`：函数结束时自下而上执行，用于清理环境，关闭流等
# Socks5 服务器

![[../../../../_resources/images/Pasted image 20241111074929.png]]
## 服务器监听

使用 `net.Listen()` 开启一个服务器监听并创建服务器对象
- `server.Accept()` 获取一个请求，返回 `(client, error)`
- 使用 `go 函数调用` 开启一个并发任务
- `client` 是一个 `net.Conn` 对象，也可以使用 `bufio.NewReader` 包装
## 认证

首先接收客户端提供的认证方式，报文信息如下：
- `VER`：协议版本，socks5 为 0x05
- `NMETHODS`：认证方法数量
- `METHODS`：`NMETHODS` 个字节，00 表示不需要认证，02 表示用户名/密码认证
需要返回服务器选择的鉴权认证方式，通过 `Write` 输出一组字节
- `VER`
- `METHOD`
## 请求

请求报文：
- `VER`
- `CMD`：请求类型，CONNECT 请求为 0x01
- `RSV`：0x00，保留字符
- `ATYP`：目标地址类型，决定 `ADDR` 字段长度
	- IPv4 地址：0x01，4 字节
	- 域名：0x03，不定长度
- `DST.ADDR`：目标地址，域名类型第一个字节是长度
- `DST.PORT`：端口，2 字节
响应报文：
- `VER`
- `REP`：00 表示 succeeded
- `RSV`
- `ATYPE`：`BND.ADDR` 类型
- `BND.ADDR`：服务器绑定地址
- `BND.PORT`：服务器绑定端口
## relay

代理与服务器建立连接，在客户端与服务器之间传递数据
- 使用 `net.Dial("tcp", url)` 建立网络连接
- 使用 `io.Copy` 复制读写缓冲区
- 使用管道将输入拷贝到 dial，再将 dial 输出到 client

![[../../../../_resources/images/Pasted image 20241111224019.png]]