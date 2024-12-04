向一个进程发送信号，多用于控制进程状态，也可以用于进程间的消息传递

```shell
kill -<signal> <pid>
```

- `pid`：目标进程 id
- `signal`：信号，可以是全称、简称或序号，推荐使用全称
	- 如 `SIGKILL` 简称为 `KILL`，序号为 9
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
