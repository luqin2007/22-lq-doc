# 上古神兵利器 sed

sed 会逐行扫描输入数据存入模式空间（一个缓冲区），将其与给定模式匹配，匹配成功则执行对应 sed 指令，否则跳过，最终输出结果

```shell
命令 | sed <选项> '匹配条件和操作指令'
sed <选项> '匹配条件和操作指令' 文件...
```

> [!note] 多组匹配条件和操作指令使用 `;` 分隔
# 基本指令

## 命令选项

| 命令选项        | 功能描述                             |
| ----------- | -------------------------------- |
| -n，--silent | 屏蔽默认输出，不会把数据显示在屏幕上               |
| -r          | 支持扩展正则                           |
| -i[SUFFIX]  | 直接修改源文件，如果设置了SUFFIX后缀名，sed会将数据备份 |
| -e          | 指定需要执行的sed指令，支持使用多个-e参数          |
| -f          | 指定需要执行的脚本文件，需要提前将sed指令写入文件中      |

## 操作指令

| 操作指令              | 功能描述                    |
| ----------------- | ----------------------- |
| p                 | 打印当前匹配的数据行              |
| l                 | 打印当前匹配的数据行，显示控制字符，如回车符等 |
| =                 | 打印当前读取的数据行数             |
| a text            | 在匹配的数据行后面追加文本内容         |
| i text            | 在匹配的数据行前面插入文本内容         |
| d                 | 删除匹配的数据行整行内容（行删除）       |
| c text            | 将匹配的数据行整行内容替换为特定的文本内容   |
| r filename        | 从文件中读取数据并追加到匹配的数据行后面    |
| w filename        | 将当前匹配到的数据写入特定的文件中       |
| q [exit code]     | 立刻退出sed脚本               |
| s/regexp/replace/ | 使用正则匹配，将匹配到的数据替换为特定的内容  |

## 数据定位

| 格式                 | 功能描述                          |
| ------------------ | ----------------------------- |
| number             | 直接根据行号匹配数据                    |
| first~step         | 从frst行开始，步长为step，匹配所有满足条件的数据行 |
| $                  | 匹配最后一行                        |
| /regexp/，\cregexpc | 使用正则表达式匹配数据行，c可以是任意字符         |
| addr1,addr2        | 直接使用行号定位，匹配从addr1到addr2的所有行   |
| addr1,+N           | 直接使用行号定位，匹配从addr1开始及后面的N行     |
| !                  | 条件取反                          |
## 实例

### 根据行号选择行

> [!example] 输出 `/etc/hosts` 的文件内容
> ```shell
> sed 'p' /etc/hosts
> ```

> [!bug] 上面的代码每行会输出两次，因为 `sed` 本身会打印一次
> ```shell
> sed -n 'p' /etc/hosts
> ```

> [!example] 输出 `/etc/hosts` 文件第二行
> ```shell
> sed -n '2p' /etc/hosts
> ```

> [!example] 输出 `/etc/hosts` 文件第 1-3 行
> ```shell
> sed -n '1,3p' /etc/hosts
> ```

> [!example] 输出 `/etc/hosts` 文件第二行到文件末尾
> ```shell
> sed -n '2,$p' /etc/hosts
> ```

> [!example] 输出 `/etc/hosts` 文件第 2,5 行
> ```shell
> sed -n '2p;5p' /etc/hosts
> ```

> [!example] 输出 `/etc/hosts` 文件第 2 行及后 3 行（第 2-5 行）
> ```shell
> sed -n '2,+3p' /etc/hosts
> ```

> [!example] 输出 `/etc/hosts` 文件偶数行（从第 2 行开始，步长为 2）
> ```shell
> sed -n '2~2p' /etc/hosts
> ```

> [!example] 输出 `/etc/hosts` 文件第 1,2,4,5,7,8,... 行（跳过 3,6,9... 行），带行号
> 使用 `cat -n` 为文件添加行号，通过管道传递给 `sed`
> ```shell
> cat -n /etc/hosts | sed -n '3~3!p'
> ```
### 根据内容选择行

> [!example] 输出 `/etc/passwd` 中以 `/bash` 结尾的行
> ```shell
> sed -n '/bash$/p' /etc/passwd
> ```

