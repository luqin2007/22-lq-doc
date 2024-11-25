package state

import "go-luacompiler/vm"

type luaState struct {
	stack *luaStack
}

func New(stackSize int) *luaState {
	return &luaState{
		stack: newLuaState(stackSize),
	}
}

// pushLuaStack 调用栈入栈
func (self *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = self.stack
	self.stack = stack
}

// popLuaStack 调用栈出栈
func (self *luaState) popLuaStack() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}

// 执行闭包
func (self *luaState) runLuaClosure() {
	for {
		// 执行指令
		inst := vm.Instruction(self.Fetch())
		inst.Execute(self)
		// RETURN 指令退出
		if inst.Opcode() == vm.OP_RETURN {
			break
		}
	}
}
