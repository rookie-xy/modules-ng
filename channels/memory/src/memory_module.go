/*
 * Copyright (C) 2017 Meng Shi
 */

package memory

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    MEMORY_MODULE = CHANNEL_MODULE|0x01000000
    MEMORY_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type Memory struct {
    *Module_t
}

var	memory = String_t{ len("memory"), "memory" }
var memoryCommands = []Command_t{

    { memory,
      MEMORY_CONFIG,
      memoryBlock,
      0,
      0,
      nil },

    NilCommand,
}

func memoryBlock(c *Configure, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := MEMORY_CONFIG|CONFIG_VALUE
    Block(c, Modules, MEMORY_MODULE, flag)

    return Ok
}

var memoryModule = &Module{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    memoryCommands,
    CHANNEL_MODULE,
}

func init() {
    Modules = Load(Modules, &Memory{Module_t:memoryModule})
}