> [!example] 输出 `/etc/passwd` 中包含以 `s` 开头，以 `:x` 结尾，中间包含三个字符的内容的行
> ```shell
> sed -n '/s...:x/p' /etc/passwd
> ```

> [!example] 输出 `/etc/passwd` 中包含数字的行
> ```shell
> sed -n '/[0-9]/p' /etc/passwd
> ```

> [!example] 输出 `/etc/services` 中以 `http` 开头的行
> ```shell
> sed -n '/^http/p' /etc/services
> ```

> [!example] 输出 `/etc/protocols` 中包含 `icmp` 或 `igmp` 的行
> `|` 条件需要 `-r` 开启扩展正则
> ```shell
> sed -rn '/^(icmp|igmp)/p' /etc/protocols
> ```

> [!example] 输出 `/etc/shells` 中包含 `bash` 的行
> `\xbashx` 中开头结尾的 `x` 可以是任意字符
> ```shell
> sed -n '\xbashxp' /etc/shells
> ```

> [!example] 输出 `/etc/shells` 中不包含 `bash` 的行
> ```shell
> sed -n '/bash/!p' /etc/shells
> ```

> [!example] 输出 `/etc/shells`，同时输出控制字符
> ```shell
> sed -n 'l' /etc/shells
> ```

### 修改整行内容

所有样例均不会修改原文件，而是将修改后的内容输出。如需要保存详见[[#文件保存]]

> [!example] 在 `/etc/hosts` 第一行后追加一行 `add test line`
> ```shell
> sed '1a add test line' /etc/hosts
> ```

> [!example] 在 `/etc/hosts` 第一行前添加一行 `add test line`
> ```shell
> sed '1i add test line' /etc/hosts
> ```

> [!example] 在 `/etc/hosts` 所有包含 `new` 的行后添加一行 `add test line`
> ```shell
> sed '/new/a add test line' /etc/hosts
> ```

> [!example] 删除 `/etc/hosts` 所有偶数行，输出带有行号
> ```shell
> cat -n /etc/hosts | sed '2~2d'
> ```

> [!example] 删除空白行
> ```shell
> sed '/^$/d' /etc/hosts
> ```

> [!example] 删除注释行（`#` 开头）
> ```shell
> sed '/^#/d' /etc/hosts
> ```

> [!example] 将第二行替换为 `modity line`
> ```shell
> sed '2c modity line' /etc/hosts
> ```

> [!example] 将所有行替换为 `all modity`
> ```shell
> sed 'c all modity' /etc/hosts
> ```

> [!example] 在 `/etc/hosts` 每行后追加主机名（`/etc/hostname` 文件内容）
> ```shell
> sed 'r /etc/hostname' /etc/hosts
> ```

> [!example] 在 `/etc/hosts` 结尾追加主机名（`/etc/hostname` 文件内容）
> ```shell
> sed '$r /etc/hostname' /etc/hosts
> ```
### 修改局部内容

- [ ] TODO P298 警告后
### 文件保存

> [!example] 在 `/etc/hosts` 第一行后追加一行 `add test line`，保存并将原文件备份为 `hosts.bak`
> 使用 `-i` 保存，后接一个后缀名作为备份文件名后缀。此时控制台不会有输出
> ```shell
> sed -i.bak '1a add test line' /etc/hosts
> ```

> [!example] 将 `/etc/hosts` 另存为 `/tmp/hosts.bak`
> 使用 `w` 将输出的内容存入指定文件。此时控制台仍有输出
> ```shell
> sed 'w /tmp/hosts.bak' /etc/hosts
> ```

> [!example] 将 `/etc/hosts` 第 1-3 行提取为 `/tmp/hosts.bak`
> ```shell
> sed '1,3w /tmp/hosts.bak' /etc/hosts
> ```
### 退出

> [!example] 读到第三行时终止 `sed`
> ```shell
> sed '3q' /etc/hosts
> ```
# 高级指令

---

# 例：配置 vsftpd 脚本
# 例：配置 DHCP 脚本
# 例：克隆 KVM 虚拟机
# 例：使用 libguestfs 管理 KVM 虚拟机
# 例：配置 SSH 安全策略
# 例：基于 GRUB 配置文件修改内核启动参数
# 例：网络爬虫
# 例：点名抽奖
