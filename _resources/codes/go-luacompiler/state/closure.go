package state

import (
    "go-luacompiler/api"
    "go-luacompiler/binchunk"
)

type GoFunction func(state api.LuaState) int

type Closure struct {
    proto  *binchunk.Prototype
    goFunc GoFunction
}

func newLuaClosure(proto *binchunk.Prototype) *Closure {
    return &Closure{proto: proto}
}

func newGoClosure(goFunc GoFunction) *Closure {
    return &Closure{goFunc: goFunc}
}
