/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    FILE_MODULE = INPUT_MODULE|MAIN_MODULE
    FILE_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type Files struct {
    *Module_t
}

var file = String_t{ len("file"), "file" }
var fileCommands = []Command_t{

    { file,
      FILE_CONFIG,
      fileBlock,
      0,
      0,
      nil },

    NilCommand,
}

func fileBlock(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := FILE_CONFIG|CONFIG_VALUE
    Block(c, Modules, FILE_MODULE, flag)

    return Ok
}

var fileModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    fileCommands,
    INPUT_MODULE,
}

func init() {
    Modules = Load(Modules, &Files{Module_t:fileModule})
}
