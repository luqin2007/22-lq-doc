package state

import (
	"fmt"
	"go-luacompiler/binchunk"
)

func (self *luaState) Load(chunk []byte, chunkName string, mode string) int {
	// 暂时先只实现加载二进制数据
	prop := binchunk.Undump(chunk)
	closure := newLuaClosure(prop)
	self.stack.push(closure)
	return 0
}

func (self *luaState) Call(nArgs, nResult int) {
	val := self.stack.get(-nArgs - 1)
	if c, ok := val.(*Closure); ok {
		fmt.Printf("call %s<%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		self.callLuaClosure(nArgs, nResult, c)
	} else {
		panic("not a function or closure!")
	}
}

func (self *luaState) LoadProto(n int) {
	// TODO
}

// callLuaClosure 调用闭包（函数）
//
//	nArgs 实际传入的参数数量
//	nResults 实际需要的参数数量
func (self *luaState) callLuaClosure(nArgs, nResults int, closure *Closure) {
	nRegs := int(closure.proto.MaxStackSize) // 函数所需寄存器大小
	nParams := int(closure.proto.NumParams)  // 函数需要传入的参数
	isVararg := closure.proto.IsVararg == 1  // 函数是否包含变长参数

	// 创建闭包（函数）调用栈
	newStack := newLuaState(nRegs + 20)
	newStack.closure = closure

	// 从当前栈中提取参数和闭包（函数），并将参数存入被调函数栈
	funcAndArgs := self.stack.popN(nArgs + 1)
	newStack.pushN(funcAndArgs[1:], nParams)
	newStack.top = nRegs
	if nArgs > nParams && isVararg {
		newStack.varargs = funcAndArgs[nParams+1:]
	}

	// 将被调函数插入主调函数栈帧，执行后出栈
	self.pushLuaStack(newStack)
	self.runLuaClosure()
	self.popLuaStack()

	// 提取函数执行结果，存入主调函数栈帧
	if nResults != 0 {
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
	}
}
