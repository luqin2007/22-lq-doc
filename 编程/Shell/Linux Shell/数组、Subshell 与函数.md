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


# 函数