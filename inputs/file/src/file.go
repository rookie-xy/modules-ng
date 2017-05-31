/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)


type File struct {
    *Module_t
}

func InitFile() *File {
    return &File{}
}

func (r *File) Listen() int {
    listen := NewListen()
    listen.Accept()
    return Ok
}

func (r *File) Reader() int {
    reader := NewReader()
    reader.Next()
    return Ok
}
