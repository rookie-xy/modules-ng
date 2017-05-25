/*
 * Copyright (C) 2017 Meng Shi
 */

package stdout

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    STDOUT_MODULE = OUTPUT_MODULE|MAIN_MODULE
    STDOUT_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type Stdout struct{
    *Module_t
}

var stdout = String_t{ len("stdout"), "stdout" }
var stdoutCommands = []Command_t{

    { stdout,
      STDOUT_CONFIG,
      stdoutBlock,
      0,
      0,
      nil },

    NilCommand,
}

func stdoutBlock(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := STDOUT_CONFIG|CONFIG_VALUE
    Block(c, Modules, STDOUT_MODULE, flag)

    return Ok
}

var stdoutModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    stdoutCommands,
    OUTPUT_MODULE,
}

func init() {
    Modules = Load(Modules, &Stdout{stdoutModule})
}