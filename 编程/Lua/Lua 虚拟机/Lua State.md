Lua 是一个可嵌入的脚本语言，其他程序嵌入 Lua 链接库就可以使用 Lua  API 获取脚本执行能力。

Lua API 有一百多个函数，包含 `lua_` 开头的基本函数和 `luaL_` 开头的辅助函数，辅助 API 完全通过基本函数实现。相当的 Lua API 都是在操作虚拟栈。

> [!note] 虚拟栈
> Lua 3.1 引入 `lua_State`，4.0 起 `lua_State` 结构走向前台，方便用户切换多个解释器实例，使用 `lua_newstate()` 创建
> 
> `lua_State` 可以看做一个不纯粹的栈（或者说是一个寄存器），支持通过索引访问其中的元素

![[../../../_resources/images/Lua 嵌入 API 2024-11-20 23.37.10.excalidraw]]

# Lua 栈

Lua 栈是宿主语言与 Lua 语言沟通的桥梁，是 Lua State 封装的最基础状态
## 数据类型

Lua 是动态数据类型，变量本身不携带类型信息，变量值携带类型信息。

```lua
local a, b, c
a = 3.14          -- number
b = a             -- number
c = false         -- boolean
c = { 1, 2, 3 }   -- table
c = "hello"       -- string
```

在 Lua 语言中，数据分为 8 种类型，使用 `type(v)` 可以查看。其中 5 种可以直接映射到 Go 语言类型

> [!note] Lua 5.3 `number` 分为浮点数和整型，但没有在语言层面体现，而是用于虚拟机优化。

| Lua 类型     | Go 类型             |
| ---------- | ----------------- |
| `nil`      | `nil`             |
| `boolean`  | `bool`            |
| `number`   | `int64`，`float64` |
| `string`   | `string`          |
| `table`    |                   |
| `function` |                   |
| `thread`   |                   |
| `userdata` |                   |
为每个类型创建对应的常量，由于 Lua 栈按索引存取值，额外增加一个无效类型 `LUA_TNONE`（-1）

```reference
file: "@/_resources/codes/go-luacompiler/api/consts.go"
```
## 栈索引

Lua 栈索引有以下特点：
- 绝对索引：索引从 1 开始，表示栈底
- 相对索引：索引可以为负数，-1 表示栈顶
- 设栈容量为 `n`，栈顶索引为 `top`，可接受索引为 `[1, n]`，有效索引为 `[1, top]`
	- 写入值时必须提供有效索引
	- 读取值时可以提供可接受索引，无效的可接受索引返回 `nil`
![[../../../_resources/images/Lua State 2024-11-21 00.19.41.excalidraw|80%]]
# LuaState

LuaState 接口主要包括以下几类函数：
- LuaStack 栈操作函数，包括辅助插入函数
- LuaStack 索引访问函数

```reference fold
file: "@/_resources/codes/go-luacompiler/api/lua_state_before8.go"
lang: "go"
```

## Rotate

Rotate 将 `[index, top]` 区间的值向栈顶方向旋转 n 次，n < 0 则逆向旋转。`Insert`，`Remove` 都可以使用该函数实现

![[../../../_resources/images/Lua State 2024-11-21 01.38.58.excalidraw|80%]]
Lua 内部通过对 `stack` 的三次 `reverse` 操作实现

![[../../../_resources/images/Lua State 2024-11-21 01.50.24.excalidraw|80%]]
具体实现如下：

```reference
file: "@/_resources/codes/go-luacompiler/state/api_stack.go"
lang: "go"
start: 46
end: 60
```
## SetTop

设置栈顶索引。当给定值小于现有栈顶时，相当于执行 n 次 `pop`；反之，相当于执行 n 次 `push(nil)`

> [!note] Pop(n) 方法可以认为是 SetTop(-n-1)

## IsXxx

判断某个位置是否为某种类型。涉及到三种判断方法：
- `None`，`Nil`：直接判断 `Type(index)`
- `String`：同样判断 `Type(index)`，涉及到类型转换可以是 `LUA_TSTRING` 或 `LUA_TNUMBER`
- `Integer`，`Number`：获取值后判断其类型
## ToXxx

获取 index 位置的对应类型值。大多数包含 `X` 后缀的函数可以判断类型是否匹配

>[!error] `ToString`，`ToStringX` 支持数字类型，但会将对应类型更新成字符串

```reference
file: "@/_resources/codes/go-luacompiler/state/api_access.go"
lang: "go"
start: 106
end: 119
```
