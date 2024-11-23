package vm

import "go-luacompiler/api"

/*
指令：FORPREP

伪代码：

	R(A) -= R(A+2)
	PC += sBx
*/
func forPrep(i Instruction, vm api.LuaVM) {
	a, sbx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPSUB)
	vm.Replace(a)
	vm.AddPC(sbx)
}

/*
指令：FORLOOP

伪代码：

		R(A) += R(A+2)
		if R(A) <?= R(A+1) {
		    PC += sBx
		    R(A+3) = R(A)
	    }
*/
func forLoop(i Instruction, vm api.LuaVM) {
	a, sbx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPADD)
	vm.Replace(a)

	isStepPositive := vm.ToNumber(a+2) > 0
	if (isStepPositive && vm.Compare(a, a+1, api.LUA_OPLE)) ||
		(!isStepPositive && vm.Compare(a+1, a, api.LUA_OPLE)) {
		vm.AddPC(sbx)
		vm.Copy(a, a+3)
	}
}
