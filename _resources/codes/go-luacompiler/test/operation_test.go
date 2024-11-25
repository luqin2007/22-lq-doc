package test

import (
	"github.com/stretchr/testify/assert"
	"go-luacompiler/api"
	"go-luacompiler/state"
	"testing"
)

func TestLuaOp(t *testing.T) {
	ls := state.New(20)

	ls.PushInteger(1)
	ls.PushString("2.0")
	ls.PushString("3.0")
	ls.PushNumber(4.0)
	assert.Equal(t, "[1][\"2.0\"][\"3.0\"][4.000000]", printStack(ls))

	ls.Arith(api.LUA_OPADD)
	assert.Equal(t, "[1][\"2.0\"][7]", printStack(ls))

	ls.Arith(api.LUA_OPBNOT)
	assert.Equal(t, "[1][\"2.0\"][-8]", printStack(ls))

	ls.Len(2)
	assert.Equal(t, "[1][\"2.0\"][-8][3]", printStack(ls))

	ls.Concat(3)
	assert.Equal(t, "[1][\"2.0-83\"]", printStack(ls))

	ls.PushBoolean(ls.Compare(1, 2, api.LUA_OPEQ))
	assert.Equal(t, "[1][\"2.0-83\"][false]", printStack(ls))
}
