package file

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

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

var (
    group   = String_t{ len("group"), "group" }
    types   = String_t{ len("type"), "type" }
    paths   = String_t{ len("paths"), "paths" }
    publish = String_t{ len("publish"), "publish" }
    codec   = String_t{ len("codec"), "codec" }

    fileInput FileInput
)

var fileInputCommands = []Command_t{

    { group,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(fileInput.group),
      nil },

    { types,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(fileInput.types),
      nil },

    { paths,
      FILE_CONFIG|CONFIG_ARRAY,
      SetArray,
      0,
      unsafe.Offsetof(fileInput.paths),
      nil },

    { publish,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(fileInput.publish),
      nil },

    { codec,
      FILE_CONFIG|CONFIG_BLOCK,
      SetCodec,
      0,
      unsafe.Offsetof(fileInput.codec),
      nil },

    NilCommand,
}
