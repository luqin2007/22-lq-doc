# systemctl

Linux 中开机自启的软件一般称为服务，可通过 `systemctl` 管理

```shell
systemctl start|stop|status|enable|disable 服务名
```

* start：启动
* stop：关闭
* status：状态查询
* enable：开启开机自启
* disable：关闭开机自启

主要的系统内置服务包括：
* NetworkManager：主网络服务
* network：副网络服务
* firewalld：防火墙服务
* sshd：ssh 服务
## target

```shell
systemctl get-default
systemctl set-default [target]
```

| level | target            |             |
| ----- | ----------------- | ----------- |
| 0     |                   | 关机          |
| 1     |                   | 单用户         |
| 2     |                   | 字符多用户，无网络   |
| 3     | mutli-user.target | 字符多用户       |
| 4     |                   | 字符多用户，unset |
| 5     | graphical.target  | 图形化多用户      |
| 6     |                   | 重启          |

