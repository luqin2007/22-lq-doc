Linux 创建远程桌面 VNC 服务

```bash
x0vncserver [options]
```

| 选项                  | 默认值     | 说明                   |
| ------------------- | ------- | -------------------- |
| AcceptKeyEvents     | 1       | 允许操作键盘               |
| AcceptPointerEvents | 1       | 允许使用鼠标               |
| AlwaysShared        | 0       | 允许同时被多人连接            |
| SecurityTypes       | VncAuth | 连接主机的认证方式，None 表示无密码 |
| rfbport             | 5908    | 登录端口                 |
