# 闭包

> [!note] 闭包：Closure，按词法作用域捕获了非局部变量的嵌套函数
> 
> 大多数语言使用静态作用域处理非局部变量问题，在编译时确定变量名绑定的是哪个变量。
> 
> Bash、PowerShell 等语言使用动态作用域，变量名绑定哪个变量在运行时确定

> [!tip] 理论上所有 Lua 函数都属于闭包，`main` 函数从外部捕获了 `_ENV` 变量
## Upvalue

闭包内捕获的非局部变量，可通过 `luac -l -l` 查看

```lua
local u, v, w

function f()
    u = v -- 捕获 u, v
end
```

![[../../../_resources/images/Pasted image 20241127235945.png]]
## 扁平闭包

Flat Closure，需要借助外围函数来获得更外围函数局部变量的闭包

```lua
local u, v, w

local function f()
    -- 捕获 u, v 供 g() 捕获
    local function g()
        u == v -- 捕获 u, v
    end
end
```

![[../../../_resources/images/Pasted image 20241128001551.png]]

![[../../../_resources/images/Pasted image 20241128001618.png]]
## 全局变量

全局变量存于 `_ENV` 表中，若函数引用了全局变量，则隐式捕获了 `_ENV` 变量

```lua
local function f()
    local function g()
        x = y -- 捕获全局变量 x, y
    end
end
```

![[../../../_resources/images/Pasted image 20241128001734.png]]
# 底层结构

在 `closure` 结构体中添加一个 `upvals` 字段即可

```reference hl:11
file: "@/_resources/codes/go-luacompiler/state/closure.go"
lang: "go"
start: 8
end: 12
```

1. 为 `main` 函数添加 `_ENV`

```reference hl:22-26
file: "@/_resources/codes/go-luacompiler/state/api_call.go"
lang: "go"
start: 10
end: 28
```

2. 加载函数（闭包）时初始化 `upvalues`

> [!note] `luaStack` 结构添加一个 `openuvs`，即 `OpenUpvalue`，表示开放的 `Upvalue` 集合。
> 开放（Open）状态：Upvalue 捕获的外围函数局部变量仍在栈上，此时直接引用局部函数即可；
> 
> 闭合（Closed）状态：Upvalue 捕获的变量位于其他位置

```reference hl:50-69
file: "@/_resources/codes/go-luacompiler/state/api_call.go"
lang: "go"
start: 45
end: 70
```

3. 完善 `LuaStack` 接口
	- `PushGoFunction`：将 Upvalue 存入 Closure
	- `LuaUpvalueIndex(i int) int`：将注册表伪索引转换为 `Upvalue` 索引
	- `get`，`set`，`isValid`：添加伪索引获取 Upvalue 值
# 指令