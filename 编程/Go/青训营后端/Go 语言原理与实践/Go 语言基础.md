Go 语言特点：高性能，高并发，丰富标准库，快速编译，垃圾回收等
# HelloWorld

```go
package main  
  
import "fmt"  
  
func main() {  
    fmt.Println("Hello World")  
}
```

- 语句尾部没有 `;`
# 变量

变量和常量都支持类型推导。

变量声明：
-  `var <name> [<type>]`
- `<name> := <value>`

常量声明：
- `const <name> [<type>]`
# 流程
## 判断

`if` 与 C 判断类似，区别在于：
- 判断条件无括号  `()`
- `if`、`else` 所在行必须紧跟 `{}`，不能省略或另起一行
## 循环

只有 `for` 循环，分为三种写法：
- `for {}`
- `for <init>; <condition>; <next> {}`
- `for <condition> {}`
## 分支

`go` 语言 `switch` 默认没有 `fall-through`，不需要在每个 `case` 之后添加 `break`

`switch` 有两种形式。一种是传统的 ` switch ` 形式，但变量没有括号，可以支持各种类型

```go
switch <var> {
case <value1>: 
	// do something
case <valueN>:
    // do something
default:
    // do something
}
```

`switch` 还可以用于替代一组 `if-else` 结构，每个 `condition` 是一个 `bool` 判断表达式

```go
switch {
case <condition1>:
    // do something
case <conditionN>:
    // do something
default:
    // do something
}
```
# 集合
## 数组与切片

使用 `var <name>[count]<type>` 声明数组，数组类型表示为 `[]<type>`

切片表示可变长度数组，使用 `make([]<type>, <count>)` 创建，使用 `arr = append(arr, values...)` 扩容

数组与切片都支持类似 `python` 的切片操作
## map

使用 `make(map[<key>]<value>)` 创建 `map`，`<key>`、`value` 表示键值类型，也可以直接使用 `map[<key>]<value>{ k1:v1, k2:v2 }` 创建
- 可以使用下标访问值，可使用两个变量接收，第二个表示是否存在：`v, ok := arr[key]`
- 删除值：`delete(<map>, <key>)`
## range

可以将一个数组转换成下标+内容的形式

```go
package main
import "fmt"
func main() {
    nums := []int{2,3,4}
    for i, num := range nums {
        fmt.Println("index:", i, "num:", num)
    }
}
```

也可以以键值对形式遍历 `map`

```go
package main
import "fmt"
func main() {
    m := map[string]string { "a": "A", "b": "B" }
    for k, v := range m {
        fmt.Println(k, v)
    }
}
```
# 函数

变量类型后置，且支持返回多个值

```go
func <name>(<v> <type>, ...) <return-type> {}
func <name>(<v> <type>, ...) (<rtype1, rtype2, ...>) {
    return a, b, ...
}
```
# 指针

go 支持指针，同样使用 `*`，`&`

```go
func add2(n *int) {
    *n += 2
}

n:=5
add2(&n)
```
# 结构体

使用 `type <name> struct {}` 声明

```go
type user struct {
    name     string
    password string
}
```

结构体支持方法，在结构体之外声明，与普通函数相比在 `func` 与函数名之间多一个括号内表示结构体类型，也支持结构体指针类型

```go
type user struct { ... }
func (u user) checkPassword(pwd string) bool { ... }
func (u *user) resetPassword(pwd string) { ... }
```
# 异常处理

常用一个额外的返回值作为异常信息返回

```go
import "errors"
func findUser(name string) (v *user, err error)
```
# 标准库

使用 `import` 引入库依赖，多个库使用 `import {}`

- 字符串工具：`strings`
- 格式化输出：`fmt.Println`、`fmt.Printf`
	- `%v`：打印各种类型变量
- Json：`encoding/json`，序列化使用 `json.Marshal`，反序列化使用 `json.Unmarshal`
	- 在字段后加一个 \`json:"<name>"\` 标记，可自定义输出的键名
- 时间：`time`
- 字符与其他类型转换：`strconv`
- 进程相关：`os`，`os/exec`
	- `os.Args`：进程运行参数
	- `os.Getenv`，`os.Setenv`：环境变量
	- `exec.Command`：创建子进程