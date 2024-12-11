# ps

使用 `ps` 可以查看当前进程

```shell
ps [-e] [-f]
```

* `-e`：显示所有进程，否则只显示当前用户的进程
* `-f`：以完全格式化的形式展示，显示全部信息
![[Pasted image 20240806165244.png]]

* `UID`：进程所属用户 ID
* `PID`：进程 id
* `PPID`：进程的父进程 id
* `C`：CPU 占用率（百分比）
* `STIME`：进程启动时间
* `TTY`：启动进程的终端号，`?` 表示非终端启动
* `TIME`：进程占用 CPU 时间
* `CMD`：进程的启动路径或启动命令
# kill

向一个进程发送信号，多用于控制进程状态，也可以用于进程间的消息传递

```shell
kill -<signal> <pid>
```

- `pid`：目标进程 id
- `signal`：信号，可以是全称、简称或序号，推荐使用全称
	- 如 `SIGKILL` 简称为 `KILL`，序号为 9，可用于强制结束进程
	- 忽略时，默认为 `SIGTERM`

使用 `kill -l` 可以看到所有信号

```shell
kill -l
```

常用的信号有：

| 信号      | 说明              |
| ------- | --------------- |
| SIGTERM | 省略时的默认信号，终止进程   |
| SIGINT  | 中断进程，相当于 Ctrl+C |
| SIGSTOP | 暂停进程            |
| SIGCONT | 恢复进程            |
| SIGTSTP | 等同于 Ctrl+Z      |
| SIGKILL | 强制杀死进程          |
