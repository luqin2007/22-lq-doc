# 运算符

Lua 共有 25 个运算符，逻辑运算符可以直接映射到 Go 对应运算符
## 算术运算符

其他运算符都可以直接映射到 Go 运算符，比较特殊的有：
- 整除（`//`）向 `-∞` 取整，而不是简单的截断
- 取模（`%`）可通过整除定义（`a%b=a-(a//b)*b`）
- 乘方（`^`）和字符串连接（字符串 `..`）具有右结合性，其余为左结合
## 位运算

其他运算符都可以直接映射到 Go 运算符，比较特殊的有：
- 右移为无符号右移
- 左右移 `-n` 位相当于向反方向移动 `n` 位
# 自动类型转换

运算过程中的类型包括：

- 算术运算符
	- 除法、乘方
		- 操作数为整数时，提升为浮点
		- 操作数为字符串时，尝试解析成浮点
		- 结果为浮点
	- 其他
		- 若所有运算数都是整数，则执行整型运算
		- 若操作数中有浮点，全部提升为浮点数后运算，结果为浮点数
- 位运算符
	- 若操作数为浮点数，但实际是整数（如 10.0），且没有超出 `int64` 范围，转换为整型
	- 若操作数为字符串，先尝试转换为整型，失败则尝试转换为浮点再转换为整型
	- 结果为整型
- 字符串拼接
	- 操作数为数字，则转换为字符串
# LuaState 实现
## 隐式类型转换

主要是整型、浮点、字符串的互相转换

```reference
file: "@/_resources/codes/go-luacompiler/number/parser.go"
lang: "go"
start: 5
end: 15
```

在此基础上，实现以任意 Lua 类型转换为整型、浮点型的方法。至此 `ToIntegerX` 和 `ToNumberX` 也可以用这两个方法重写

```reference
file: "@/_resources/codes/go-luacompiler/state/lua_value.go"
lang: "go"
start: 27
end: 60
```

## 算术运算符、位运算符

实现不能直接映射到 Go 运算的运算函数

```reference fold
file: "@/_resources/codes/go-luacompiler/number/math.go"
lang: "go"
start: 5
end: 49
```

将所有运算符都以二元运算符的形式封装，以便后期统一调用

```reference
file: "@/_resources/codes/go-luacompiler/state/arith.go"
lang: "go"
start: 8
end: 29
```

然后以整型、浮点型分类，便于选择方法

```reference
file: "@/_resources/codes/go-luacompiler/state/api_arith.go"
lang: "go"
start: 10
end: 25
```

```reference
file: "@/_resources/codes/go-luacompiler/state/api_arith.go"
lang: "go"
start: 45
end: 70
```
