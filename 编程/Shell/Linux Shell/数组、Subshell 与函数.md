# 请开始你的表演，数组、Subshell 与函数
# 数组

Bash 支持一维数组和关联数组，暂不支持多维数组。
## 一维数组

直接使用下标访问即可，索引可以是正负整数及其表达式、`*` 和 `@`

| 索引                  | 功能             | 备注                                                       |
| ------------------- | -------------- | -------------------------------------------------------- |
| `n`                 | 获取或修改第 n 个元素   |                                                          |
| `-n`                | 获取或修改倒数第 n 个元素 |                                                          |
| `*`，`@`             | 输出所有变量         | 书上说 `*` 是作为一个整体输出，在循环中视为一个元素，但实测 `*` 和 `@` 都可以用于遍历所有数组元素 |
| `!arr[*]`，`!arr[@]` | 输出所有下标         |                                                          |
> [!note] 索引不连续时，`*` 不会输出不存在的索引

```shell
name[0]="Jacob"
name[1]="Rose"
name[2]="Vicky"
name[3]="Rick"
name[4+4]="TinTin"

echo $name
echo ${name[0]}
echo ${name[1]}
echo ${name[4]}
echo ${name[8]}
echo ${name[-1]}
echo ${name[-2]}
echo ${name[-9]}
echo ${!name[@]}
echo ${!name[*]}

for i in ${name[*]}
do
    echo "*-$i"
done
for j in ${name[@]}
do
    echo "@-$j"
done
```

也可以使用 `()` 创建数组，下标从 0 开始连续创建，以 ` ` 分隔

```shell
name=(Jacob Rose Vicky Rick  TinTin)
echo ${name[*]}
echo ${!name[*]}
```
## 关联数组

使用 `declare -A <数组名>` 创建关联数组，关联数组下标可以是任意字符串。

> [!note] 普通数组与关联数组之间不可以互相转化

```shell
declare -A man
man[name]=TOM
man[age]=26
man[addr]=Beijing
man[phone]=13666666666

echo ${man[name]}
echo ${man[*]}
echo ${!man[*]}
```

关联数组也支持 `()` 初始化

```shell
declare -A woman
woman=([name]=lucy [phone]=13999999999 [age]=27 [addr]=Xian)

echo ${woman[name]}
echo ${woman[*]}
echo ${!woman[*]}
```

---

> [!example] 统计 Nginx 日志
```reference fold
file: "@/_resources/codes/linuxshell/nginx_log.sh"
lang: "shell"
```
# Subshell
## fork

使用 `()` 可以创建一个子进程。子进程继承父进程的上下文。
- 每进入一层子进程，`BASH_SUBSHELL+1`

```reference fold
file: "@/_resources/codes/linuxshell/subshell.sh"
lang: "shell"
```

除此之外，以下符号也会开启子进程：
- 管道运算 `|`
- 分组符号 `()`、分组替换 `$()`
- 后台程序 `&`
- 执行其他程序或脚本

使用 `>` 导出文件不会开启新进程
使用 `source` 载入其他脚本不会开启新进程
## exec

使用 `exec` 执行 Shell 指令将开启一个子进程，并将子进程替代父进程，因此执行结束后会退出。

通常使用 `fork` 的形式执行 `exec` 的代码用于防止覆盖当前脚本

> [!warning] 当 `exec` 后接 `>` 时不会替换当前进程
## source

执行一个脚本，不打开子进程
# 函数

函数声明有多种

```shell
函数名() {
    # do something
}

function 函数名() {
    # do something
}

function 函数名 {
    # do something
}
```

- 直接使用函数名调用函数，不需要 `()`
- 使用 `unset 函数名` 可以取消函数

```reference
file: "@/_resources/codes/linuxshell/usage.sh"
lang: "shell"
```

- 使用 `$n` 访问参数，参数使用空格分隔

```reference
file: "@/_resources/codes/linuxshell/check_service.sh"
lang: "shell"
```
