package api

type LuaVM interface {
	LuaState

	// PC 返回 PC，用于测试
	PC() int
	// AddPC 修改 PC，用于跳转指令
	AddPC(n int)
	// Fetch 取指令，并将 PC 移至下一条指令
	Fetch() uint32
	// GetConst 取出指定常量，推至栈顶
	GetConst(index int)
	// GetRK 将指定常量或栈值推入栈顶，用于 OpArgK 类型参数
	// rk > 0xFF 时表示常量（rk&0xFF），否则为寄存器索引（rk+1）
	GetRK(rk int)
}