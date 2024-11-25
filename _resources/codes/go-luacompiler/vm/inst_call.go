package vm

import "go-luacompiler/api"

/*
指令：CLOSURE (iABx)

伪代码：

	R(A) := closure(KPROTO[Bx])
*/
func closure(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}

/*
指令：CALL (iABC)

伪代码：

	func := R(A)
	R(A), R(A+1), ..., R(A+C-2) := func(R(A+1), R(A+2), ..., R(A+B-1))
*/
func call(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1

	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

func lua_return(i Instruction, vm api.LuaVM) {

}

// _pushFuncAndArgs 将函数和参数推入栈顶
func _pushFuncAndArgs(a, b int, vm api.LuaVM) int {
	if b >= 1 {
		// 有参数
		vm.CheckStack(b)
		for i := a; i < a+b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	} else {
		// _fix
	}
}

func _popResults(a, c int, vm api.LuaVM) {
	if c == 1 {
		// 无返回
	} else if c > 1 {
		for i := a + c - 2; i >= a; i-- {
			vm.Replace(i)
		}
	} else {
		// 暂时将值放在栈中，并存入目标位置
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}

func _fixStack(a int, vm api.LuaVM) {
	// 取出目标位置
	x := int(vm.ToInteger(-1))
	vm.Pop(1)

	vm.CheckStack(x - a)
	for i := a; i < x; i++ {
		vm.PushValue(i)
	}
	// TODO RegisterCount 当前函数寄存器数量
	vm.Rotate(vm.RegisterCount()+1, x-a)
}
