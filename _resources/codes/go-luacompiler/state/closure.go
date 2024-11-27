package state

import (
    "go-luacompiler/api"
    "go-luacompiler/binchunk"
)

type Closure struct {
    proto  *binchunk.Prototype
    goFunc api.GoFunction
    upvals []*upvalue
}

type upvalue struct {
    val *luaValue
}

func newLuaClosure(proto *binchunk.Prototype) *Closure {
    c := &Closure{proto: proto}
    if nUpvals := len(proto.Upvalues); nUpvals > 0 {
        c.upvals = make([]*upvalue, 0, nUpvals)
    }
    return c
}

func newGoClosure(goFunc api.GoFunction) *Closure {
    return &Closure{goFunc: goFunc}
}
